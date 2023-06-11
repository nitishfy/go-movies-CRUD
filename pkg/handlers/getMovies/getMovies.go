package getMovies

import (
	"encoding/json"
	"github.com/nitishfy/go-movies-CRUD/pkg/types"
	"net/http"
)

// GetMovies will display all the movies in the JSON format
func GetMovies(w http.ResponseWriter, r *http.Request) {
	// set the response content to JSON format
	w.Header().Set("Content-Type", "json")

	// create a new JSON encoder that will encode slice of movies
	err := json.NewEncoder(w).Encode(types.MoviesList)
	if err != nil {
		return
	}
}
