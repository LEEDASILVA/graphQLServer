package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("secret")

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// store the claims
	claims := token.Claims.(jwt.MapClaims)

	// set Tokens
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tokenString, nil
}

// parseToken parses a jwt token and returns the username
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		log.Panic(err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", nil
	}
}
