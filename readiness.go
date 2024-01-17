package main

import "net/http"

// Esta sera la funcion que nos sirva para verificar el estado del servidor, nos da una respuesta simple.

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
