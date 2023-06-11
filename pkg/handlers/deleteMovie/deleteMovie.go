package deleteMovie

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nitishfy/go-movies-CRUD/pkg/types"
	"net/http"
)

// DeleteMovie deletes a movie
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	params := mux.Vars(r)
	for index, item := range types.MoviesList {
		if item.ID == params["id"] {
			types.MoviesList = append(types.MoviesList[:index], types.MoviesList[index+1:]...)
			break
		}
	}
	// return the remaining Movies in JSON format after deleting
	err := json.NewEncoder(w).Encode(types.MoviesList)
	if err != nil {
		return
	}
}
