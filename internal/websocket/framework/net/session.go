package net

import "sync"

type Session struct {
	sync.RWMutex
	Cid    string
	Client *WsConnection
	Uid    int64
	data   map[string]any
}

func NewSession(cid string, client *WsConnection) *Session {
	return &Session{
		Cid:    cid,
		Client: client,
		data:   make(map[string]any),
	}
}

func (s *Session) Put(key string, value any) {
	s.Lock()
	defer s.Unlock()
	s.data[key] = value
}

func (s *Session) Get(key string) (any, bool) {
	s.RLock()
	defer s.RUnlock()
	value, ok := s.data[key]
	return value, ok
}

func (s *Session) AddData(uid int64, data map[string]any) {
	s.Lock()
	defer s.Unlock()
	if s.Uid == uid {
		for k, v := range data {
			s.data[k] = v
		}
	}
}
