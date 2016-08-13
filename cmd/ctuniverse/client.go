// Original file from gorilla/websocket chat exampl
//
// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Modified by Zachary Giles
// Additional code is under the MIT License, a copy of which is found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/zgiles/ctuniverse"
	"log"
	"net/http"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func goodOrigin(r *http.Request) bool {
	return true
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: goodOrigin,
}

// Short explaination
// hub is a pointer to the hub so the threads can all access it. contention will be taken care of by chan's etc
// conn is a pointer to the websocket instance for this connection so the threads can access it.
// send is a channel to get messages into this connection. The hub will distribute into it. messages coming from this will go into the hub via a channel inside the hub instance
// uuid is this connections id, shows the client has identified itself.. maybe another way later
// attributes are the additional data about the client, how it wants to receive things, flags, etc.
type Client struct {
	hub          *Hub
	conn         *websocket.Conn
	send_object  chan *ctuniverse.UniverseObject
	send_control chan *ctuniverse.UniverseControl
	uuid         string
	attributes   map[string]string
}

func (c *Client) write(mt int, payload []byte) error {
	return c.conn.WriteMessage(mt, payload)
}

// Write to the websocket from the Hub
func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()
	// forever loop
	for {
		select {
		case message, chanopen := <-c.send_object:
			if !chanopen {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if message.Owner != c.uuid {
				w, writeerr := c.conn.NextWriter(websocket.TextMessage)
				if writeerr != nil {
					log.Printf("error: %v", writeerr)
					return
				}
				o := ctuniverse.UniverseMessage{"object", message}
				b, berr := json.Marshal(o)
				if berr != nil {
					log.Printf("error: %v", berr)
					return
				}
				w.Write(b)
				// Maybe optimize for more messages at once like in the chat example, keep simple for now and just close
				closeerr := w.Close()
				if closeerr != nil {
					log.Printf("error: %v", closeerr)
					return
				}
			} // uuids equal
		case message, chanopen := <-c.send_control:
			if !chanopen {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			log.Printf("control: %+v", message)
		}
	}
}

// Write to the Hub from websocket
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		var raw json.RawMessage
		r := ctuniverse.UniverseMessage{O: &raw}
		rerr := json.Unmarshal(message, &r)
		if rerr != nil {
			log.Printf("error: %+v", rerr)
			break
		}
		switch r.Messagetype {
		case "object":
			var o ctuniverse.UniverseObject
			oerr := json.Unmarshal(raw, &o)
			if oerr != nil {
				log.Printf("error: decoding error 4, %+v", o)
				break
			}
			c.uuid = o.Owner
			c.hub.broadcast_object <- &o
		case "control":
			var o ctuniverse.UniverseControl
			oerr := json.Unmarshal(raw, &o)
			if oerr != nil {
				log.Printf("error: decoding error 5")
				break
			}
			c.hub.broadcast_control <- &o
		case "id":
			var o ctuniverse.UniverseID
			oerr := json.Unmarshal(raw, &o)
			if oerr != nil {
				log.Printf("error: decoding error 6")
				break
			}
			c.uuid = o.Uuid
		default:
			log.Printf("Messagetype did not conform to any standard")
		}
	}
}

func wshandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	c := &Client{hub: hub, conn: conn, send_object: make(chan *ctuniverse.UniverseObject, 256), send_control: make(chan *ctuniverse.UniverseControl, 256)}
	c.hub.register <- c
	log.Printf("New Client: %+v", c)
	go c.writePump()
	c.readPump()
}
