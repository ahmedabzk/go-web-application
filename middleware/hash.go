package middleware

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)
// hashing function
func GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil{
		log.Fatal(err)
	}
	return string(hash)
}
