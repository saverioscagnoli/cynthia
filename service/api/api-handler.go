package api

import "net/http"

func (rt *_router) Handler() http.Handler {
	rt.mux.HandleFunc("GET /api/healthcheck", rt.wrap(rt.Healthcheck))
	rt.mux.HandleFunc("GET /api/auth/login", rt.wrap(rt.AuthLogin))
	rt.mux.HandleFunc("GET /api/auth/callback", rt.wrap(rt.AuthCallback))
	rt.mux.HandleFunc("POST /api/auth/logout", rt.wrap(rt.AuthLogout))

	// public
	rt.mux.HandleFunc("GET /api/user/{id}", rt.wrap(rt.GetUser))
	rt.mux.HandleFunc("GET /api/user/{id}/banner", rt.wrap(rt.GetUserBanner))
	rt.mux.HandleFunc("GET /api/user/{id}/matches", rt.wrap(rt.GetWinStats))

	// private — registered as exact paths, takes priority over {id}
	auth := rt.authMiddleware
	rt.mux.Handle("GET /api/user/me", auth(http.HandlerFunc(rt.wrap(rt.GetLoggedUser))))
	rt.mux.Handle("PUT /api/user/username", auth(http.HandlerFunc(rt.wrap(rt.UpdateUsername))))
	rt.mux.Handle("PUT /api/user/sprite-id", auth(http.HandlerFunc(rt.wrap(rt.UpdateTrainerSprite))))
	rt.mux.Handle("PUT /api/user/banner", auth(http.HandlerFunc(rt.wrap(rt.UpdateUserBanner))))
	rt.mux.Handle("DELETE /api/user/banner", auth(http.HandlerFunc(rt.wrap(rt.DeleteUserBanner))))

	return rt.mux
}

func (rt *_router) Close() error {
	return nil
}
