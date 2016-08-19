// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Copyright 2016 Zachary Giles
// MIT License (Expat)
//
// Please see the LICENSE file

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
	CheckOrigin:     goodOrigin,
}

// Client is a struct for each websocket client
// Short explanation
// hub is a pointer to the hub so the threads can all access it. contention will be taken care of by chan's etc
// conn is a pointer to the websocket instance for this connection so the threads can access it.
// send is a channel to get messages into this connection. The hub will distribute into it. messages coming from this will go into the hub via a channel inside the hub instance
// uuid is this connections id, shows the client has identified itself.. maybe another way later
// attributes are the additional data about the client, how it wants to receive things, flags, etc.
type Client struct {
	hub         *Hub
	conn        *websocket.Conn
	sendObject  chan *ctuniverse.SpaceObject
	sendControl chan *ctuniverse.SpaceControl
	sendChat    chan *ctuniverse.SpaceChat
	uuid        string
	attributes  map[string]string
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
		case message, chanopen := <-c.sendObject:
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
				o := ctuniverse.SpaceMessage{Messagetype: "SpaceObject", O: message}
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
		case message, chanopen := <-c.sendObject:
			if !chanopen {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			w, writeerr := c.conn.NextWriter(websocket.TextMessage)
			if writeerr != nil {
				log.Printf("error: %v", writeerr)
				return
			}
			o := ctuniverse.SpaceMessage{Messagetype: "SpaceChat", O: message}
			b, berr := json.Marshal(o)
			if berr != nil {
				log.Printf("error: %v", berr)
				return
			}
			w.Write(b)
			closeerr := w.Close()
			if closeerr != nil {
				log.Printf("error: %v", closeerr)
				return
			}
		case message, chanopen := <-c.sendControl:
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
		r := ctuniverse.SpaceMessage{O: &raw}
		rerr := json.Unmarshal(message, &r)
		if rerr != nil {
			log.Printf("error: %+v", rerr)
			break
		}
		switch r.Messagetype {
		case "SpaceObject":
			var o ctuniverse.SpaceObject
			oerr := json.Unmarshal(raw, &o)
			if oerr != nil {
				log.Printf("error: decoding error 4, %+v", o)
				break
			}
			c.uuid = o.Owner
			c.hub.broadcastObject <- &o
		case "SpaceControl":
			var o ctuniverse.SpaceControl
			oerr := json.Unmarshal(raw, &o)
			if oerr != nil {
				log.Printf("error: decoding error 5")
				break
			}
			c.hub.broadcastControl <- &o
		case "SpaceID":
			var o ctuniverse.SpaceID
			oerr := json.Unmarshal(raw, &o)
			if oerr != nil {
				log.Printf("error: decoding error 6")
				break
			}
			c.uuid = o.UUID
		case "SpaceChat":
			var o ctuniverse.SpaceChat
			oerr := json.Unmarshal(raw, &o)
			if oerr != nil {
				log.Printf("error: decoding error 7, %+v", o)
				break
			}
			c.hub.broadcastChat <- &o
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
	c := &Client{hub: hub, conn: conn, sendObject: make(chan *ctuniverse.SpaceObject, 256), sendControl: make(chan *ctuniverse.SpaceControl, 256), sendChat: make(chan *ctuniverse.SpaceChat, 256)}
	c.hub.register <- c
	log.Printf("New Client: %+v", c)
	go c.writePump()
	c.readPump()
}
