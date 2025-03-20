package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mariohalucyn/todo-app/helpers"
	"github.com/mariohalucyn/todo-app/initializers"
	"github.com/mariohalucyn/todo-app/models"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: err.Error()})
		return
	}

	claims, status, err := helpers.Authorization(r)
	if err != nil {
		helpers.WriteJson(w, status, helpers.ApiResponse{Message: err.Error()})
		return
	}

	if err := updateUser(user, claims.Id); err != nil {
		helpers.WriteJson(w, status, helpers.ApiResponse{Message: err.Error()})
		return
	}
}

func updateUser(user models.User, id string) error {
	updatedUser := map[string]interface{}{}

	if user.FirstName != "" {
		updatedUser["first_name"] = user.FirstName
	}

	if user.LastName != "" {
		updatedUser["last_name"] = user.LastName
	}

	if user.Password != "" {
		updatedUser["password"] = user.Password
	}

	tx := initializers.DB.Model(models.User{}).Where("id = ?", id).Updates(updatedUser)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
