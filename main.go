package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Movies struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movies

const portNum = ":8000"

func main() {
	//	create a new router
	r := mux.NewRouter()

	// We will display some movies at the initial when user hits the endpoint /movies
	movies = append(movies, Movies{
		ID:    "1",
		Isbn:  "4342",
		Title: "Movie One",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Maverick",
		},
	})

	movies = append(movies, Movies{
		ID:    "2",
		Isbn:  "4341",
		Title: "Movie Two",
		Director: &Director{
			Firstname: "Alex",
			Lastname:  "Hills",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	//r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	//r.HandleFunc("/movies", createMovie).Methods("POST")
	//r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	//r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server on port %s", portNum)
	if err := http.ListenAndServe(portNum, r); err != nil {
		return
	}
}

// getMovies will display all the movies in the JSON format
func getMovies(w http.ResponseWriter, r *http.Request) {
	// set the response content to JSON format
	w.Header().Set("Content-Type", "json")

	// create a new JSON encoder that will encode slice of movies
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}
