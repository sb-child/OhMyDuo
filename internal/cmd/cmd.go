package cmd

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gproc"

	"oh-my-duo/internal/controller/duo"
	"oh-my-duo/internal/service"

	tgbot "github.com/go-telegram/bot"
)

var tgBotLock sync.WaitGroup

func StartHttpServer(ctx context.Context) {
	s := g.Server("ohmyduo-http")
	s.Group(g.Config().MustGet(ctx, "ohmyduo.rootDir").String(), func(group *ghttp.RouterGroup) {
		group.GET("/_", duo.ParamsHandler)
		group.GET("/_/:prompt", duo.PromptHandler)
		group.GET("/b/:b64", duo.Base64Handler)
	})
	s.Start()
}

func StartTelegramServer(ctx context.Context, token string) {
	tgBotTimeout := g.Config().MustGet(ctx, "ohmyduo.telegramBotHttpTimeout", 5000).Int()
	tgBotProxy := g.Config().MustGet(ctx, "ohmyduo.telegramBotHttpProxy", "").String()
	c := g.Client()
	c.SetTimeout(time.Duration(tgBotTimeout) * time.Millisecond)
	c.SetProxy(tgBotProxy)
	bot, err := tgbot.New(token,
		tgbot.WithHTTPClient(time.Duration(tgBotTimeout), c),
		tgbot.WithDefaultHandler(TelegramDefaultHandler))
	if err != nil {
		g.Log().Fatal(ctx, "Telegram bot start failed: "+err.Error())
	}
	tgBotLock.Add(1)
	go func(ctx context.Context, x *tgbot.Bot) {
		TelegramProcess(ctx, x)
	}(ctx, bot)
}

func MainProcess(ctx context.Context, parser *gcmd.Parser) (err error) {
	httpServer := g.Config().MustGet(ctx, "ohmyduo.httpServer", false).Bool()
	tgBotToken := g.Config().MustGet(ctx, "ohmyduo.telegramBotToken", "").String()
	tgCtx, tgCancel := context.WithCancel(ctx)
	service.MyDuo().Init(ctx)
	if httpServer {
		StartHttpServer(ctx)
	}
	if len(tgBotToken) > 0 {
		StartTelegramServer(tgCtx, tgBotToken)
	}
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		g.Log().Infof(ctx, "%s Signal received, stopping service...", sig.String())
		if len(tgBotToken) > 0 {
			tgCancel()
			tgBotLock.Wait()
		}
		if httpServer {
			s := g.Server("ohmyduo-http")
			s.Shutdown()
			ghttp.Wait()
		}
	})
	// go func() {
	// 	err := MainCmd(ctx, parser)
	// 	if err != nil {
	// 		g.Log().Warning(ctx, "main process exited with error:", err)
	// 		return
	// 	}
	// 	g.Log().Warning(ctx, "main process exited")
	// }()
	gproc.Listen()
	return err
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start OhMyDuo backend",
		Func:  MainProcess,
	}
)
