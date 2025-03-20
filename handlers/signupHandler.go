package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/mariohalucyn/todo-app/helpers"
	"github.com/mariohalucyn/todo-app/initializers"
	"github.com/mariohalucyn/todo-app/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

const BcryptCost = 10

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: "Failed to decode request body."})
		return
	}

	tx := initializers.DB.First(&user, "email_address = ? AND is_verified = ?", user.EmailAddress, true)
	if tx.Error == nil {
		helpers.WriteJson(w, http.StatusBadRequest, helpers.ApiResponse{Message: "This email belongs to an existing account."})
	} else if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: tx.Error.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), BcryptCost)
	if err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: "Failed to hash password."})
		return
	}

	tokenString, err := helpers.GenerateJWT(user)
	if err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: "Failed to generate token."})
		return
	}

	newUser := models.User{
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		EmailAddress:      user.EmailAddress,
		Password:          string(hashedPassword),
		VerificationToken: tokenString,
	}

	if err := createUser(newUser); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: "Failed to create user."})
		return
	}

	if err := sendVerificationEmail(user.EmailAddress, tokenString); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: "Failed to send verification email."})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func createUser(user models.User) error {
	tx := initializers.DB.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func sendVerificationEmail(email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("FROM_EMAIL"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email Verification")
	m.SetHeader("X-PM-Message-Stream", "verification")
	m.SetBody("text/html", fmt.Sprintf(`<p>Click <a href="http://localhost:8000/api/verify?token=%s">here</a> to verify your email.</p>`, token))

	d := gomail.NewDialer("smtp.postmarkapp.com", 587, os.Getenv("POSTMARKAPP_USERNAME"), os.Getenv("POSTMARKAPP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
