package main

type Hub struct {
  broadcast chan []byte

}

func newHub() *Hub {
  return &Hub{
    broadcast: make(chan []byte),
    
  }
}

fun (h *Hub) run() {

}
