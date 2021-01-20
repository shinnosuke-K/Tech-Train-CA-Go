package api

import (
	"net/http"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/usecase"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (u userHandler) Create(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (u userHandler) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (u userHandler) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
