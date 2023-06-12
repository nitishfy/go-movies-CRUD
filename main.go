package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nitishfy/go-movies-CRUD/pkg/handlers/createMovie"
	"github.com/nitishfy/go-movies-CRUD/pkg/handlers/deleteMovie"
	"github.com/nitishfy/go-movies-CRUD/pkg/handlers/getMovie"
	"github.com/nitishfy/go-movies-CRUD/pkg/handlers/getMovies"
	"github.com/nitishfy/go-movies-CRUD/pkg/handlers/updateMovie"
	"github.com/nitishfy/go-movies-CRUD/pkg/types"
	"net/http"
)

func main() {

	// We will display some Movies at the initial when user hits the endpoint /Movies
	types.MoviesList = append(types.MoviesList, types.Movies{
		ID:    "1",
		Isbn:  "4342",
		Title: "Movie One",
		Director: &types.Director{
			Firstname: "John",
			Lastname:  "Maverick",
		},
	})

	types.MoviesList = append(types.MoviesList, types.Movies{
		ID:    "2",
		Isbn:  "4341",
		Title: "Movie Two",
		Director: &types.Director{
			Firstname: "Alex",
			Lastname:  "Hills",
		},
	})
	//	create a new router
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie.GetMovieById).Methods("GET")
	r.HandleFunc("/movies", createMovie.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie.DeleteMovie).Methods("DELETE")

	fmt.Printf("starting server on port %s", types.PortNum)
	if err := http.ListenAndServe(types.PortNum, r); err != nil {
		return
	}
}
