package api

import (
	"net/http"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/usecase"
)

type GachaHandler interface {
	Draw(w http.ResponseWriter, r *http.Request)
}

type gachaHandler struct {
	gachaUseCase usecase.GachaUseCase
}

func NewGachaHandler(ug usecase.GachaUseCase) GachaHandler {
	return &gachaHandler{
		gachaUseCase: ug,
	}
}

func (g gachaHandler) Draw(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
