package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/auth"
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
		http.Error(w, "bad request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body couldn't read", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if len(body) == 0 {
		http.Error(w, "body is empty", http.StatusBadRequest)
		return
	}

	var jsonBody map[string]string
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		http.Error(w, "body couldn't convert to json", http.StatusBadRequest)
		return
	}

	name := jsonBody["name"]
	if name == "" {
		http.Error(w, "name is empty", http.StatusBadRequest)
		return
	}

	userID := uuid.New().String()
	for {
		if u.userUseCase.IsRecord(userID) {
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
		http.Error(w, "couldn't create token", http.StatusInternalServerError)
		return
	}

	if err = u.userUseCase.Add(userID, name, regTime); err != nil {
		http.Error(w, "couldn't create account", http.StatusInternalServerError)
		return
	}

	type response struct {
		Token string `json:"token"`
	}

	res := new(response)
	res.Token = token
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	return
}

func (u userHandler) Get(w http.ResponseWriter, r *http.Request) {

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
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	userId, err := auth.Get(xToken, "user_id")
	if err != nil {
		log.Println(err)
		http.Error(w, "your token don't have user_id", http.StatusBadRequest)
		return
	}

	account, err := u.userUseCase.Get(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type response struct {
		Name string `json:"name"`
	}

	res := new(response)
	res.Name = account.Name
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	return
}

func (u userHandler) Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
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
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	userID, err := auth.Get(xToken, "user_id")
	if err != nil {
		log.Println(err)
		http.Error(w, "your token don't have user_id", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body couldn't read", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if len(body) == 0 {
		http.Error(w, "body is empty", http.StatusBadRequest)
		return
	}

	var jsonBody map[string]string
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		http.Error(w, "body couldn't convert to json", http.StatusBadRequest)
		return
	}

	name := jsonBody["name"]
	if name == "" {
		http.Error(w, "name is empty", http.StatusBadRequest)
		return
	}

	if err := u.userUseCase.Update(userID, name); err != nil {
		http.Error(w, "couldn't update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
