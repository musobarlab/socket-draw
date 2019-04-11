package main

import (
	"log"
	"net/http"

	"github.com/musobarlab/socket-draw/core"
)

func main() {

	manager := core.NewManager()

	handler := &core.Handler{Manager: manager}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handler.WsHandler)

	go manager.Run()

	log.Println("Listening...")
	http.ListenAndServe(":8000", nil)
}
