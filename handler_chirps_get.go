package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Invalid chirp ID", err)
		return
	}

	chirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp", err)
		return
	}
	respondWithJSON(w, http.StatusOK, Chirp{
		ID: 		chirp.ID,
		CreatedAt: 	chirp.CreatedAt,
		UpdatedAt:	chirp.UpdatedAt,
		Body: 		chirp.Body,
		UserID: 	chirp.UserID,
	})
}

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
		return
	}

	authorID := uuid.Nil
	authorIDString := r.URL.Query().Get("author_id")
	if authorIDString != "" {
		authorID, err = uuid.Parse(authorIDString)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID", err)
			return
		}
	}
	
	chirps := []Chirp{}
	for _, chirp := range dbChirps {
		if authorID != uuid.Nil && chirp.UserID != authorID {
			continue
		}
		chirps = append(chirps, Chirp{
			ID: 		chirp.ID,
			CreatedAt: 	chirp.CreatedAt,
			UpdatedAt: 	chirp.UpdatedAt,
			Body: 		chirp.Body,
			UserID: 	chirp.UserID,
		})
	}
	respondWithJSON(w, http.StatusOK, chirps)
}
