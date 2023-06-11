package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
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
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

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

// getMovie displays a movie in the JSON format
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// createMovie creates a movie, data will be sent in JSON format in request body(r.Body)
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	var movie Movies                           // this is the movie you'll create
	_ = json.NewDecoder(r.Body).Decode(&movie) // Decode the JSON format
	movie.ID = strconv.Itoa(rand.Intn(100000)) // Generating a random ID for the movie in string format
	movies = append(movies, movie)

	// send a response back (movie created) in the JSON format
	err := json.NewEncoder(w).Encode(movie)
	if err != nil {
		return
	}
}

// updateMovie updates the existing movie (delete the existing movie and then create a new movie with same id)
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set the content type as JSON
	w.Header().Set("Content-Type", "json")
	params := mux.Vars(r)

	// Delete the movie
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			// create a new movie that will be replaced with deleted movie
			var movie Movies
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = item.ID
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

// deleteMovie deletes a movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	// return the remaining movies in JSON format after deleting
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}
