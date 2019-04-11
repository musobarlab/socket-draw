package core

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Manager model
type Manager struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	BroadCast  chan *Message
	Upgrader   websocket.Upgrader
	sync.RWMutex
}

// NewManager function
func NewManager() *Manager {
	clients := make(map[*Client]bool)
	broadCast := make(chan *Message)
	register := make(chan *Client)
	unregister := make(chan *Client)
	upgrader := websocket.Upgrader{
		CheckOrigin: func(*http.Request) bool {
			return true
		},
	}

	return &Manager{
		Clients:    clients,
		BroadCast:  broadCast,
		Register:   register,
		Unregister: unregister,
		Upgrader:   upgrader,
	}

}

// Run function will run Manager process
func (p *Manager) Run() {
	for {
		select {
		case client := <-p.Register:
			p.AddClient(client, true)
		case client := <-p.Unregister:
			if _, ok := p.Clients[client]; ok {
				p.DeleteClient(client)
			}
		case m := <-p.BroadCast:
			messageByte, _ := json.Marshal(m)
			// send to every client that is currently connected
			for client := range p.Clients {
				p.RLock()

				select {
				case client.Send <- messageByte:
				default:
					close(client.Send)
					p.DeleteClient(client)
				}

				p.RUnlock()
			}
		}
	}
}

//AddClient function will push new client to the map clients
func (p *Manager) AddClient(key *Client, b bool) {
	p.Lock()
	p.Clients[key] = b
	p.Unlock()
}

//DeleteClient function will delete client by specific key from map clients
func (p *Manager) DeleteClient(key *Client) {
	p.Lock()
	delete(p.Clients, key)
	p.Unlock()
}
