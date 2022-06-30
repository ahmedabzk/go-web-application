package controllers

import (
	"encoding/json"
	"github.com/ahmed/go-web/helpers"
	"github.com/ahmed/go-web/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InsertAMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Methods", "POST")
	var movie models.Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	helpers.InsertOneMovie(movie)
	err := json.NewEncoder(res).Encode(movie)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func SearchAMovieById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Methods", "GET")
	params := mux.Vars(req)
	helpers.SearchForAMovie(params["id"])

	_ = json.NewEncoder(res).Encode(params["id"])
}

func DeleteOneMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(req)

	//pass the movie id to the delete function
	helpers.DeleteOneMovie(params["id"])

}

func DeleteAllMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Methods", "DELETE")
	//return the count of deleted movies
	count := helpers.DeleteAllMovies
	_ = json.NewEncoder(res).Encode(count)

}

func ShowAllMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Methods", "GET")
	movies := helpers.DisplayAllMovies()
	_ = json.NewEncoder(res).Encode(movies)
}
