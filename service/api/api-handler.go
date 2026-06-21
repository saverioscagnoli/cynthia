package api

import "net/http"

func (rt *_router) Handler() http.Handler {
	pmux := http.NewServeMux()

	rt.mux.HandleFunc("GET /healthcheck", rt.wrap(rt.Healthcheck))
	rt.mux.HandleFunc("GET /auth/login", rt.wrap(rt.AuthLogin))
	rt.mux.HandleFunc("GET /auth/callback", rt.wrap(rt.AuthCallback))

	rt.mux.Handle("/user/", rt.authMiddleware(pmux))

	pmux.HandleFunc("GET /user/me", rt.wrap(rt.GetLoggedUser))
	pmux.HandleFunc("PUT /user/username", rt.wrap(rt.UpdateUsername))
	pmux.HandleFunc("PUT /user/sprite-id", rt.wrap(rt.UpdateTrainerSprite))
	pmux.HandleFunc("GET /user/banner", rt.wrap(rt.GetUserBanner))
	pmux.HandleFunc("PUT /user/banner", rt.wrap(rt.UpdateUserBanner))
	pmux.HandleFunc("DELETE /user/banner", rt.wrap(rt.DeleteUserBanner))

	return rt.mux
}

func (rt *_router) Close() error {
	return nil
}
