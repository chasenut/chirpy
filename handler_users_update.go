package main

import (
	"encoding/json"
	"net/http"

	"github.com/chasenut/chirpy/internal/auth"
	"github.com/chasenut/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUsersUpdate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email		string	`json:"email"`
		Password	string	`json:"password"`
	}

	type response struct {
		User
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT token", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	dec := json.NewDecoder(r.Body)
	params := parameters{}
	err = dec.Decode(&params) 
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters", err)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create password hash", err)
		return
	}

	user, err := cfg.db.UpdateUserCredentials(r.Context(), database.UpdateUserCredentialsParams{
		ID: 			userID,
		Email: 			params.Email,
		HashedPassword:	hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update user", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User{
			ID: 			user.ID,
			CreatedAt: 		user.CreatedAt,
			UpdatedAt: 		user.UpdatedAt,
			Email: 			user.Email,
			IsChirpyRed: 	user.IsChirpyRed,
		},
	})
}
