package connector

import (
	"eeGame/internal/websocket/framework/net"
	game "eeGame/internal/websocket/game/config"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Connector struct {
	isRunning bool
	wsManager *net.Manager
	handlers  net.LogicHandle
}

func Default() *Connector {
	return &Connector{
		handlers: make(net.LogicHandle),
	}
}

func (c *Connector) Run(serveId string) {
	if !c.isRunning {
		c.wsManager = net.NewManager()
		c.wsManager.ConnectorHandles = c.handlers
		c.Serve(serveId)
	}
}

func (c *Connector) Close() {
	if c.wsManager != nil {
		c.wsManager.Close()
	}
}

func (c *Connector) Serve(serveId string) {
	c.isRunning = true
	// 加载配置文件
	conf := game.Conf.GetConnector(serveId)
	if conf == nil {
		logrus.Error("not found connector game")
	}
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.ClientPort)
	c.wsManager.ServeId = serveId
	c.wsManager.Run(addr)
}

func (c *Connector) RegisterHandler(handlers net.LogicHandle) {
	c.handlers = handlers
}
