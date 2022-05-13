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
	Name string `json:"name"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	var m Movie
	json.NewDecoder(r.Body).Decode(&m)
	movies = append(movies, m)
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

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "get movie\n")
	for idx, ele := range movies {
		if strconv.Itoa(ele.Id) == vars["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
		}
	}
}

var movies []Movie

func main() {
	movies = append(movies, Movie{1001, "Rush", &Director{"Chan"}})
	movies = append(movies, Movie{1002, "Panda", &Director{"Lem"}})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", addMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	http.ListenAndServe(":2022", r)
}
