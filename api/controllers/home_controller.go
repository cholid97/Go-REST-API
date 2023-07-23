package controllers

import (
	"net/http"

	"example/web-service-gin/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To this API")
}
