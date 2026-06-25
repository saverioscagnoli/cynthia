package api

import (
	"log/slog"
	"net/http"
	"time"
)

type RequestContext struct {
	ReqSnowflake string
	Logger       *slog.Logger
}

func (ctx *RequestContext) Error(w http.ResponseWriter, msg string, status int, additionalMsg *string) {
	message := msg

	if additionalMsg != nil {
		message = msg + " " + *additionalMsg
	}

	ctx.Logger.Error("backend", "status", status, "error", message)
	http.Error(w, msg, status)
}

func (ctx *RequestContext) Claims(r *http.Request) *Claims {
	return r.Context().Value(claimsKey).(*Claims)
}

type HttpRouterHandler func(http.ResponseWriter, *http.Request, RequestContext)

func (rt *_router) wrap(fn HttpRouterHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ctx := RequestContext{
			ReqSnowflake: rt.snowflake.Next(),
		}

		ctx.Logger = rt.logger.With("req", ctx.ReqSnowflake, "remote", r.RemoteAddr)
		fn(w, r, ctx)
		ctx.Logger.Info("backend", "method", r.Method, "path", r.URL.Path, "time", time.Since(start))
	}
}
