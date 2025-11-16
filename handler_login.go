package main

import (
	"encoding/json"
	"net/http"

	"github.com/chasenut/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password 	string	`json:"password"`
		Email		string	`json:"email"`
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
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve such user", err)
		return
	}

	ok, err := auth.CheckPassowordHash(params.Password, user.HashedPassword)
	if !ok || err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", nil)
		return
	}
	respondWithJSON(w, http.StatusOK, User{
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email: user.Email,
	})
}
