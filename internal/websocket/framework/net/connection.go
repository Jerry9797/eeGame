package net

type IConnection interface {
	Close()
	SendMsg(buf []byte)
	GetSession() *Session
}

type Message struct {
	Cid     string
	Payload []byte
}
