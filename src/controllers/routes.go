package controllers

import (
	"github.com/RaguSoftEng/inviteservice_api/src/middlewares"
	"github.com/didip/tollbooth"
)

func (server *Server) initializeRoutes() {
	// Home rout
	server.Router.HandleFunc("/", middlewares.PrepareJSON(server.Home)).Methods("GET")

	// Admin login rout
	server.Router.HandleFunc("/login", middlewares.PrepareJSON(server.Login)).Methods("POST")

	// Generate invite token rout
	server.Router.HandleFunc("/invite", middlewares.PrepareJSON(middlewares.SetAuthentication(server.GenerateToken))).Methods("POST")

	// Get all tokens by admin rout
	server.Router.HandleFunc("/invite", middlewares.PrepareJSON(middlewares.SetAuthentication(server.GetTokens))).Methods("GET")

	// Disable token by admin
	server.Router.HandleFunc("/invite/{token}", middlewares.PrepareJSON(middlewares.SetAuthentication(server.DisableToken))).Methods("PUT")

	// Guest login using invite token with request limit
	limiter := tollbooth.NewLimiter(1, nil)

	server.Router.Handle("/invite/{token}", tollbooth.LimitFuncHandler(limiter, server.ValidateToken)).Methods("GET")
}
