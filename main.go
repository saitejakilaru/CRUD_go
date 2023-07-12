package main

import(

	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Moviename string `json:"Moviename"`
	Director *Director `json:"directorname"`

}

type Director struct{

	Fname string `json:"fname"`
	Lname string `json:"lname"`

}

var movielist []Movie

func getAllMovies(wr http.ResponseWriter, rd *http.Request){
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(movielist)
}

func deleteMovieName(wr http.ResponseWriter, rd *http.Request){
	wr.Header().Set("Content-Type", "application/json")
	param := mux.Vars(rd)
	for i,item:= range movielist{
		if item.ID == param["id"]{
			movielist = append(movielist[:i], movielist[i+1:]...)
			break
		}

	}
	json.NewEncoder(wr).Encode(movielist)
}

func getMovieById(wr http.ResponseWriter, rd *http.Request){
	wr.Header().Set("Content-Type", "application/json")
	param := mux.Vars(rd)
	for _,item := range movielist{
		if item.ID == param["id"]{
			json.NewEncoder(wr).Encode(item)
			return
		}
	}
}

func createNewMovie(wr http.ResponseWriter, rd *http.Request){
	wr.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(rd.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movielist = append(movielist, movie)
	json.NewEncoder(wr).Encode(movie)
}


func updateMovieName(wr http.ResponseWriter, rd *http.Request){
	wr.Header().Set("Content-Type", "application/json")
	param := mux.Vars(rd)


	for i, item := range movielist{
		if item.ID == param["id"]{
			movielist= append(movielist[:i], movielist[i+1:]...)
			var movie Movie
			_=json.NewDecoder(rd.Body).Decode(&movie)
			movie.ID=param["id"]
			movielist= append(movielist, movie)
			json.NewEncoder(wr).Encode(movie)
			return

		}
	}


}

func main(){
	route:=mux.NewRouter()

	movielist = append(movielist, Movie{ID: "1", Isbn: "573839", Moviename: "Movie1", Director: &Director{Fname:"james", Lname:"cameron"}})
	movielist = append(movielist, Movie{ID: "2", Isbn: "45829", Moviename: "Movie2", Director: &Director{Fname:"john", Lname:"paul"}})
	movielist = append(movielist, Movie{ID: "3", Isbn: "0193749", Moviename: "Movie3", Director: &Director{Fname:"Peaky", Lname:"Blinders"}})


	route.HandleFunc("/movies",getAllMovies).Methods("GET")
	route.HandleFunc("/movies/{id}",getMovieById).Methods("GET")
	route.HandleFunc("/movies",createNewMovie).Methods("POST")
	route.HandleFunc("/movies/{id}",updateMovieName).Methods("PUT")
	route.HandleFunc("/movies/{id}",deleteMovieName).Methods("DELETE")

	fmt.Printf("starting the server at port : 8000\n")
	log.Fatal(http.ListenAndServe(":8000",route))

}