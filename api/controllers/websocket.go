package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"golang.org/x/net/websocket"
)

// SocketController main
type SocketController struct{}

// NewSocketController represents the controller for operating on the JWT resource
func NewSocketController() *SocketController {
	return &SocketController{}
}

// Socket websocket
func (s SocketController) Socket(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	websocket.Handler(SocketHandler).ServeHTTP(w, r)
}

// Echo websocket
func (s SocketController) Echo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	websocket.Handler(EchoHandler).ServeHTTP(w, r)
}

// SocketHandler Send a buffer to socket
func SocketHandler(ws *websocket.Conn) {
	var in []byte
	if err := websocket.Message.Receive(ws, &in); err != nil {
		return
	}
	fmt.Printf("Received: %s\n", string(in))
	websocket.Message.Send(ws, in)
}

// EchoHandler function
func EchoHandler(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
