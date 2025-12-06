package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chasenut/chirpy/internal/auth"
	"github.com/chasenut/chirpy/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password 			string	`json:"password"`
		Email				string	`json:"email"`
	}

	type response struct {
		User
		Token			string	`json:"token"`
		RefreshToken	string	`json:"refresh_token"`
	}

	dec := json.NewDecoder(r.Body)
	params := parameters{}
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

	accessToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create authentication token", err)
		return
	}

	refreshToken := auth.MakeRefreshToken()

	dbRefreshToken, err := cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token: refreshToken,
		UserID: user.ID,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 60),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create new refresh token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID: 			user.ID,
			CreatedAt: 		user.CreatedAt,
			UpdatedAt: 		user.UpdatedAt,
			Email: 			user.Email,
			IsChirpyRed: 	user.IsChirpyRed,
		},
		Token: accessToken,
		RefreshToken: dbRefreshToken.Token,
	})
}
