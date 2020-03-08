package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/db"

	"github.com/dgrijalva/jwt-go"

	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
)

type Server struct {
	Engine *http.ServeMux
}

type Model struct {
	db *gorm.DB
}

func NewServer() *Server {
	return &Server{
		Engine: http.NewServeMux(),
	}
}

func (model *Model) createUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var jsonBody map[string]string
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
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
	account.UserId = createUserId()
	// 要検討
	go func(account *db.User) {
		for {
			if account.IsRecord(model.db) {
				account.UserId = createUserId()
			} else {
				break
			}
		}
	}(&account)

	createTimeUTC := time.Now().UTC()
	jst, _ := time.LoadLocation("Asia/Tokyo")
	createTimeJST := createTimeUTC.In(jst)

	account.RegTimeJST = createTimeJST
	account.UpdateTimeJST = createTimeJST

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  account.UserId,
		"user": account.UserName,
		"nbf":  account.RegTimeJST,
		"iat":  account.RegTimeJST,
	})

	keyData, err := ioutil.ReadFile(os.Getenv("KEY_PATH"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokenString, err := token.SignedString(keyData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	account.Token = tokenString

	if err = account.Insert(model.db); err != nil {
		fmt.Println(1)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(map[string]string{
		"token": account.Token,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
	return
}

func (model *Model) getUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := r.Header.Get("x-token")
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s", "Unexpected signing method")

		} else {
			keyData, err := ioutil.ReadFile(os.Getenv("KEY_PATH"))
			if err != nil {
				return nil, err
			}
			return keyData, nil
		}
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := parsedToken.Claims.(jwt.MapClaims)
	user, err := db.Get(model.db, token["sub"].(string))
	if err != nil {
		log.Println(err)
		return
	}

	res, err := json.Marshal(map[string]string{
		"name": user.UserName,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
	return
}

func (model *Model) updateUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := r.Header.Get("x-token")
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s", "Unexpected signing method")

		} else {
			keyData, err := ioutil.ReadFile(os.Getenv("KEY_PATH"))
			if err != nil {
				return nil, err
			}
			return keyData, nil
		}
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokenMap := parsedToken.Claims.(jwt.MapClaims)

	var accountInfo db.User
	accountInfo.UserId = tokenMap["sub"].(string)
	if accountInfo.IsRecord(model.db) {
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
	updateTimeUTC := time.Now().UTC()
	jst, _ := time.LoadLocation("Asia/Tokyo")
	updateTimeJST := updateTimeUTC.In(jst)
	updateInfo.UpdateTime = updateTimeUTC
	updateInfo.UpdateTimeJST = updateTimeJST

	if err := db.Update(model.db, updateInfo); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func (router *Server) Init() error {
	// db.open の処理
	connectedDB, err := db.Open()
	if err != nil {
		return err
	}

	model := Model{db: connectedDB}

	// http method ごとの処理(handler)
	router.Engine.HandleFunc("/user/create", model.createUserHandler)
	router.Engine.HandleFunc("/user/get", model.getUserHandler)
	router.Engine.HandleFunc("/user/update", model.updateUserHandler)

	return nil
}

func (router *Server) Run(port string) {
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router.Engine)
	if err != nil {
		log.Fatal(err)
	}
}

func createUserId() string {
	userId := uuid.Must(uuid.NewRandom())
	return strings.ReplaceAll(userId.String(), "-", "")
}

func main() {

	server := NewServer()
	if err := server.Init(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(port)
}
