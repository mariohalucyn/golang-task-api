package handlers

import (
	"net/http"

	"github.com/mariohalucyn/todo-app/helpers"
	"github.com/mariohalucyn/todo-app/initializers"
	"github.com/mariohalucyn/todo-app/models"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var todos []models.Todo

	claims, status, err := helpers.Authorization(r)
	if err != nil {
		helpers.WriteJson(w, status, helpers.ApiResponse{Message: err.Error()})
		return
	}

	if err := helpers.GetUserByClaimsIssuer(claims.Issuer, &user); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: err.Error()})
		return
	}

	if err := getTodosByUserID(user.ID, &todos); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: err.Error()})
		return
	}

	helpers.WriteJson(w, http.StatusOK, todos)
}

func getTodosByUserID(userID uint, todos *[]models.Todo) error {
	tx := initializers.DB.Find(todos, "user_id = ?", userID)
	return tx.Error
}
