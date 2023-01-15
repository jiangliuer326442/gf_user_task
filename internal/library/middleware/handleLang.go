package middleware

import (
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func HandleLang(r *ghttp.Request) {
	lang := gconv.String(r.Get("lang", "en"))
	ctx := gi18n.WithLanguage(r.Context(), lang)
	r.SetCtx(ctx)

	r.Middleware.Next()
}
