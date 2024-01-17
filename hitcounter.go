package main

import (
	"fmt"
	"net/http"
)

// Crearemos un struct que nos permitira guardar el numero de veces que ha sido cargada la pagina desde que encendio el servidor
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
	htmlContent := fmt.Sprintf(`
	<html>

		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>

	</html>`, cfg.fileSystemHits)

	w.Header().Add("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, htmlContent)

}
