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

	"github.com/dgrijalva/jwt-go"

	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
)

type Server struct {
	db     *gorm.DB
	Engine *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		Engine: http.NewServeMux(),
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {

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

	createTimeUTC := time.Now().UTC()
	//jst, _ := time.LoadLocation("Asia/Tokyo")
	//createTimeJST := createTimeUTC.In(jst)
	userId := createUserId()
	name := jsonBody["name"]
	if name == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userId,
		"user": name,
		"nbf":  createTimeUTC,
		"iat":  createTimeUTC,
	})

	keyData, err := ioutil.ReadFile("./rsa/jwtRS256.key")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokenString, err := token.SignedString(keyData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(map[string]string{
		"token": tokenString,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
	return
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {

}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (router *Server) Init() error {
	// db.open の処理

	// http method ごとの処理(handler)
	router.Engine.HandleFunc("/user/create", createUserHandler)
	router.Engine.HandleFunc("/user/get", getUserHandler)
	router.Engine.HandleFunc("/user/update", updateUserHandler)

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
