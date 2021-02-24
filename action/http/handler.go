package http

import (
	"context"
	"fmt"
	"github.com/functionalfoundry/graphqlws"
	"github.com/gorilla/websocket"
	graphqlgo "github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
	"hcc/piccolo/action/graphql"
	piccoloConfig "hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"net/http"
	"strings"
	"sync"
	"time"
)

func telegrafSubscriptionQueryTimeChange(query string, newTime string) string {
	newQuery := ""

	query = strings.Replace(query, "\t", "", -1)
	query = strings.Replace(query, "\n", "", -1)

	querySlice := strings.Split(query, ",")
	for i := range querySlice {
		if strings.Contains(querySlice[i], "time") {
			if strings.Contains(querySlice[i], "$time") {
				continue
			}
			querySlice[i] = strings.Replace(querySlice[i], " ", "", -1)
			s := strings.Split(querySlice[i], ":")
			if len(s) == 2 {
				querySlice[i] = "time: " + "\"" + newTime + "\""
			}
		}
		newQuery += querySlice[i] + ","
	}
	if strings.Contains(newQuery, "$uuid") &&
		!strings.Contains(newQuery, "$time") {
		newQuery = strings.Replace(newQuery, "$uuid: String!", "$time: String!, $uuid: String!", -1)
		newQuery = strings.Replace(newQuery, "uuid: $uuid", "time: $time, uuid: $uuid", -1)
	}

	//logger.Logger.Println("newQuery", newQuery)

	return newQuery
}

func telegrafSubscriptionGetNewTime(dataStr string) string {
	newTime := ""

	if strings.Contains(dataStr, "values") {
		dataStr = strings.TrimSpace(dataStr)
		dataStr = strings.Replace(dataStr, " ", "", -1)
		dataStr = strings.Replace(dataStr, "\t", "", -1)
		dataStr = strings.Replace(dataStr, "\n", "", -1)
		dataStr = strings.Replace(dataStr, "\\", "", -1)
		dataStr = strings.Replace(dataStr, "\"", "", -1)
		dataStr = strings.Replace(dataStr, "[", "", -1)
		dataStr = strings.Replace(dataStr, "]", "", -1)
		slices := strings.Split(dataStr, ":")
		length := len(slices)
		for i := range slices {
			if strings.Contains(slices[i], "values") && length >= i+2 {
				s := strings.Split(slices[i+1], ",")
				newTime = s[0] + "000000"
				break
			}
		}
	}

	//logger.Logger.Println("newTime", newTime)

	return newTime
}

func newSubscriptionHandler(config graphqlws.HandlerConfig) http.Handler {
	var upgrader = websocket.Upgrader{
		CheckOrigin:  func(r *http.Request) bool { return true },
		Subprotocols: []string{"graphql-ws"},
	}

	wsLogger := graphqlws.NewLogger("handler")

	var connections = make(map[graphqlws.Connection]bool)
	connLock := sync.Mutex{}
	cancelLock := sync.Mutex{}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var ws, err = upgrader.Upgrade(w, r, nil)
			cancelList := make(map[string]string)

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
				Authenticate: config.Authenticate,
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

						newTime := ""
						data.Query = strings.TrimSpace(data.Query)
						if strings.HasPrefix(data.Query, "subscription") &&
							strings.Contains(data.Query, "telegraf") {
							go func(
								goroutineConn graphqlws.Connection,
								goroutineOpID string,
								goroutineData *graphqlws.StartMessagePayload) {
								for true {
									ctx := context.Background()

									query := telegrafSubscriptionQueryTimeChange(goroutineData.Query, newTime)
									goroutineData.Variables["time"] = newTime

									params := graphqlgo.Params{
										Schema:         graphql.Schema,
										RequestString:  query,
										VariableValues: goroutineData.Variables,
										OperationName:  goroutineData.OperationName,
										Context:        ctx,
									}
									//logger.Logger.Println("query", query)
									//logger.Logger.Println("goroutineData.Variables", goroutineData.Variables)
									result := graphqlgo.Do(params)

									dataStr := fmt.Sprintf("%v", result.Data)
									newTime = telegrafSubscriptionGetNewTime(dataStr)

									graphqlData := graphqlws.DataMessagePayload{
										Data: result.Data,
										Errors: graphqlws.ErrorsFromGraphQLErrors(result.Errors),
									}
									goroutineConn.SendData(goroutineOpID, &graphqlData)
									if graphqlData.Errors != nil {
										logger.Logger.Println("subscription websocket Error: ", graphqlData.Errors)
									}

									for _, connID := range cancelList {
										if connID == goroutineConn.ID() {
											cancelLock.Lock()
											delete(cancelList, connID)
											cancelLock.Unlock()
											return
										}
									}

									time.Sleep(time.Millisecond * time.Duration(piccoloConfig.GraphQL.SubscriptionInterval))
								}
							}(conn, opID, data)
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
