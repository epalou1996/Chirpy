package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	// Hay que establecer un puerto y el root del filepath y la variable del contador de hits
	const filepathRoot = "."
	const port = "8080"
	var cfg apiConfig
	cfg.fileSystemHits = 0

	//Con la funcion NewRouter de chi, creamos un router o un multiplexer, el que sera nuestro basico
	r := chi.NewRouter()

	// Creamos  los Handlers a aqui, el doble handle es porque estamos usando Chi, con ServeMux no es necesario
	// Sin embargo chi es mas efectivo y sencillo a la hora de establecer que metodos queremos permitir para que rutas.

	handlerApp := cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	// En este caso con ServeMux podriamos haber hecho solo un handle de "/app/" que nos hubiera servido para "/app"
	// y los subdirectorios que estan dentro.
	r.Handle("/app/*", handlerApp)
	r.Handle("/app", handlerApp)

	// Router para los elemento de la API y lo montamos sobre el router principal, Importante montar despues de declarar los handles
	// Funciona al reves tambien pero bootdev parecia tener algun tipo de conflicto intentando conectarse al metodo de la otra forma.
	rApi := chi.NewRouter()

	// Con esto simplemente establecemos que metodo queremos, en este caso el get.
	// En caso de querer mas podriamos ponerlos y creo que estaria bien.
	rApi.Get("/reset", cfg.middlewareMetricsReset)
	rApi.Get("/healthz", handlerReadiness)
	rApi.Post("/validate_chirp", validateChirp)
	r.Mount("/api/", rApi)

	// Realizamos lo mismo para las rutas de admin(aun no entiendo la necesidad)
	rAdmin := chi.NewRouter()
	rAdmin.Get("/metrics", cfg.handlerMetrics)
	r.Mount("/admin/", rAdmin)

	// Agregamos un CORS al Mux, aun no entiendo muy bien para que.
	corsMux := middlewareCors(r)

	// Creamos el servidor llamando al pointer de http con su funcion server, dandole el handler creado
	servidor := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	// Ponemos el servidor  en modo escucha,  que es como debe estar un servidor jeje.
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(servidor.ListenAndServe())
}
