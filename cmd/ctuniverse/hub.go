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

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	broadcastObject  chan *ctuniverse.SpaceObject // Inbound messages from the clients
	broadcastControl chan *ctuniverse.SpaceControl
	broadcastChat    chan *ctuniverse.SpaceChat
	register         chan *Client     // Register requests from the clients
	unregister       chan *Client     // Unregister requests from clients
	clients          map[*Client]bool // Registered clients
}

func newHub() *Hub {
	return &Hub{
		broadcastObject:  make(chan *ctuniverse.SpaceObject),
		broadcastControl: make(chan *ctuniverse.SpaceControl),
		broadcastChat:    make(chan *ctuniverse.SpaceChat),
		register:         make(chan *Client),
		unregister:       make(chan *Client),
		clients:          make(map[*Client]bool),
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
				close(client.sendObject)
				close(client.sendControl)
				close(client.sendChat)
			}
		case message := <-h.broadcastObject:
			// here push message to fellow servers
			for client := range h.clients {
				select {
				case client.sendObject <- message:
				default:
					delete(h.clients, client)
					close(client.sendObject)
					close(client.sendControl)
					close(client.sendChat)
				}
			}
		case message := <-h.broadcastChat:
			// here push message to fellow servers
			for client := range h.clients {
				select {
				case client.sendChat <- message:
				default:
					delete(h.clients, client)
					close(client.sendObject)
					close(client.sendControl)
					close(client.sendChat)
				}
			}
		}
	}
}
