package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strings"
)

func validateChirp(w http.ResponseWriter, r *http.Request) {
	type validateChirpResponse struct {
		Valid string `json:"cleaned_body"`
	}

	// El Body tiene que estar en mayusculas porque sino puede decidir no funcionar, no entiendo porque.
	type parameters struct {
		Body string `json:"body"`
	}

	chirp := parameters{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&chirp)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	const maxChirpLength = 140
	if len(chirp.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is to long")
		return
	}
	respondWithJson(w, http.StatusOK, validateChirpResponse{
		Valid: cleanChirps(chirp.Body),
	})

}

// Esta no fue mi idea inicial pero bueno
// Creamos 2 funciones que nos ayudaran a costruir una respuesta apropiada. dicha respuesta apropiada se genera en respondWithJson, respondWithError nos sirve para llamar
// a la primera, declarando errores y mensajes de error en el body.
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"Error"`
	}
	respondWithJson(w, code, errorResponse{
		Error: msg,
	})

}

// Necesita una interfaz de mensaje, ya sea valido o erroneo, y responde enviando el codigo apropiado, y un mensaje en formato json
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}

func cleanChirps(msg string) string {
	badWordsArray := [3]string{"kerfuffle", "sharbert", "fornax"}
	badWordsSlice := badWordsArray[:]

	placeHolder := strings.Split(msg, " ")
	for i, word := range placeHolder {
		if slices.Contains(badWordsSlice, strings.ToLower(word)) {
			placeHolder[i] = "****"
		}
	}
	return strings.Join(placeHolder, " ")

}
