package core

import (
	"fmt"
	"log"
	"net/http"
)

// Handler model
type Handler struct {
	Manager *Manager
}

// WsHandler handler
func (h *Handler) WsHandler(res http.ResponseWriter, req *http.Request) {
	sock, err := h.Manager.Upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Fatal(err)
	}

	id := req.Header.Get("Sec-Websocket-Key")
	fmt.Println(id)

	var client Client
	client.ID = id
	client.Conn = sock
	client.Send = make(chan []byte)
	client.Manager = h.Manager

	h.Manager.Register <- &client

	// Read message
	go client.Read()

	// Write message
	go client.Write()
}
