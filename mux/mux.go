package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Director *Director `json:"director"`
}

type Director struct {
	Name string `json:"director"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	var m Movie
	json.NewDecoder(r.Body).Decode(&m)
	fmt.Fprintf(w, "movie: %+v", m)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "get movie\n")
	for _, ele := range movies {
		if strconv.Itoa(ele.Id) == vars["id"] {
			json.NewEncoder(w).Encode(ele)
			// fmt.Fprint(w, )
		}
	}
}

var movies []Movie

func main() {
	movies = append(movies, Movie{1001, "Rush", &Director{"Chan"}})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", addMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
