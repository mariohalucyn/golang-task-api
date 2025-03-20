package handlers

import (
	"github.com/mariohalucyn/todo-app/helpers"
	"github.com/mariohalucyn/todo-app/initializers"
	"github.com/mariohalucyn/todo-app/models"
	"net/http"
	"os"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	tx := initializers.DB.Model(models.User{}).Where("verification_token = ?", token).Updates(map[string]interface{}{
		"is_verified":        true,
		"verification_token": nil,
	})

	if tx.Error != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: tx.Error.Error()})
		return
	}

	http.Redirect(w, r, os.Getenv("FRONTEND_ADDRESS"), http.StatusFound)
}
