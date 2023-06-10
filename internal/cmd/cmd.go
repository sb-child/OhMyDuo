package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"my-duo/internal/controller/duo"
	"my-duo/internal/utils"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start MyDuo backend",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			g.Log().Info(ctx, utils.SplitText("1dd 23456v werty喵 呜..,,//啊啊啊234f,,fe."))
			s.Group(g.Config().MustGet(ctx, "myduo.rootDir").String(), func(group *ghttp.RouterGroup) {
				group.GET("/b/:b64", duo.Base64Handler)
				group.GET("/_", duo.ParamsHandler)
			})
			s.Run()
			return nil
		},
	}
)
