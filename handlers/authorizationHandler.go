package handlers

import (
	"github.com/mariohalucyn/todo-app/helpers"
	"net/http"
)

func Authorization(w http.ResponseWriter, r *http.Request) {
	_, status, err := helpers.Authorization(r)
	if err != nil {
		helpers.WriteJson(w, status, helpers.ApiResponse{Message: err.Error()})
		return
	}
}
