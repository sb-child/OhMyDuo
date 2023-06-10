package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"my-duo/internal/controller/duo"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start MyDuo backend",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group(g.Config().MustGet(ctx, "myduo.rootDir").String(), func(group *ghttp.RouterGroup) {
				group.GET("/b/:b64", duo.Base64Handler)
				group.GET("/_", duo.ParamsHandler)
			})
			s.Run()
			return nil
		},
	}
)
