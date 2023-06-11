package createMovie

import (
	"encoding/json"
	"github.com/nitishfy/go-movies-CRUD/pkg/types"
	"math/rand"
	"net/http"
	"strconv"
)

// CreateMovie creates a movie, data will be sent in JSON format in request body(r.Body)
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	var movie types.Movies                     // this is the movie you'll create
	_ = json.NewDecoder(r.Body).Decode(&movie) // Decode the JSON format
	movie.ID = strconv.Itoa(rand.Intn(100000)) // Generating a random ID for the movie in string format
	types.MoviesList = append(types.MoviesList, movie)

	// send a response back (movie created) in the JSON format
	err := json.NewEncoder(w).Encode(movie)
	if err != nil {
		return
	}
}
