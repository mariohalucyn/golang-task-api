package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Task      string    `json:"task"`
	DueDate   time.Time `json:"dueDate"`
	Completed *bool     `json:"completed"`
	UserID    uint      `json:"userId"`
}
