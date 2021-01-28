package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/auth"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/usecase"
)

type CharacterHandler interface {
	List(w http.ResponseWriter, r *http.Request)
}

type characterHandler struct {
	characterUseCase usecase.CharacterUseCase
}

func NewCharaHandler(cu usecase.CharacterUseCase) CharacterHandler {
	return &characterHandler{
		characterUseCase: cu,
	}
}

func (c characterHandler) List(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "bad request method", http.StatusMethodNotAllowed)
		return
	}

	xToken := r.Header.Get("x-token")
	if xToken == "" {
		http.Error(w, "x-token is empty", http.StatusUnauthorized)
		return
	}

	if err := auth.Validate(xToken); err != nil {
		log.Println(err)
		http.Error(w, "x-token is invalid", http.StatusUnauthorized)
		return
	}

	userId, err := auth.Get(xToken, "user_id")
	if err != nil {
		log.Println(err)
		http.Error(w, "your token don't have user_id", http.StatusBadRequest)
		return
	}

	list, err := c.characterUseCase.List(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type response struct {
		Characters []*usecase.Character `json:"characters"`
	}

	res := new(response)
	res.Characters = list
	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "couldn't convert to json", http.StatusInternalServerError)
		return
	}
	return
}
