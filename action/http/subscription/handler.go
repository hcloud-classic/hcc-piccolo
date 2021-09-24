package subscription

import (
	"fmt"
	"github.com/functionalfoundry/graphqlws"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"sync"
)

var opCancelReadLock = sync.Mutex{}
var opCancelWriteLock = sync.Mutex{}
var opCancelList = make(map[string]string)

var piccoloOpIDSplitMagic = "!@#$PicC0Lo"

var connections = make(map[graphqlws.Connection]bool)
var connLock = sync.Mutex{}

func isOpStopped(conn graphqlws.Connection, opID string) bool {
	_, exist := connections[conn]
	if !exist {
		return true
	}

	if !connections[conn] {
		connLock.Lock()
		defer connLock.Unlock()
		delete(connections, conn)
		return true
	}

	cancelOpIDs := strings.Split(opCancelList[conn.ID()], piccoloOpIDSplitMagic)
	for _, cancelOpID := range cancelOpIDs {
		if cancelOpID == opID {
			opCancelWriteLock.Lock()
			opCancelList[conn.ID()] = strings.Replace(opCancelList[conn.ID()], opID+piccoloOpIDSplitMagic, "", -1)
			opCancelWriteLock.Unlock()
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

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var ws, err = upgrader.Upgrade(w, r, nil)

			if err != nil {
				fmt.Println("Failed to establish WebSocket connection", err)
				return
			}

			if ws.Subprotocol() != "graphql-ws" {
				fmt.Println("Connection does not implement the GraphQL WS protocol")
				_ = ws.Close()
				return
			}

			conn := graphqlws.NewConnection(ws, graphqlws.ConnectionConfig{
				EventHandlers: graphqlws.ConnectionEventHandlers{
					Close: func(conn graphqlws.Connection) {
						fmt.Println(
							"conn :", conn.ID(),
							"user :", conn.User(),
							"Closing connection")

						connLock.Lock()
						defer connLock.Unlock()
						connections[conn] = false
					},
					StartOperation: func(
						conn graphqlws.Connection,
						opID string,
						data *graphqlws.StartMessagePayload,
					) []error {
						fmt.Println(
							"conn :", conn.ID(),
							"op :", opID,
							"user :", conn.User(),
							"Start operation")

						data.Query = strings.TrimSpace(data.Query)
						if strings.HasPrefix(data.Query, "subscription") {
							if strings.Contains(data.Query, "telegraf") {
								newTime := ""
								go telegrafSubscription(conn, opID, data, &newTime)
							} else if strings.Contains(data.Query, "resource_usage") ||
								strings.Contains(data.Query, "all_task") ||
								strings.Contains(data.Query, "all_server") ||
								strings.Contains(data.Query, "all_subnet") ||
								strings.Contains(data.Query, "list_server") ||
								strings.Contains(data.Query, "list_subnet") ||
								strings.Contains(data.Query, "list_user") {
								go graphqlCommonSubscription(conn, opID, data)
							}
						}

						return nil
					},
					StopOperation: func(conn graphqlws.Connection, opID string) {
						opCancelWriteLock.Lock()
						opCancelList[conn.ID()] += opID + piccoloOpIDSplitMagic
						opCancelWriteLock.Unlock()

						fmt.Println(
							"conn :", conn.ID(),
							"op :", opID,
							"user :", conn.User(),
							"Stopped operation")
					},
				},
			})

			connLock.Lock()
			defer connLock.Unlock()
			connections[conn] = true
		},
	)
}
