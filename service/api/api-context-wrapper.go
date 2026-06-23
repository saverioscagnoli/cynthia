package api

import (
	"log/slog"
	"net/http"

	"github.com/gofrs/uuid"
)

type RequestContext struct {
	ReqUUID uuid.UUID
	Logger  *slog.Logger
}

type HttpRouterHandler func(http.ResponseWriter, *http.Request, RequestContext)

func (rt *_router) wrap(fn HttpRouterHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqUUID, err := uuid.NewV4()

		if err != nil {
			rt.logger.Error("Cannot generate an uuid for the request", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var ctx = RequestContext{
			ReqUUID: reqUUID,
		}

		ctx.Logger = rt.logger.With("reqUUID", reqUUID, "remote", r.RemoteAddr)
		ctx.Logger.Info("Request", "method", r.Method, "path", r.URL.Path)

		fn(w, r, ctx)
	}
}
