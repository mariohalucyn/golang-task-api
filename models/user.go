package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	EmailAddress      string `json:"email"`
	Password          string `json:"password"`
	Todos             []Todo `gorm:"foreignKey:UserID"`
	IsVerified        bool   `json:"isVerified" gorm:"default:false"`
	VerificationToken string `json:"verificationToken"`
}

type LoginRequest struct {
	EmailAddress string `json:"email"`
	Password     string `json:"password"`
	StaySignedIn bool   `json:"staySignedIn"`
}
