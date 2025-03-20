package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mariohalucyn/todo-app/helpers"
	"github.com/mariohalucyn/todo-app/initializers"
	"github.com/mariohalucyn/todo-app/models"
)

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: err.Error()})
		return
	}

	claims, status, err := helpers.Authorization(r)
	if err != nil {
		helpers.WriteJson(w, status, helpers.ApiResponse{Message: err.Error()})
		return
	}

	user, err := helpers.GetUserByEmail(claims.Issuer)
	if err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: err.Error()})
		return
	}

	todoFromDB, err := helpers.GetTodoByID(vars["id"])
	if err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: err.Error()})
		return
	}

	if user.ID != todoFromDB.UserID {
		helpers.WriteJson(w, http.StatusBadRequest, helpers.ApiResponse{Message: "You are not authorized to perform this action"})
		return
	}

	if err := updateTodo(todo, vars["id"]); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateTodo(todo models.Todo, id string) error {
	updatedTodo := map[string]interface{}{}
	var t time.Time

	if todo.Task != "" {
		updatedTodo["task"] = todo.Task
	}

	if todo.Completed != nil {
		updatedTodo["completed"] = todo.Completed
	}

	if todo.DueDate != t {
		updatedTodo["due_date"] = todo.DueDate
	}

	tx := initializers.DB.Model(models.Todo{}).Where("id = ?", id).Updates(updatedTodo)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
