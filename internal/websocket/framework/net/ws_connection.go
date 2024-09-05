package net

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"sync/atomic"
	"time"
)

type WsConnection struct {
	Cid        string
	Conn       *websocket.Conn
	Manager    *Manager
	ReadMsgCh  chan *Message
	WriteMsgCh chan []byte
	Session    *Session
}

func (ws *WsConnection) GetSession() *Session {
	return ws.Session
}

var (
	cidBase      int64 = 1000
	pongWait           = 10 * time.Second
	pingInterval       = (pongWait * 9) / 10 // 需要比pongWait少
)

func NewWsClient(wsConn *websocket.Conn, manager *Manager) *WsConnection {
	cid := fmt.Sprintf("%s-%s-%s", uuid.New().String(), manager.ServeId, atomic.AddInt64(&cidBase, 1))
	client := &WsConnection{
		Cid:        cid,
		Conn:       wsConn,
		Manager:    manager,
		WriteMsgCh: make(chan []byte, 1024),
		ReadMsgCh:  manager.ReadMsgCh,
	}
	client.Session = NewSession(cid, client)
	return client
}

func (ws *WsConnection) Run() {
	go ws.readMsg()
	go ws.writeMsg()
	// Heartbeat detection
	ws.Conn.SetPongHandler(ws.PongHandler)
}

// 将消息放入ws.ReadMsgCh处理
func (ws *WsConnection) readMsg() {
	defer func() {
		time.Sleep(3 * time.Second)
		ws.Manager.DelClient(ws)
	}()
	ws.Conn.SetReadLimit(1024)
	if err := ws.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		logrus.Errorf("ws.Conn.SetReadDeadline error: %v \n", err)
		return
	}
	for {
		messageType, msg, err := ws.Conn.ReadMessage()
		if err != nil {
			logrus.Errorf("read message error: %v \n", err)
			break
		}
		// 接受二进制消息
		if messageType == websocket.BinaryMessage {
			if ws.ReadMsgCh != nil {
				// 写到manager，manager处理
				ws.ReadMsgCh <- &Message{
					Cid:     ws.Cid,
					Payload: msg,
				}
			}
		} else {
			logrus.Errorf("unsupported message type : %d", messageType)
		}
	}
}

func (ws *WsConnection) SendMsg(buf []byte) {
	ws.WriteMsgCh <- buf
}

func (ws *WsConnection) writeMsg() {
	ticker := time.NewTicker(pingInterval)
	for {
		select {
		case msg, ok := <-ws.WriteMsgCh:
			if ok {
				if err := ws.Conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
					logrus.Errorf("ws.Conn.WriteMessage error: %v \n", err)
				}
			} else {
				logrus.Errorf("ws.WriteMsgCh is : %v \n", ok)
			}
		case <-ticker.C:
			if ws.Manager.clients[ws.Cid] != nil {
				if err := ws.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
					logrus.Errorf("pong error: %v \n", err)
				}
				if err := ws.Conn.WriteMessage(websocket.BinaryMessage, nil); err != nil {
					logrus.Errorf("ws.Conn[%s] ping error: %v \n", ws.Cid, err)
				}
			}
		}
	}
}

func (ws *WsConnection) Close() {
	if ws.Conn != nil {
		err := ws.Conn.Close()
		if err != nil {
			logrus.Errorf("ws.Conn.Close error: %v \n", err)
		}
	}
}

func (ws *WsConnection) PongHandler(appData string) error {
	logrus.Info("pong...")
	if err := ws.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		logrus.Errorf("pong error: %v \n", err)
		return err
	}
	return nil
}
