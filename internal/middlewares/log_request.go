package middlewares

import (
	"github.com/valyala/fasthttp"
)

// LogRequest logs the query that is being treated.
func LogRequest(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		autheliaCtx := &AutheliaCtx{RequestCtx: ctx}
		logger := NewRequestLogger(autheliaCtx)

		logger.Trace("Request hit")
		next(ctx)
		logger.Tracef("Replied (status=%d)", ctx.Response.StatusCode())
	}
}
