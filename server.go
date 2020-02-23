package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
	if r.Method == http.MethodPost {
		response := map[string]string{
			"name": createUserId(),
		}
		resJson, err := json.Marshal(response)
		if err != nil {
			log.Fatal()
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resJson)
	} else {

	}
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
