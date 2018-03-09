package main

type Conn interface {
	Send([]byte) error
	ID() string
}

type packet struct {
	conn    Conn
	message []byte
}
