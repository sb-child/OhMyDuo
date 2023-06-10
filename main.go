package main

import (
	_ "oh-my-duo/internal/logic"
	_ "oh-my-duo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"oh-my-duo/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
)

func main() {
	cmd.Main.Run(gctx.New())
}
