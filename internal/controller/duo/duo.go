package duo

import (
	"encoding/base64"
	"my-duo/internal/consts"
	"my-duo/internal/service"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func setHeader(r *ghttp.Request) {
	r.Response.Header().Set("Cache-Control", "max-age=3600")
	r.Response.Header().Set("Content-Type", "image/png")
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
	_ = ctx
	_ = decoded
}

func ParamsHandler(r *ghttp.Request) {
	ctx := r.GetCtx()
	rounded := r.GetQuery("r", true).Bool()
	character := r.GetQuery("c", consts.Duo.ToString()).String()
	language := r.GetQuery("l", consts.English.ToString()).String()
	var origin, translated string
	if r.GetQuery("o") == nil || r.GetQuery("t") == nil {
		origin = "Hi there, This is MyDuo!"
		translated = "大家好，我是麦多!"
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
	_ = ctx
	r.Response.Write(service.MyDuo().Draw(ctx, elem))
	setHeader(r)
}
