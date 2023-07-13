package main

import{

	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

}

type Movie struct{
	ID string `json:"id"`
	isbn string `json:"isbn"`
	Title string `json:"Title"`
	Director *Director `json:"director"`

}

type Director struct{

	Firstname string `json:"fname"`
	Lastname string `json:"lname"`

}

var movies []Movie

func main(){
	r:=MUX.NewRouter()

	movies = append(movies, Movie[ID: "1", Isbn: "573839", Title: "Movie1", Director: &Director(fname:"james", lname:"cameron")])
	movies = append(movies, Movie[ID: "2", Isbn: "45829", Title: "Movie2", Director: &Director(fname:"john", lname:"paul")])


	r.HandleFunc("/movies",getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getmovie).Methods("GET")
	r.HandleFunc("/movies",createmovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updatemovie).Methods("POST")
	r.HandleFunc("/movies/{id}",deletemovie).Methods("POST")
}