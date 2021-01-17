package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/db"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/util"

	"github.com/jinzhu/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Controller {
	return &Controller{DB: db}
}

func (ctr *Controller) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var jsonBody map[string]string
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var account db.User

	name := jsonBody["name"]
	if name == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	account.UserName = name
	account.UserId = util.CreateUserId()
	for {
		if account.IsRecord(ctr.DB) {
			account.UserId = util.CreateUserId()
		} else {
			break
		}
	}

	createTimeJST := util.GetJSTTime()
	account.RegAt = createTimeJST
	account.UpdateAt = createTimeJST

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": account.UserId,
		"nbf": account.RegAt,
		"iat": account.RegAt,
	})

	keyData, err := ioutil.ReadFile(os.Getenv("KEY_PATH"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokenString, err := token.SignedString(keyData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = account.Insert(ctr.DB); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(map[string]string{
		"token": tokenString,
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
	return
}

func (ctr *Controller) GetUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := r.Header.Get("x-token")
	parsedToken, err := util.ParsedJWTToken(tokenString)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := parsedToken.Claims.(jwt.MapClaims)
	user, err := db.Get(ctr.DB, token["sub"].(string))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(map[string]string{
		"name": user.UserName,
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
	return
}

func (ctr *Controller) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := r.Header.Get("x-token")
	parsedToken, err := util.ParsedJWTToken(tokenString)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokenMap := parsedToken.Claims.(jwt.MapClaims)

	var accountInfo db.User
	accountInfo.UserId = tokenMap["sub"].(string)
	if accountInfo.IsRecord(ctr.DB) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var jsonBody map[string]string
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var updateInfo db.User
	updateInfo.UserId = tokenMap["sub"].(string)
	updateInfo.UserName = jsonBody["name"]

	updateInfo.UpdateAt = util.GetJSTTime()

	if err := db.Update(ctr.DB, updateInfo); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
