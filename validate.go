package main

import (
	"encoding/json"
	"net/http"
	"unicode/utf8"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if utf8.RuneCountInString(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long", nil)
		return
	}

	resp := returnVals{
		Valid: true,
	}
	dat, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dat)
}
