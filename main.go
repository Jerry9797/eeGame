package main

import (
	_ "eeGame/internal/logic"

	_ "eeGame/internal/packed"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "github.com/sirupsen/logrus"

	"github.com/gogf/gf/v2/os/gctx"

	"eeGame/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
