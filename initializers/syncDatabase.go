package initializers

import (
	"log"

	"github.com/mariohalucyn/todo-app/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatal(err)
	}
}
