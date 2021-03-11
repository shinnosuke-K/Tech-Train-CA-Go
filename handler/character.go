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

	userID, err := auth.Get(r.Header, "user_id")
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
