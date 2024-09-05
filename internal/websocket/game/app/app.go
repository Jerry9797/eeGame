package app

import (
	"eeGame/internal/websocket/framework/connector"
	game "eeGame/internal/websocket/game/config"
	"eeGame/internal/websocket/game/route"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(ctx context.Context, serveId string) error {
	game.InitConfig("./internal/websocket/game/config")
	clo := func() {}
	go func() {
		conn := connector.Default()
		// 注册本地处理器
		conn.RegisterHandler(route.Register())
		conn.Run(serveId)
		// 获取关闭函数
		clo = conn.Close
	}()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGHUP)
	stop := func() {
		clo()
		time.Sleep(3 * time.Second)
		logrus.Info("connector app stopped!")
	}
	for {
		select {
		case <-ctx.Done():
			fmt.Println("")
			return nil
		case s := <-stopCh:
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				stop()
				logrus.Warn("connector server exit")
				return nil
			case syscall.SIGHUP:
			default:
				return nil
			}
		}
	}

}
