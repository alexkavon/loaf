package client

import (
	"github.com/gorilla/websocket"
)

type ClientManager struct {
	//clients    map[*Client]bool
	clients	map[*websocket.Conn]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	user   string
	socket *websocket.Conn
	send   chan []byte
}

func NewClientManager() *ClientManager {
	return &ClientManager{make(map[*websocket.Conn]bool), make(chan []byte), make(chan *Client), make(chan *Client)}
}

func (cm *ClientManager) NewClient(ws *websocket.Conn) {
	//cm.clients[&Client{"asdf", ws, make(chan []byte)}] = true
	cm.clients[ws] = true
}

func (cm *ClientManager) DeleteClient(ws *websocket.Conn) {
	delete(cm.clients, ws)
}
