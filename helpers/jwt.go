package helpers

import (
	"log"
	"os"

	// "os"
    "github.com/dgrijalva/jwt-go"
	// "github.com/joho/godotenv"
)

var SECRET_KEY string

func GenerateJwt() (string, error) {
	if SECRET_KEY := os.Getenv("SECRET_KEY"); SECRET_KEY == "" {
		log.Fatal("environment variable not provided\n")
	}

	key := []byte(SECRET_KEY)
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(key)

	if err != nil {
		log.Println("error in jwt generation")
		return "", err
	}
	return tokenString, nil
}
