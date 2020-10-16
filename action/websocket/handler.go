package websocket

import (
	"encoding/json"
	"fmt"
	"hcc/piccolo/lib/logger"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type connectionACKMessage struct {
	OperationID string `json:"id,omitempty"`
	Type        string `json:"type"`
	Payload     struct {
		Query string `json:"query"`
	} `json:"payload,omitempty"`
}

// Subscriber : Used to receive GraphQL subscription requests
type Subscriber struct {
	ID            int
	Conn          *websocket.Conn
	RequestString string
	OperationID   string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	Subprotocols: []string{"graphql-ws"},
}

// Subscribers : Holding Subscriber structs
var Subscribers sync.Map

// Handler : HTTP handler for websocket
func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Logger.Printf("websocket.Handler(): failed to do websocket upgrade: %v", err)
		return
	}
	connectionACK, err := json.Marshal(map[string]string{
		"type": "connection_ack",
	})
	if err != nil {
		logger.Logger.Printf("websocket.Handler(): failed to marshal ws connection ack: %v", err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, connectionACK); err != nil {
		logger.Logger.Printf("websocket.Handler(): failed to write to ws connection: %v", err)
		return
	}
	go func() {
		for {
			_, p, err := conn.ReadMessage()
			fmt.Println(p)
			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				continue
			}
			if err != nil {
				logger.Logger.Printf("websocket.Handler(): failed to read websocket message: %v", err)
				continue
			}
			var msg connectionACKMessage
			if err := json.Unmarshal(p, &msg); err != nil {
				logger.Logger.Printf("websocket.Handler(): failed to unmarshal: %v", err)
				continue
			}
			fmt.Println(msg)
			if msg.Type == "start" {
				length := 0
				Subscribers.Range(func(key, value interface{}) bool {
					length++
					return true
				})
				var subscriber = Subscriber{
					ID:            length + 1,
					Conn:          conn,
					RequestString: msg.Payload.Query,
					OperationID:   msg.OperationID,
				}
				Subscribers.Store(subscriber.ID, &subscriber)
			}
		}
	}()
}
