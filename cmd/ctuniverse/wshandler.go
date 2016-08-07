package main

import (
	"io"
	"golang.org/x/net/websocket"
)

func wshandler(ws *websocket.Conn) {
	io.WriteString(ws, "{ \"uuid\": \"123-212-321-231\" }")
}
