package main

import (
	"log"
	"net/http"
)

func main() {

	// Hay que establecer un puerto y el root del filepath y la variable del contador de hits
	const filepathRoot = "."
	const port = "8080"
	var cfg apiConfig
	cfg.fileSystemHits = 0

	//Con la funcion NewServeMux, creamos un multiplexer
	mux := http.NewServeMux()

	// Creamos  los Handlers a aqui
	handlerApp := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	mux.Handle("/app/", cfg.middlewareMetricsInc(handlerApp))

	mux.HandleFunc("/reset", cfg.middlewareMetricsReset)
	mux.HandleFunc("/metrics", cfg.handlerMetrics)
	mux.HandleFunc("/healthz", handlerReadiness)

	// Agregamos un CORS al Mux, aun no entiendo muy bien para que.
	corsMux := middlewareCors(mux)

	// Creamos el servidor llamando al pointer de http con su funcion server, dandole el handler creado
	servidor := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	// Ponemos el servidor  en modo escucha,  que es como debe estar un servidor jeje.
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(servidor.ListenAndServe())
}

// Esta sera la funcion que nos sirva para verificar el estado del servidor, nos da una respuesta simple.
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
