package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/mariohalucyn/todo-app/helpers"
	"github.com/mariohalucyn/todo-app/initializers"
	"github.com/mariohalucyn/todo-app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, &helpers.ApiResponse{Message: "Failed to decode request body."})
		return
	}

	tx := initializers.DB.First(&user, "email_address = ? AND is_verified = ?", loginRequest.EmailAddress, true)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			helpers.WriteJson(w, http.StatusBadRequest, helpers.ApiResponse{Message: "Enter a valid email address."})
		} else {
			helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: "Error retrieving user."})
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		helpers.WriteJson(w, http.StatusBadRequest, &helpers.ApiResponse{Message: "Your account or password is incorrect."})
		return
	}

	cookieExpTime := getCookieExpirationTime(loginRequest.StaySignedIn)
	tokenString, err := helpers.GenerateJWT(user)
	if err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.ApiResponse{Message: "Failed to generate token."})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Path:     "/",
		Expires:  cookieExpTime,
		Secure:   true,
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
}

func getCookieExpirationTime(staySignedIn bool) time.Time {
	if staySignedIn {
		return time.Now().AddDate(0, 1, 0)
	}
	return time.Time{}
}
