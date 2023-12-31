package duo

import (
	"encoding/base64"
	"net/http"
	"oh-my-duo/internal/consts"
	"oh-my-duo/internal/service"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/grand"
)

func setHeader(r *ghttp.Request, toJpeg bool) {
	r.Response.Header().Set("Cache-Control", "max-age=60")
	if toJpeg {
		r.Response.Header().Set("Content-Type", "image/jpeg")
		r.Response.Header().Set("Content-Disposition",
			"inline; filename=\"oh-my-duo-"+grand.S(8)+".jpg\"")
	} else {
		r.Response.Header().Set("Content-Type", "image/png")
		r.Response.Header().Set("Content-Disposition",
			"inline; filename=\"oh-my-duo-"+grand.S(8)+".png\"")
	}
}

func Base64Handler(r *ghttp.Request) {
	ctx := r.GetCtx()
	b64Text := r.Get("b64", "").String()
	decoded, err := base64.RawURLEncoding.DecodeString(b64Text)
	if err != nil {
		r.SetError(gerror.NewCode(gcode.CodeInvalidParameter, "BASE64 decode failed"))
		r.Response.Status = http.StatusInternalServerError
		r.Response.Writeln("BASE64 decode failed")
		r.Exit()
	}
	data := consts.MyDuoElements{}
	if err = gjson.DecodeTo(decoded, &data); err != nil {
		r.SetError(gerror.NewCode(gcode.CodeInvalidParameter, "json decode failed"))
		r.Response.Status = http.StatusInternalServerError
		r.Response.Writeln("json decode failed")
		r.Exit()
	}
	data.Character = consts.MyDuoCharactersFromString(data.Character.ToString())
	data.Language = consts.MyDuoLanguageFromString(data.Language.ToString())
	r.Response.Write(service.MyDuo().Draw(ctx, data, data.ToJpeg))
	setHeader(r, data.ToJpeg)
}

func PromptHandler(r *ghttp.Request) {
	ctx := r.GetCtx()
	promptText := r.Get("prompt", "").String()
	// 指定原文和翻译 I-am-Duo|我是多儿
	// 指定原文, 翻译和角色 lily|I-am-Duo|我是多儿
	// todo: 指定原文, 翻译, 角色和语言 english|lily|I-am-Duo|我是多儿
	rounded := true // todo
	character := consts.Duo.ToString()
	language := consts.English.ToString()
	var origin, translated string
	// -- -> -
	// - -> space
	// a--b-cc-d--e
	// a-b cc d-e
	_prompt_splited := strings.Split(promptText, "--")
	promptText = ""
	for i := 0; i < len(_prompt_splited); i++ {
		promptText += strings.ReplaceAll(_prompt_splited[i], "-", " ")
		if len(_prompt_splited)-i > 1 {
			promptText += "-"
		}
	}
	prompts := strings.Split(promptText, "|")
	switch len(prompts) {
	case 2:
		origin, translated = prompts[0], prompts[1]
	case 3:
		character = prompts[0]
		origin, translated = prompts[1], prompts[2]
	}
	if len(origin) <= 0 || len(translated) <= 0 {
		origin = "Hi there, I'm Duo! Can you play with me?"
		translated = "大家好，我是多儿！你能和我玩吗？"
	}
	elem := consts.MyDuoElements{
		Rounded:        rounded,
		Character:      consts.MyDuoCharactersFromString(character),
		Language:       consts.MyDuoLanguageFromString(language),
		OriginText:     origin,
		TranslatedText: translated,
	}
	r.Response.Write(service.MyDuo().Draw(ctx, elem))
	setHeader(r, false)
}

func ParamsHandler(r *ghttp.Request) {
	ctx := r.GetCtx()
	toJpeg := r.GetQuery("j", false).Bool()
	rounded := r.GetQuery("r", true).Bool()
	character := r.GetQuery("c", consts.Duo.ToString()).String()
	language := r.GetQuery("l", consts.English.ToString()).String()
	var origin, translated string
	if r.GetQuery("o") == nil || r.GetQuery("t") == nil {
		origin = "Hi there, I'm Duo! Can you play with me?"
		translated = "大家好，我是多儿！你能和我玩吗？"
	} else {
		origin = r.GetQuery("o").String()
		translated = r.GetQuery("t").String()
	}
	elem := consts.MyDuoElements{
		Rounded:        rounded,
		Character:      consts.MyDuoCharactersFromString(character),
		Language:       consts.MyDuoLanguageFromString(language),
		OriginText:     origin,
		TranslatedText: translated,
	}
	r.Response.Write(service.MyDuo().Draw(ctx, elem, toJpeg))
	setHeader(r, toJpeg)
}
