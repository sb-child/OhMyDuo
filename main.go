package main

import (
	_ "my-duo/internal/logic"
	_ "my-duo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"my-duo/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
)

func main() {
	cmd.Main.Run(gctx.New())
}
