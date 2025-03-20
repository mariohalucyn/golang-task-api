package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mariohalucyn/todo-app/helpers"
	"github.com/mariohalucyn/todo-app/initializers"
	"github.com/mariohalucyn/todo-app/models"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	var user models.User

	claims, status, err := helpers.Authorization(r)
	if err != nil {
		helpers.WriteJson(w, status, &helpers.ApiResponse{Message: err.Error()})
		return
	}

	if err := helpers.GetUserByClaimsIssuer(claims.Issuer, &user); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: err.Error()})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		helpers.WriteJson(w, http.StatusBadRequest, helpers.ApiResponse{Message: "Failed to decode request body."})
		return
	}

	if err := createNewTodo(todo, user.ID); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: "Failed to create Todo."})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func createNewTodo(todo models.Todo, userID uint) error {
	if todo.Completed == nil {
		b := false
		todo.Completed = &b
	}

	tx := initializers.DB.Create(&models.Todo{
		Task:      todo.Task,
		DueDate:   todo.DueDate,
		Completed: todo.Completed,
		UserID:    userID,
	})
	return tx.Error
}
