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

// Short explaination
// picking up the conn so that go routines on the Client can use the conn
// we have a send chan so that messages can get queued into the Client, however, recv is handled by a go routine.. no need for a channel.
// we need a send chan because we are required to ensure only one writer is sending to the client
type Client struct {
	uuid string
	attributes map[string]string
	conn *websocket.Conn
	send chan []byte
}

// This is write to client from send chan
func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}
	// forever loop
	for {
		select {
			message, ok := <-c.send:
				if !ok {
					// somethings broke. Send a control channel close
					c.write(websocket.CloseMessage, []byte{})
					return
				}
				c.conn.SetWriteDeadline()
		}
	}
}

func wshandler(h *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// we dont know what it's id is yet, so we will leave it blank and error when trying to use it.
	c := &Client{ conn: conn, send: make(chan []byte, 256) }
	fmt.Println("New Client")
	go c.writePump()
	c.readPump()

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
