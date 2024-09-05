package cmd

import (
	"context"
	"eeGame/internal/cmd/middleware"
	"eeGame/internal/controller/ee"
	"eeGame/internal/controller/trade"
	"eeGame/internal/websocket/game/app"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"eeGame/internal/controller/user"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 启动gtoken,生成jwt...等
			gfAdminToken, err := StartBackendGToken()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					middleware.CORS,
				)
				//不需要登录的路由组绑定
				group.Bind(
					user.New().Register,
				)

				//需要登录的路由组绑定
				group.Group("/api", func(group *ghttp.RouterGroup) {
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						trade.New(),
						ee.NewEe(),
					)
				})
			})
			// 启动websocket
			go func() {
				err := app.Run(ctx, "connector001")
				if err != nil {
					return
				}
			}()
			s.Run()
			return nil
		},
	}
)
