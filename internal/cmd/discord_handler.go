package cmd

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/gogf/gf/v2/frame/g"
)

func DiscordProcess(ctx context.Context, bot *discordgo.Session) {
	g.Log().Warning(ctx, "Discord bot started.")
	bot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) { 

	 })()
	g.Log().Warning(ctx, "Discord bot stopped.")
	dcBotLock <- struct{}{}
}
