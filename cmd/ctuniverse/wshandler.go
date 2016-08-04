package main

import (
	"io"
	"golang.org/x/net/websocket"
)

func wshandler(ws *websocket.Conn) {
	io.WriteString(ws, "HELLO")
}
