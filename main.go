package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)

	for index, item := range movies {
		if item.ID == parms["id"] {
			movies = append(movies[:index], movies[index+1:]...)
		}
		break
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	for _, item := range movies {
		if item.ID == parms["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// parms := mux.Vars(r)
	// movie := Movie{
	// 	ID:    strconv.Itoa(len(movies)),
	// 	Isbn:  parms["isbn"],
	// 	Title: parms["title"],
	// 	Director: &Director{
	// 		Firstname: parms["firstname"],
	// 		Lastname:  parms["lastname"],
	// 	},
	// }
	// movies = append(movies, movie)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(len(movies))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	for index, item := range movies {
		if item.ID == parms["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			movie.ID = item.ID
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}

}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "1234", Title: "First Movie", Director: &Director{Firstname: "Ravi", Lastname: "Kumar"}})
	movies = append(movies, Movie{ID: "2", Isbn: "1235", Title: "Movie Two", Director: &Director{Firstname: "Khuss", Lastname: "Prajapati"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
