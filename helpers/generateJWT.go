package helpers

import (
	"github.com/golang-jwt/jwt"
	"github.com/mariohalucyn/todo-app/initializers"
	"github.com/mariohalucyn/todo-app/models"
	"strconv"
	"time"
)

func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"jti": strconv.Itoa(int(user.ID)),
		"iss": user.EmailAddress,
		"iat": time.Now().Unix(),
		"exp": time.Now().AddDate(0, 1, 0).Unix(),
	})
	return token.SignedString(initializers.EcdsaPrivateKey)
}
