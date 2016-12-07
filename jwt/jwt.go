package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	"errors"
	"github.com/smhouse/pi/db"
)

//CreateToken creates new access token for specified user
func CreateToken(tokenUser *db.User_t, secret string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": tokenUser.Name,
		"email": tokenUser.Email,
		"nbf": time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

//CheckToken verifies token and returns user
func CheckToken(tokenString string, secret string) (*db.User_t, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		rUser := db.User_t{
			Email:		claims["email"].(string),
			Password:	"",
			Name:		claims["name"].(string),
		}

		return &rUser, nil
	}

	return nil, errors.New("Unable to validate token")
}
