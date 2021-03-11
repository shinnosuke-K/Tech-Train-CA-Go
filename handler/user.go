package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/handler/response"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/auth"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/logger"
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

	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, nil, "bad request method")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err, "body couldn't read")
		return
	}
	defer r.Body.Close()

	if len(body) == 0 {
		response.Error(w, http.StatusBadRequest, nil, "body is empty")
		return
	}

	var jsonBody map[string]string
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		response.Error(w, http.StatusBadRequest, err, "body couldn't convert to json")
		return
	}

	name := jsonBody["name"]
	if name == "" {
		response.Error(w, http.StatusBadRequest, nil, "name is empty")
		return
	}

	userID := uuid.New().String()
	for {
		if u.userUseCase.IsRecord(userID) {
			logger.Log.Error(fmt.Sprintf("duplicate user_id : %s", userID))
			userID = uuid.New().String()
		}
		break
	}

	regTime := time.Now().Local()

	token, err := auth.CreateJwtToken(map[string]interface{}{
		"user_id": userID,
		"nbf":     regTime,
		"iat":     regTime,
	})

	if err != nil {
		response.Error(w, http.StatusInternalServerError, nil, "couldn't create token")
		return
	}

	if err = u.userUseCase.Add(userID, name, regTime); err != nil {
		response.Error(w, http.StatusInternalServerError, err, "couldn't create account")
		return
	}

	type resToken struct {
		Token string `json:"token"`
	}

	res := new(resToken)
	res.Token = token
	response.WriteJSON(w, res)
}

func (u userHandler) Get(w http.ResponseWriter, r *http.Request) {

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

	account, err := u.userUseCase.Get(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	type resUser struct {
		Name string `json:"name"`
	}

	res := new(resUser)
	res.Name = account.Name
	response.WriteJSON(w, res)
}

func (u userHandler) Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err, "body couldn't read")
		return
	}
	defer r.Body.Close()

	if len(body) == 0 {
		response.Error(w, http.StatusBadRequest, nil, "body is empty")
		return
	}

	var jsonBody map[string]string
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		response.Error(w, http.StatusBadRequest, err, "body couldn't convert to json")
		return
	}

	name := jsonBody["name"]
	if name == "" {
		response.Error(w, http.StatusBadRequest, nil, "name is empty")
		return
	}

	if err := u.userUseCase.Update(userID, name); err != nil {
		response.Error(w, http.StatusInternalServerError, err, "couldn't update user")
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
