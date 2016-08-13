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
	"github.com/zgiles/ctuniverse"
)

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	broadcast  chan *ctuniverse.UniverseMessage // Inbound messages from the clients
	register   chan *Client                     // Register requests from the clients
	unregister chan *Client                     // Unregister requests from clients
	clients    map[*Client]bool                 // Registered clients
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan *ctuniverse.UniverseMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
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
				close(client.send)
			}
		case message := <-h.broadcast:
			// here push message to fellow servers
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
