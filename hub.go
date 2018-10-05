package main

import (
	"context"
	"sync"
)

type Hub struct {
	ctx    context.Context
	cancel func()

	clients    map[string]Conn
	mu         *sync.RWMutex
	register   chan Conn
	unregister chan Conn

	packet chan packet
}

func NewHub() *Hub {
	ctx, cancel := context.WithCancel(context.Background())
	return &Hub{
		ctx:        ctx,
		cancel:     cancel,
		clients:    make(map[string]Conn),
		mu:         new(sync.RWMutex),
		register:   make(chan Conn),
		unregister: make(chan Conn),
		packet:     make(chan packet),
	}
}

func (h *Hub) AddConn(c Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[c.ID()] = c
}

func (h *Hub) RemoveConn(c Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.clients[c.ID()]; ok {
		delete(h.clients, c.ID())
	}
}

func (h *Hub) MessageHandle(p packet) {
}

func (h *Hub) Broadcast(m []byte) error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, c := range h.clients {
		c.Send(m)
	}

	return nil
}

func (h *Hub) GetClientLen() int {
	return len(h.clients)
}

func (h *Hub) Run() {
	for {
		select {
		case conn := <-h.register:
			h.AddConn(conn)
		case conn := <-h.unregister:
			h.RemoveConn(conn)
		case p := <-h.packet:
			h.MessageHandle(p)
		case <-h.ctx.Done():
			return
		}
	}
}
