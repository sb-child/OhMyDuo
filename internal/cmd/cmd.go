package cmd

import (
	"context"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gproc"

	"oh-my-duo/internal/controller/duo"
	"oh-my-duo/internal/service"

	tgbot "github.com/go-telegram/bot"
)

var tgBotLock = make(chan struct{})
var dcBotLock = make(chan struct{})

func StartHttpServer(ctx context.Context) {
	s := g.Server("ohmyduo-http")
	// s.EnablePProf()
	s.Group(g.Config().MustGet(ctx, "ohmyduo.rootDir").String(), func(group *ghttp.RouterGroup) {
		group.GET("/_", duo.ParamsHandler)
		group.GET("/_/:prompt", duo.PromptHandler)
		group.GET("/b/:b64", duo.Base64Handler)
	})
	s.Start()
}

func StartTelegramServer(ctx context.Context, token string) {
	tgBotTimeout := g.Config().MustGet(ctx, "ohmyduo.telegramBotTimeout", 5000).Int()
	tgBotProxy := g.Config().MustGet(ctx, "ohmyduo.telegramBotProxy", "").String()
	c := g.Client()
	c.SetTimeout(time.Duration(tgBotTimeout) * time.Millisecond)
	c.SetProxy(tgBotProxy)
	bot, err := tgbot.New(token,
		tgbot.WithHTTPClient(time.Duration(tgBotTimeout), c),
		tgbot.WithDefaultHandler(TelegramDefaultHandler))
	if err != nil {
		g.Log().Fatal(ctx, "Telegram bot start failed: "+err.Error())
		// todo: auto restart
		tgBotLock <- struct{}{}
	}
	go func(ctx context.Context, x *tgbot.Bot) {
		TelegramProcess(ctx, x)
	}(ctx, bot)
}

func StartDiscordServer(ctx context.Context, token string) {
	dcBotTimeout := g.Config().MustGet(ctx, "ohmyduo.discordBotTimeout", 20000).Int()
	dcBotProxy := g.Config().MustGet(ctx, "ohmyduo.discordBotProxy", "").String()
	bot, err := discordgo.New("Bot " + token)
	c := g.Client()
	c.SetTimeout(time.Duration(dcBotTimeout) * time.Millisecond)
	c.SetProxy(dcBotProxy)
	bot.Client = &c.Client
	if err != nil {
		g.Log().Fatal(ctx, "Discord bot start failed: "+err.Error())
		// todo: auto restart
		dcBotLock <- struct{}{}
	}
	go func(ctx context.Context, x *discordgo.Session) {
		DiscordProcess(ctx, x)
	}(ctx, bot)
}

func MainProcess(ctx context.Context, parser *gcmd.Parser) (err error) {
	httpServer := g.Config().MustGet(ctx, "ohmyduo.httpServer", false).Bool()
	tgBotToken := g.Config().MustGet(ctx, "ohmyduo.telegramBotToken", "").String()
	dcBotToken := g.Config().MustGet(ctx, "ohmyduo.discordBotToken", "").String()
	tgCtx, tgCancel := context.WithCancel(ctx)
	dcCtx, dcCancel := context.WithCancel(ctx)
	service.MyDuo().Init(ctx)
	if httpServer {
		StartHttpServer(ctx)
	}
	if len(tgBotToken) > 0 {
		StartTelegramServer(tgCtx, tgBotToken)
	}
	if len(dcBotToken) > 0 {
		StartDiscordServer(dcCtx, dcBotToken)
	}
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		g.Log().Infof(ctx, "%s Signal received, stopping service...", sig.String())
		if len(tgBotToken) > 0 {
			tgCancel()
			<-tgBotLock
		}
		if len(dcBotToken) > 0 {
			dcCancel()
			<-dcBotLock
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
