package helpers

import (
	"github.com/mariohalucyn/todo-app/initializers"
	"github.com/mariohalucyn/todo-app/models"
)

func GetUserByClaimsIssuer(issuer string, user *models.User) error {
	tx := initializers.DB.First(&user, "email_address = ?", issuer)
	return tx.Error
}
