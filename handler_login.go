package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chasenut/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password 			string	`json:"password"`
		Email				string	`json:"email"`
		ExpiresInSeconds 	int64	`json:"expires_in_seconds"`
	}

	type response struct {
		User
		Token			string	`json:"token"`
		RefreshToken	string	`json:"refresh_token"`
	}

	dec := json.NewDecoder(r.Body)
	params := &parameters{}
	err := dec.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	match, err := auth.CheckPassowordHash(params.Password, user.HashedPassword)
	if !match || err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	expirationTime := time.Hour
	if params.ExpiresInSeconds > 0 && params.ExpiresInSeconds < 3600 {
		expirationTime = time.Duration(params.ExpiresInSeconds) * time.Second
	}

	accessToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, expirationTime)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create authentication token", err)
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID: 		user.ID,
			CreatedAt: 	user.CreatedAt,
			UpdatedAt: 	user.UpdatedAt,
			Email: 		user.Email,
		},
		Token: accessToken,
	})
}
