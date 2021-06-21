package graphql

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/graphql-go/graphql"
	piccoloWebsocket "hcc/piccolo/action/websocket"
	"hcc/piccolo/lib/logger"
)

// Publisher : Used for publish GraphQL subscription
func Publisher() {
	piccoloWebsocket.Subscribers.Range(func(key, value interface{}) bool {
		subscriber, ok := value.(*piccoloWebsocket.Subscriber)
		if !ok {
			return true
		}
		payload := graphql.Do(graphql.Params{
			Schema:        SchemaExported,
			RequestString: subscriber.RequestString,
		})
		message, err := json.Marshal(map[string]interface{}{
			"type":    "data",
			"id":      subscriber.OperationID,
			"payload": payload,
		})
		if err != nil {
			logger.Logger.Println("websocket.Publisher(): failed to marshal message: " + err.Error())
			return true
		}
		if err := subscriber.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			if err == websocket.ErrCloseSent {
				piccoloWebsocket.Subscribers.Delete(key)
				return true
			}
			logger.Logger.Println("websocket.Publisher(): failed to write to ws connection: " + err.Error())
			return true
		}
		return true
	})
}
