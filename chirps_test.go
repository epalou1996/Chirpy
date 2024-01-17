package main

import "testing"

func TestCleanChirps(t *testing.T) {
	result := cleanChirps("Yo kerfuffle todo el dia hermano, fornax")
	expected := "Yo **** todo el dia hermano, ****"

	if result != expected {
		t.Errorf("cleanChirps('Yo kerfuffle todo el dia hermano, fornax') returned %s, expected %s", result, expected)
	}
}
