// This controller responsible for admin user login
package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/RaguSoftEng/inviteservice_api/src/models"
	"github.com/RaguSoftEng/inviteservice_api/src/repository"
	"github.com/RaguSoftEng/inviteservice_api/util"
)

// Validate Admin user and generate jwt token
// parameters : username string, password string
// return jwt token
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		util.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := repository.ValidateUser(server.DB, user.Username, user.Password)
	if err != nil {
		util.ERROR(w, http.StatusUnprocessableEntity, errors.New("[ ERROR ] Username or Password incorrect."))
		return
	}

	util.JSON(w, http.StatusOK, token)
}
