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
	"github.com/zgiles/ctuniverse"
)

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	broadcast_object  chan *ctuniverse.SpaceObject // Inbound messages from the clients
	broadcast_control chan *ctuniverse.SpaceControl
	register          chan *Client     // Register requests from the clients
	unregister        chan *Client     // Unregister requests from clients
	clients           map[*Client]bool // Registered clients
}

func newHub() *Hub {
	return &Hub{
		broadcast_object:  make(chan *ctuniverse.SpaceObject),
		broadcast_control: make(chan *ctuniverse.SpaceControl),
		register:          make(chan *Client),
		unregister:        make(chan *Client),
		clients:           make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send_object)
				close(client.send_control)
			}
		case message := <-h.broadcast_object:
			// here push message to fellow servers
			for client := range h.clients {
				select {
				case client.send_object <- message:
				default:
					delete(h.clients, client)
					close(client.send_object)
					close(client.send_control)
				}
			}
		}
	}
}
