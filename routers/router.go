package router

import (
	MovieController "github.com/ahmed/go-web/controllers"
	UserController "github.com/ahmed/go-web/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	//user router
	router.HandleFunc("/users/signup", UserController.SignUp).Methods("POST")
	router.HandleFunc("/users/login", UserController.Login).Methods("POST")

	//movie router
	router.HandleFunc("/users/search", MovieController.SearchAMovieById).Methods("POST")
	router.HandleFunc("/users/create", MovieController.InsertAMovie).Methods("POST")
	router.HandleFunc("/users/displayAll", MovieController.ShowAllMovies).Methods("GET")
	router.HandleFunc("/users/deleteOne", MovieController.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/users/deleteAll", MovieController.DeleteAllMovies).Methods("DELETE")
	return router
}
