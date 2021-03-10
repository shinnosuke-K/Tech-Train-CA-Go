package handler

import (
	"encoding/json"
	"net/http"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/auth"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/logger"
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
		logger.Log.Error(err.Error())
		http.Error(w, "x-token is invalid", http.StatusUnauthorized)
		return
	}

	userID, err := auth.Get(xToken, "user_id")
	if err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, "your token don't have user_id", http.StatusBadRequest)
		return
	}

	list, err := c.characterUseCase.List(userID)
	if err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type response struct {
		Characters []*usecase.Character `json:"characters"`
	}

	res := new(response)
	res.Characters = list
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, "couldn't convert to json", http.StatusInternalServerError)
		return
	}
	return
}
