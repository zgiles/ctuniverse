package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

type client struct {
	uuid string
	attributes map[string]string
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := client{}
	fmt.Println("New Client")
	for {
		_ = conn.WriteMessage(websocket.TextMessage, []byte("{ \"uuid\": \"123-212-321-231\" }"))
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg);
		conn.Close()
		return
	}
}
