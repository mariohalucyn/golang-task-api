package handlers

import (
	"github.com/gorilla/mux"
	"github.com/mariohalucyn/todo-app/helpers"
	"github.com/mariohalucyn/todo-app/initializers"
	"github.com/mariohalucyn/todo-app/models"
	"net/http"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todo models.Todo

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

	if err := deleteTodo(todo, vars["id"]); err != nil {
		helpers.WriteJson(w, status, helpers.ApiResponse{Message: err.Error()})
	}

	w.WriteHeader(http.StatusOK)
}

func deleteTodo(todo models.Todo, id string) error {
	tx := initializers.DB.Delete(&todo, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
