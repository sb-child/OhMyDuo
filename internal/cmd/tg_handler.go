package cmd

import (
	"context"
	"oh-my-duo/internal/consts"
	"strings"

	tgbot "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

func TelegramProcess(ctx context.Context, bot *tgbot.Bot) {
	g.Log().Warning(ctx, "Telegram bot started.")
	u, err := bot.GetMe(ctx)
	if err != nil {
		g.Log().Warning(ctx, "Failed to get telegram bot info: ", err.Error())
	}
	g.Log().Infof(ctx, "Telegram bot: @%s [%d]", u.Username, u.ID)
	bot.Start(ctx)
	g.Log().Warning(ctx, "Telegram bot stopped.")
	tgBotLock.Done()
}

func TelegramDefaultHandler(ctx context.Context, bot *tgbot.Bot, update *tgmodels.Update) {
	g.Log().Infof(ctx, "Telegram bot Event[%d]", update.ID)
	if update.InlineQuery != nil {
		g.Log().Infof(ctx, "Inline query: [%d]->[%s]",
			update.InlineQuery.From.ID, update.InlineQuery.Query)
		// urls, chars := TelegramParseDuoCommand(ctx, update.InlineQuery.Query)
		TelegramSendDuoImage(ctx, bot, update.InlineQuery, update.InlineQuery.Query)
	}
}
func TelegramParseDuoCommand(ctx context.Context, cmd string) (urls []string, characters []string) {
	serverBase := g.Config().MustGet(ctx, "ohmyduo.telegramBotImageServer", "").String()
	cmd = strings.TrimSpace(cmd)
	if len(cmd) <= 0 {
		// give some help text
		return []string{consts.MyDuoElements{
			Rounded:        true,
			Character:      consts.Lin,
			Language:       consts.English,
			OriginText:     "Please input the origin text on your message box.",
			TranslatedText: "请在聊天框中输入原文。",
		}.ToUrl(serverBase)}, []string{consts.Lin.ToString()}
	}
	cmds := strings.Split(cmd, "|")
	for i := 0; i < len(cmds); i++ {
		cmds[i] = strings.TrimSpace(cmds[i])
	}
	switch len(cmds) {
	case 1:
		// give some help text
		return []string{consts.MyDuoElements{
			Rounded:        true,
			Character:      consts.Duo,
			Language:       consts.English,
			OriginText:     cmds[0],
			TranslatedText: "[然后请插入\"|\"符号并键入翻译文本]",
			ToJpeg:         true,
		}.ToUrl(serverBase)}, []string{consts.Duo.ToString()}
	case 2:
		// list characters
		r := make([]consts.MyDuoElements, 0, consts.MAX_MyDuoCharacters)
		for i := 0; i < consts.MAX_MyDuoCharacters; i++ {
			r = append(r, consts.MyDuoElements{
				Rounded:        true,
				Character:      consts.MyDuoCharacters(i),
				Language:       consts.English,
				OriginText:     cmds[0],
				TranslatedText: cmds[1],
				ToJpeg:         true,
			})
		}
		rr := make([]string, 0, consts.MAX_MyDuoCharacters)
		rrr := make([]string, 0, consts.MAX_MyDuoCharacters)
		for i := 0; i < consts.MAX_MyDuoCharacters; i++ {
			rr = append(rr, r[i].ToUrl(serverBase))
			rrr = append(rrr, r[i].Character.ToString())
		}
		return rr, rrr
	case 3:
		char := consts.MyDuoCharactersFromString(cmds[2])
		return []string{consts.MyDuoElements{
			Rounded:        true,
			Character:      char,
			Language:       consts.English,
			OriginText:     cmds[0],
			TranslatedText: cmds[1],
			ToJpeg:         true,
		}.ToUrl(serverBase)}, []string{char.ToString()}
	default:
		return []string{consts.MyDuoElements{
			Rounded:        true,
			Character:      consts.Lily,
			Language:       consts.English,
			OriginText:     "Failed to parse your command!",
			TranslatedText: "无法解析您的命令！",
			ToJpeg:         true,
		}.ToUrl(serverBase)}, []string{consts.Lily.ToString()}
	}
	// return consts.MyDuoElements{}, gerror.New("failed to parse")
}

func TelegramSendDuoImage(ctx context.Context, bot *tgbot.Bot, req *tgmodels.InlineQuery, cmd string) {
	results := make([]tgmodels.InlineQueryResult, 0)
	r, rr := TelegramParseDuoCommand(ctx, cmd)
	for i, v := range r {
		kbd := [][]tgmodels.InlineKeyboardButton{
			{{
				Text:                         "俺也试试|Try this prompt",
				SwitchInlineQueryCurrentChat: cmd,
			}},
			{{
				Text: "仓库地址|View this repository",
				URL:  "https://github.com/sb-child/OhMyDuo",
			}},
		}
		results = append(results, &tgmodels.InlineQueryResultPhoto{
			ID:           grand.S(64),
			PhotoURL:     v,
			ThumbnailURL: v,
			PhotoWidth:   793,
			PhotoHeight:  793,
			Title:        rr[i],
			Description:  "Oh My Duo~",
			Caption:      rr[i] + " | Result from github sb-child/OhMyDuo",
			ReplyMarkup:  tgmodels.InlineKeyboardMarkup{InlineKeyboard: kbd},
		})
		// g.Log().Info(ctx, v)
	}

	result := tgbot.AnswerInlineQueryParams{
		InlineQueryID: req.ID,
		Results:       results,
	}
	g.Log().Infof(ctx, "Inline query: [%s] replying...", req.ID)
	_, err := bot.AnswerInlineQuery(ctx, &result)
	if err != nil {
		g.Log().Errorf(ctx, "Inline query: [%s] reply failed: %s", req.ID, err.Error())
		return
	}
	g.Log().Infof(ctx, "Inline query: [%s] replied", req.ID)
}
