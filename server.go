package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/xid"

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
	return xid.New().String()
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
