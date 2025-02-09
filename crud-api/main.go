package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


type Movie struct{
	ID string `json:"id"`
	Isbn string  `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
 Firstname string `json:"firstname"`
 Lasname string `json:"lastname"`
}

var movies []Movie

func getMovies(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func getMovie(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range movies{
		if item.ID == params["id"]{
		  json.NewEncoder(res).Encode(item)
		  return
		}
	}
}

func createMovie(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	var movie Movie

	_ = json.NewDecoder(req.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000000000))

	movies = append(movies, movie)

	json.NewEncoder(res).Encode(movie)
}

func updateMovie(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range movies{
		if item.ID == params["id"]{
		 movies = append(movies [:index], movies[index+1:]...)
		 var movie Movie

		 _ = json.NewDecoder(req.Body).Decode(&movie)
	 
		 movie.ID =  params["id"]
	 
		 movies = append(movies, movie)

		 json.NewEncoder(res).Encode(movies)

		 break
		}
	}
}

func deleteMovie(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range movies{

		if item.ID == params["id"]{
          movies = append(movies [:index], movies[index+1:]...)
		  break
		}
	}

	json.NewEncoder(res).Encode(movies)
}


func main() {
  router := mux.NewRouter()

  movies = append(movies, Movie{ID: "1", Isbn: "438526", Title: "Movie one", Director: &Director{Firstname: "John", Lasname: "Doe"}  })

  router.HandleFunc("/movies", getMovies).Methods("GET")
  router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
  router.HandleFunc("/movies", createMovie).Methods("POST")
  router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
  router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")


  fmt.Printf("Starting server at port 8000\n")
  
  error := http.ListenAndServe(":8000", router)

  if error != nil {
	  log.Fatal(error)
  }
}