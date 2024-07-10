package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewMux(sc *ServerConfig) (http.Handler, error) {
	// ref: https://qiita.com/tjun/items/3eea798905b597ec83c0
	// ref: https://peraimaru.work/go-chi%E3%81%A7%E9%9D%99%E7%9A%84%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%82%92%E9%85%8D%E4%BF%A1%E3%81%97%E3%81%9F%E3%81%84%E6%99%82/
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	fileServer := http.FileServer(http.Dir(sc.Static))
	mux.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)
	})

	mux.Get("/index", index)
	return mux, nil
}
