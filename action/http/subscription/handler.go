package subscription

import (
	"github.com/functionalfoundry/graphqlws"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"sync"
)

var opCancelLock = sync.Mutex{}
var opCancelList = make(map[string]string)

var piccoloOpIDSplitMagic = "!@#$PicC0Lo"
var piccoloOpStopAllMagic = "!@#$PizC0Lo"

func isOpStopped(connID string, opID string) bool {
	if opCancelList[connID] == piccoloOpStopAllMagic {
		opCancelLock.Lock()
		delete(opCancelList, connID)
		opCancelLock.Unlock()
		return true
	}

	cancelOpIDs := strings.Split(opCancelList[connID], piccoloOpIDSplitMagic)
	for _, cancelOpID := range cancelOpIDs {
		if cancelOpID == opID {
			opCancelLock.Lock()
			opCancelList[connID] = strings.Replace(opCancelList[connID], opID+piccoloOpIDSplitMagic, "", -1)
			opCancelLock.Unlock()
			return true
		}
	}

	return false
}

// NewSubscriptionHandler : Return customized subscription handler
func NewSubscriptionHandler() http.Handler {
	var upgrader = websocket.Upgrader{
		CheckOrigin:  func(r *http.Request) bool { return true },
		Subprotocols: []string{"graphql-ws"},
	}

	wsLogger := graphqlws.NewLogger("handler")

	var connections = make(map[graphqlws.Connection]bool)
	connLock := sync.Mutex{}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var ws, err = upgrader.Upgrade(w, r, nil)

			if err != nil {
				wsLogger.Warn("Failed to establish WebSocket connection", err)
				return
			}

			if ws.Subprotocol() != "graphql-ws" {
				wsLogger.Warn("Connection does not implement the GraphQL WS protocol")
				_ = ws.Close()
				return
			}

			conn := graphqlws.NewConnection(ws, graphqlws.ConnectionConfig{
				EventHandlers: graphqlws.ConnectionEventHandlers{
					Close: func(conn graphqlws.Connection) {
						wsLogger.WithFields(log.Fields{
							"conn": conn.ID(),
							"user": conn.User(),
						}).Debug("Closing connection")

						opCancelLock.Lock()
						opCancelList[conn.ID()] = piccoloOpStopAllMagic
						opCancelLock.Unlock()
						connLock.Lock()
						delete(connections, conn)
						connLock.Unlock()
					},
					StartOperation: func(
						conn graphqlws.Connection,
						opID string,
						data *graphqlws.StartMessagePayload,
					) []error {
						wsLogger.WithFields(log.Fields{
							"conn": conn.ID(),
							"op":   opID,
							"user": conn.User(),
						}).Debug("Start operation")

						data.Query = strings.TrimSpace(data.Query)
						if strings.HasPrefix(data.Query, "subscription") {
							if strings.Contains(data.Query, "telegraf") {
								newTime := ""
								go telegrafSubscription(conn, opID, data, &newTime)
							} else if strings.Contains(data.Query, "resource_usage") ||
								strings.Contains(data.Query, "all_task") ||
								strings.Contains(data.Query, "all_server") ||
								strings.Contains(data.Query, "all_subnet") {
								go graphqlCommonSubscription(conn, opID, data)
							}
						}

						return nil
					},
					StopOperation: func(conn graphqlws.Connection, opID string) {
						opCancelLock.Lock()
						opCancelList[conn.ID()] += opID + piccoloOpIDSplitMagic
						opCancelLock.Unlock()
					},
				},
			})

			connLock.Lock()
			defer connLock.Unlock()
			connections[conn] = true
		},
	)
}
