package helpers

import (
	"github.com/mariohalucyn/todo-app/initializers"
	"github.com/mariohalucyn/todo-app/models"
)

func GetTodoByID(id string) (*models.Todo, error) {
	var todo models.Todo
	tx := initializers.DB.First(&todo, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &todo, nil
}
