package utils

import (
	"fmt"

	"time"

	"github.com/golang-jwt/jwt/v4"
)

const SECRET_KEY = "janganlupabobok"

func GenerateJWT(username string) (string, error) {

	var mySigningKey = []byte(SECRET_KEY)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
