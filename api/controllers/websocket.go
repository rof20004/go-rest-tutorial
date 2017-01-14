package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"encoding/json"

	"github.com/gorilla/websocket"
)

// SocketController main
type SocketController struct{}

// NewSocketController represents the controller for operating on the JWT resource
func NewSocketController() *SocketController {
	return &SocketController{}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Socket websocket
func (s SocketController) Socket(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	type Reply struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	obj := &Reply{}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		if error := json.Unmarshal(p, obj); error != nil {
			return
		}

		r, _ := json.Marshal(obj)

		if err = conn.WriteMessage(messageType, r); err != nil {
			return
		}
	}
}
