package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"oh-my-duo/internal/controller/duo"
	"oh-my-duo/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start OhMyDuo backend",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			service.MyDuo().Init(ctx)
			s := g.Server()
			s.Group(g.Config().MustGet(ctx, "ohmyduo.rootDir").String(), func(group *ghttp.RouterGroup) {
				group.GET("/_", duo.ParamsHandler)
				group.GET("/_/:prompt", duo.PromptHandler)
				group.GET("/b/:b64", duo.Base64Handler)
			})
			s.Run()
			return nil
		},
	}
)
