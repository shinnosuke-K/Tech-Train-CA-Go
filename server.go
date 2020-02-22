package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/xid"

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

func (router *Server) Init() error {
	// db.open の処理
	// http method ごとの処理(handler)

	return nil
}

func (router *Server) Run(port string) {
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router.Engine)
	if err != nil {
		log.Fatal(err)
	}
}

func tryUuidGoogle() {
	userId := uuid.Must(uuid.NewRandom())
	fmt.Printf("%s : %d\n", userId, len(userId.String()))
}

func tryUuidXid() {
	userId := xid.New()
	fmt.Printf("%s : %d\n", userId.String(), len(userId.String()))
}

//func tryUuidAnarchar() {
//	userId := shortuuid.New()
//	fmt.Printf("%s : %d\n", userId, len(userId.String()))
//}

// 使うid生成パッケージ github.com/rs/xid
func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Google/uuid")
		tryUuidGoogle()
		fmt.Println("xid/uuid")
		tryUuidXid()
		//fmt.Println("anarchar")
		//tryUuidAnarchar()
	}

	server := NewServer()
	if err := server.Init(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	//server.Run(port)
}
