package net

import (
	"eeGame/internal/websocket/framework/protocol"
	game "eeGame/internal/websocket/game/config"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	websocketUpgrade = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type CheckOriginHandle func(r *http.Request) bool

type Manager struct {
	sync.RWMutex
	websocketUpgrade  *websocket.Upgrader
	CheckOriginHandle CheckOriginHandle
	ServeId           string
	serverId          string
	clients           map[string]IConnection
	ReadMsgCh         chan *Message
	EventHandlers     map[protocol.PackageType]MsgHandler
	ConnectorHandles  LogicHandle
}

type HandleFunc func(session *Session, body []byte) (any, error)
type LogicHandle map[string]HandleFunc

type MsgHandler func(packet *protocol.Packet, c IConnection) error

func (m *Manager) Run(addr string) {
	// 处理ReadMsgCh（前端客户端）中的消息
	go m.ReadMsgHandle()

	http.HandleFunc("/", m.Serve)
	m.setupEventHandlers()
	if err := http.ListenAndServe(addr, nil); err != nil {
		logrus.Errorf("ws http.ListenAndServe error %v \n", err)
	}
}

func (m *Manager) Serve(w http.ResponseWriter, r *http.Request) {
	if m.websocketUpgrade == nil {
		m.websocketUpgrade = &websocketUpgrade
	}
	// 将HTTP连接升级为WebSocket连接
	wsConn, err := m.websocketUpgrade.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("Could not open websocket connection : %v \n", err)
		return
	}
	// 开始使用 conn 进行WebSocket通信
	client := NewWsClient(wsConn, m)
	// 保存客户端
	m.addClient(client)
	client.Run()
}

// 消息解析
func (m *Manager) decodeMsg(message *Message) {
	msg, err := protocol.Decode(message.Payload)
	if err != nil {
		logrus.Errorf("ws manager decode message error: %v \n", err)
		return
	}
	if err := m.routeHandle(msg, message.Cid); err != nil {

	}

}

func (m *Manager) Close() {

}

func (m *Manager) addClient(client *WsConnection) {
	m.Lock()
	defer m.Unlock()
	m.clients[client.Cid] = client
}

func (m *Manager) DelClient(ws *WsConnection) {
	for cid, client := range m.clients {
		if cid == ws.Cid {
			client.Close()
			delete(m.clients, cid)
		}
	}
}

func (m *Manager) ReadMsgHandle() {
	for {
		select {
		case msg, ok := <-m.ReadMsgCh:
			if ok {
				m.decodeMsg(msg)
				logrus.Infof("manager m.ReadMsgCh %v \n", msg.Payload)
			}
		}
	}
}

func (m *Manager) routeHandle(msg *protocol.Packet, cid string) error {
	client, ok := m.clients[cid]
	if ok {
		handler, ok := m.EventHandlers[msg.Type]
		if ok {
			return handler(msg, client)
		} else {
			return fmt.Errorf("not found eventHandlers")
		}
	}
	return fmt.Errorf("not found clients[%s]", cid)
}

func (m *Manager) setupEventHandlers() {
	// 根据消息类型，添加处理器
	m.EventHandlers[protocol.Handshake] = m.HandshakeHandler
	m.EventHandlers[protocol.HandshakeAck] = m.HandshakeAckHandler
	m.EventHandlers[protocol.Heartbeat] = m.HeartbeatHandler
	m.EventHandlers[protocol.Data] = m.DataHandler
	m.EventHandlers[protocol.Kick] = m.KickHandler
}

func (m *Manager) HandshakeHandler(packet *protocol.Packet, c IConnection) error {
	logrus.Info("receiver message HandshakeHandler...")
	res := protocol.HandshakeResponse{
		Code: 200,
		Sys: protocol.Sys{
			Heartbeat: 3,
		},
	}
	marshal, _ := json.Marshal(res)
	buf, err := protocol.Encode(packet.Type, marshal)
	if err != nil {
		logrus.Errorf("protocol.Encode error: %v \n", err)
		return err
	}
	c.SendMsg(buf)
	return nil
}

func (m *Manager) HandshakeAckHandler(packet *protocol.Packet, c IConnection) error {
	logrus.Info("receiver message HandshakeAckHandler...")

	return nil
}

func (m *Manager) HeartbeatHandler(packet *protocol.Packet, c IConnection) error {
	logrus.Info("receiver message HeartbeatHandler...")
	var res []byte
	marshal, _ := json.Marshal(res)
	buf, err := protocol.Encode(packet.Type, marshal)
	if err != nil {
		logrus.Errorf("protocol.Encode error: %v \n", err)
		return err
	}
	c.SendMsg(buf)
	return nil
}

// 消息处理
func (m *Manager) DataHandler(packet *protocol.Packet, c IConnection) error {
	logrus.Infof("DataHandler... msg: %v \n", packet.Body)
	code := packet.RouteCode
	route := packet.GetRouteByCode(code)
	split := strings.Split(route, ".")
	if len(split) != 3 {
		return fmt.Errorf("message.Route not supported")
	}
	serverType := split[0]
	// 获取serverType对应的服务器
	connectorConfig := game.Conf.GetConnectorByServerType(serverType)
	handlerMethod := fmt.Sprintf("%s.%s", split[1], split[2])
	if connectorConfig != nil {
		handleFunc, ok := m.ConnectorHandles[handlerMethod]
		if ok {

			data := packet.Body.Data
			mdata, err2 := json.Marshal(data)
			if err2 != nil {
				return err2
			}
			session := c.GetSession()
			session.Uid = int64(packet.Body.Uid)
			session.data[strconv.Itoa(packet.Body.Uid)] = data
			data, err := handleFunc(c.GetSession(), mdata)
			if err != nil {
				return err
			}
			marshal, errM := json.Marshal(data)
			if errM != nil {
				logrus.Errorf("DataHandler marshal data error %v \n", errM)
				return errM
			}
			c.SendMsg(marshal)
		}
	}
	return nil
}

func (m *Manager) KickHandler(packet *protocol.Packet, c IConnection) error {
	logrus.Info("KickHandler...")

	return nil
}

func (m *Manager) selectDst(serverType string) (string, error) {
	serverConfig, ok := game.Conf.ServersConf.TypeServer[serverType]
	if !ok {
		return "", fmt.Errorf("not found server %s", serverType)

	}
	//TODO 使用策略选择一个

	return serverConfig[0].ID, nil
}

func NewManager() *Manager {
	return &Manager{
		ReadMsgCh:     make(chan *Message, 1024),
		clients:       make(map[string]IConnection),
		EventHandlers: map[protocol.PackageType]MsgHandler{},
	}
}
