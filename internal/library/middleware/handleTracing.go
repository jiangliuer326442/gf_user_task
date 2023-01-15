package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
)

func HandleTracing(r *ghttp.Request) {
	ctx, span := gtrace.NewSpan(r.Context(), r.Router.Method+":"+r.Router.Uri)
	defer span.End()
	r.SetCtx(ctx)

	r.Middleware.Next()
}
