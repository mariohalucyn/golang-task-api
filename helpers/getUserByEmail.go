package helpers

import (
	"github.com/mariohalucyn/todo-app/initializers"
	"github.com/mariohalucyn/todo-app/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	tx := initializers.DB.First(&user, "email_address = ?", email)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
