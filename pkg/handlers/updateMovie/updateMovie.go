package updateMovie

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nitishfy/go-movies-CRUD/pkg/types"
	"net/http"
)

// UpdateMovie updates the existing movie (delete the existing movie and then create a new movie with same id)
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	// set the content type as JSON
	w.Header().Set("Content-Type", "json")
	params := mux.Vars(r)

	// Delete the movie
	for index, item := range types.MoviesList {
		if item.ID == params["id"] {
			types.MoviesList = append(types.MoviesList[:index], types.MoviesList[index+1:]...)

			// create a new movie that will be replaced with deleted movie
			var movie types.Movies
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = item.ID
			types.MoviesList = append(types.MoviesList, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}
