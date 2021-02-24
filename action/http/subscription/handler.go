package subscription

import (
	"github.com/functionalfoundry/graphqlws"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"sync"
)

var cancelLock = sync.Mutex{}
var cancelList = make(map[string]string)

func isConnectionClosed(connID string) bool {
	for _, cancelID := range cancelList {
		if cancelID == connID {
			cancelLock.Lock()
			delete(cancelList, cancelID)
			cancelLock.Unlock()
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
								go func(goroutineConn graphqlws.Connection,
									goroutineOpID string,
									goroutineData *graphqlws.StartMessagePayload,
									goroutineNewTime *string) {
									telegrafSubscription(goroutineConn, goroutineOpID, goroutineData, goroutineNewTime)
								}(conn, opID, data, &newTime)
							} // else if strings.Contains(data.Query, "XXX") {} <-- TODO: Add more subscriptions here
						}

						return nil
					},
					StopOperation: func(conn graphqlws.Connection, opID string) {
						cancelLock.Lock()
						cancelList[conn.ID()] = conn.ID()
						cancelLock.Unlock()
					},
				},
			})

			connLock.Lock()
			defer connLock.Unlock()
			connections[conn] = true
		},
	)
}
