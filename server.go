package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/Infra/persistence"
	handler "github.com/shinnosuke-K/Tech-Train-CA-Go/handler/api"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/handler/db"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/usecase"
)

type Server struct {
	Engine *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		Engine: http.NewServeMux(),
	}
}

var DB, _ = db.Open()

func initUserHandler() handler.UserHandler {
	userPersistence := persistence.NewUserPersistence(DB)
	userUseCase := usecase.NewUserUseCase(userPersistence)
	return handler.NewUserHandler(userUseCase)
}

func (router *Server) Init() {

	userHandler := initUserHandler()
	router.Engine.HandleFunc("/user/create", userHandler.Create)
	router.Engine.HandleFunc("/user/get", userHandler.Get)

}

func (router *Server) Run(port string) {
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router.Engine)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	server := NewServer()
	server.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(port)
}
