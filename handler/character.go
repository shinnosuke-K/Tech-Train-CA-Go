package handler

import (
	"net/http"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/handler/response"
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
		response.Error(w, http.StatusMethodNotAllowed, nil, "bad request method")
		return
	}

	xToken := r.Header.Get("x-token")
	if xToken == "" {
		response.Error(w, http.StatusUnauthorized, nil, "x-token is empty")
		return
	}

	if err := auth.Validate(xToken); err != nil {
		response.Error(w, http.StatusUnauthorized, err, "x-token is invalid")
		return
	}

	userID, err := auth.Get(xToken, "user_id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, err, "your token don't have user_id")
		return
	}

	list, err := c.characterUseCase.List(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	type responseList struct {
		Characters []*usecase.Character `json:"characters"`
	}

	res := new(responseList)
	res.Characters = list
	response.WriteJSON(w, res)
}
