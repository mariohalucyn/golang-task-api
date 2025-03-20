package helpers

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/mariohalucyn/todo-app/initializers"
	"net/http"
)

func Authorization(r *http.Request) (claims *jwt.StandardClaims, status int, err error) {
	signedString, err := r.Cookie("Authorization")
	if err != nil {
		return nil, http.StatusUnauthorized, err
	}

	token, err := jwt.ParseWithClaims(signedString.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("unexpected signing method: expected ECDSA")
		}
		return initializers.EcdsaPublicKey, nil
	})
	if err != nil {
		return nil, http.StatusUnauthorized, err
	}

	claims = token.Claims.(*jwt.StandardClaims)

	if err := claims.Valid(); err != nil {
		return nil, http.StatusUnauthorized, err
	}

	return claims, http.StatusOK, nil
}
