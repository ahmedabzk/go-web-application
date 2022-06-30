package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ahmed/go-web/helpers"
	"github.com/ahmed/go-web/middleware"
	"github.com/ahmed/go-web/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// SignUp function
func SignUp(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var user models.User
	_ = json.NewDecoder(req.Body).Decode(&user)
	user.Password = middleware.GetHash([]byte(user.Password))
	result, err := helpers.Collection.InsertOne(context.Background(), user)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "` + err.Error() + `"}`))
		if err != nil {
			log.Fatal(err)
			return
		}
		return
	}
	_ = json.NewEncoder(res).Encode(result)
}

//Login function
func Login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var user models.User
	var dbUser models.User
	_ = json.NewDecoder(req.Body).Decode(&user)

	filter := bson.M{"email": user.Email}
	err := helpers.Collection.FindOne(context.Background(), filter).Decode(&dbUser)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "` + err.Error() + `"}`))
	}

	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		res.Write([]byte(`{"response": "wrong email or password"}`))
		return
	}
	// generate the token
	jwtToken, err := helpers.GenerateJwt()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	res.Write([]byte(`{"token": "` + jwtToken + `"}`))

}
