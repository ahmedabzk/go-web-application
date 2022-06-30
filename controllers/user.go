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

var userCol = "users"

// SignUp function
func SignUp(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	client, err := helpers.Setup()
	collection := client.Database(helpers.DbName).Collection(userCol)
	var user models.User
	_ = json.NewDecoder(req.Body).Decode(&user)
	user.Password = middleware.GetHash([]byte(user.Password))
	result, er := collection.InsertOne(context.Background(), user)

	if er != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "` + er.Error() + `"}`))
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

	client, err := helpers.Setup()
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(helpers.DbName).Collection(userCol)
	var user models.User
	var dbUser models.User
	_ = json.NewDecoder(req.Body).Decode(&user)

	filter := bson.M{"email": user.Email}

	er := collection.FindOne(context.Background(), filter).Decode(&dbUser)

	if er != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "` + er.Error() + `"}`))
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
	jwtToken, errs := helpers.GenerateJwt()

	if errs != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "` + errs.Error() + `"}`))
		return
	}
	res.Write([]byte(`{"token": "` + jwtToken + `"}`))

}
