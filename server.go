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

	"github.com/shinnosuke-K/Tech-Train-CA-Go/controller"
)

type Server struct {
	Engine *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		Engine: http.NewServeMux(),
	}
}

func (router *Server) Init() error {
	// db.open の処理
	connectedDB, err := db.Open()
	if err != nil {
		return err
	}

	ctr := controller.New(connectedDB)

	// http method ごとの処理(handler)
	router.Engine.HandleFunc("/user/create", ctr.CreateUserHandler)
	router.Engine.HandleFunc("/user/get", ctr.GetUserHandler)
	router.Engine.HandleFunc("/user/update", ctr.UpdateUserHandler)

	return nil
}

func (router *Server) Run(port string) {
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router.Engine)
	if err != nil {
		log.Fatal(err)
	}
}

var DB, _ = db.Open()

func initUserHandler() handler.UserHandler {
	userPersistence := persistence.NewUserPersistence(DB)
	userUseCase := usecase.NewUserUseCase(userPersistence)
	return handler.NewUserHandler(userUseCase)
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
