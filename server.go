package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/handler"
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

func initUserHandler(DB *sql.DB, tx db.Transaction) handler.UserHandler {
	userPersistence := persistence.NewUserPersistence(DB)
	userUseCase := usecase.NewUserUseCase(userPersistence, tx)
	return handler.NewUserHandler(userUseCase)
}

func initGachaHandler(DB *sql.DB, tx db.Transaction) handler.GachaHandler {
	gachaPersistence := persistence.NewGachaPersistence(DB)
	gachaUseCase := usecase.NewGachaUseCase(gachaPersistence, tx)
	return handler.NewGachaHandler(gachaUseCase)
}

func initCharaHandler(DB *sql.DB) handler.CharacterHandler {
	charaPersistence := persistence.NewCharaPersistence(DB)
	charaUseCase := usecase.NewCharaUseCase(charaPersistence)
	return handler.NewCharaHandler(charaUseCase)
}

func (router *Server) Init(DB *sql.DB) {

	tx := db.NewTransaction(DB)

	userHandler := initUserHandler(DB, tx)
	router.Engine.HandleFunc("/user/create", userHandler.Create)
	router.Engine.HandleFunc("/user/get", userHandler.Get)
	router.Engine.HandleFunc("/user/update", userHandler.Update)

	gachaHandler := initGachaHandler(DB, tx)
	router.Engine.HandleFunc("/gacha/draw", gachaHandler.Draw)

	charaHandler := initCharaHandler(DB)
	router.Engine.HandleFunc("/character/list", charaHandler.List)
}

func (router *Server) Run(port string) {
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router.Engine)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	DB, err := db.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer DB.Close()

	rand.Seed(time.Now().UnixNano())

	server := NewServer()
	server.Init(DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(port)
}
