package getMovie

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nitishfy/go-movies-CRUD/pkg/types"
	"net/http"
)

// GetMovieById displays a movie in the JSON format
func GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	params := mux.Vars(r)
	for _, item := range types.MoviesList {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
