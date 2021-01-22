package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/handler/auth"
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

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body couldn't read", http.StatusBadRequest)
		return
	}

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

	userId := uuid.New().String()
	for {
		if u.userUseCase.IsRecord(userId) {
			userId = uuid.New().String()
		}
		break
	}

	regTime := time.Now().Local()

	token := auth.CreateJwt(map[string]interface{}{
		"user_id": userId,
		"nbf":     regTime,
		"iat":     regTime,
	})

	keyData, err := ioutil.ReadFile(os.Getenv("KEY_PATH"))
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	tokenString, err := token.SignedString(keyData)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = u.userUseCase.Add(userId, name, regTime); err != nil {
		http.Error(w, "couldn't create account", http.StatusInternalServerError)
		return
	}

	type response struct {
		Token string `json:"token"`
	}

	res := new(response)
	res.Token = tokenString
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

	authedUser, err := auth.ParseToken(xToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	account, err := u.userUseCase.Get(authedUser.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	type response struct {
		Name string
	}

	res := new(response)
	res.Name = account.UserName
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	return
}

func (u userHandler) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
