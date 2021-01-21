package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/handler/auth"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"

	"github.com/google/uuid"

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
		http.Error(w, "bad request method", http.StatusBadRequest)
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
		"sub": userId,
		"nbf": regTime,
		"iat": regTime,
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

	account := model.User{
		UserId:   userId,
		UserName: name,
		RegAt:    regTime,
		UpdateAt: regTime,
	}

	if err = u.userUseCase.Add(&account); err != nil {
		http.Error(w, "couldn't create account", http.StatusInternalServerError)
		return
	}

	type response struct {
		token string `json:"token"`
	}

	res := new(response)
	res.token = tokenString
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	return
}

func (u userHandler) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (u userHandler) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
