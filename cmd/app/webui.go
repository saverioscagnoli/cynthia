package main

import (
	"cynthia/www"
	"fmt"
	"io/fs"
	"net/http"
	"strings"
)

func registerWebUI(handler http.Handler) (http.Handler, error) {
	dist, err := fs.Sub(www.Dist, "dist")

	if err != nil {
		return nil, fmt.Errorf("failed to embed webui: %w", err)
	}

	fileServer := http.FileServer(http.FS(dist))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			handler.ServeHTTP(w, r)
			return
		}

		fileServer.ServeHTTP(w, r)
	}), nil
}
