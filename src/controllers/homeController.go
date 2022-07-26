package controllers

import (
	"net/http"

	"github.com/RaguSoftEng/inviteservice_api/util"
)

// Home or default controller
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	util.JSON(w, http.StatusOK, "Welcome To Catalyst Experience.")

}
