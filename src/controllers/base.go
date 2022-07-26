package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/RaguSoftEng/inviteservice_api/util"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {

	server.DB = util.Connect()

	util.Migrate()

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

// Initiate server
func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
