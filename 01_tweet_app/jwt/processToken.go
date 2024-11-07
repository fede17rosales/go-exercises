package jwt

import (
	"errors"
	"strings"
	"twittergo/models"

	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUser string

func ProcessToken(tk string, JWTSing string) (*models.Claim, bool, string, error) {
	myKey := []byte(JWTSing)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		// check runtine
	}
	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token Invalido")
	}

	return &claims, false, string(""), err

}
