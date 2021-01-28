package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	handler "github.com/shinnosuke-K/Tech-Train-CA-Go/handler/api"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/db"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/persistence"
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

func initUserHandler(db *sql.DB) handler.UserHandler {
	userPersistence := persistence.NewUserPersistence(db)
	userUseCase := usecase.NewUserUseCase(userPersistence)
	return handler.NewUserHandler(userUseCase)
}

func initGachaHander(db *sql.DB) handler.GachaHandler {
	gachaPersistence := persistence.NewGachaPersistence(db)
	gachaUseCase := usecase.NewGachaUseCase(gachaPersistence)
	return handler.NewGachaHandler(gachaUseCase)
}

func (router *Server) Init(db *sql.DB) {

	userHandler := initUserHandler(db)
	router.Engine.HandleFunc("/user/create", userHandler.Create)
	router.Engine.HandleFunc("/user/get", userHandler.Get)
	router.Engine.HandleFunc("/user/update", userHandler.Update)

	gachaHandler := initGachaHander(db)
	router.Engine.HandleFunc("/gacha/draw", gachaHandler.Draw)
}

func (router *Server) Run(port string) {
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router.Engine)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := NewServer()
	server.Init(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(port)
}
