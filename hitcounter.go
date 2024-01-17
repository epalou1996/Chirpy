package main

import (
	"fmt"
	"net/http"
)

type apiConfig struct {
	fileSystemHits int
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileSystemHits++
		next.ServeHTTP(w, r)
	})
}
func (cfg *apiConfig) middlewareMetricsReset(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Has resetado con exito el contador"))
	cfg.fileSystemHits = 0
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	x := fmt.Sprint("Hits: " + fmt.Sprint(cfg.fileSystemHits))
	w.Header().Add("Content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(x))

}
