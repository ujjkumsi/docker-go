package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	eh "github.com/ujjkumsi/docker-go/best-practices"
	"github.com/ujjkumsi/docker-go/dao"
	"github.com/ujjkumsi/docker-go/dialogflow"
	"github.com/ujjkumsi/docker-go/models"
)

var moviesDao = dao.MoviesDAO{}

func allMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func findMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func createMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = gocql.TimeUUID()

	if err := moviesDao.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, movie)
}

func updateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func deleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the RESTful api")
	errorHandler()
}

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}

func init() {
	moviesDao.Database = "movieapi"
	moviesDao.Server = "cassandra"
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", heartbeat).Methods("GET")
	r.HandleFunc("/movies", allMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", createMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", updateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", deleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", findMovieEndpoint).Methods("GET")
	r.HandleFunc("/action", dialogflow.DialogflowHandler).Methods("POST")

	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func errorHandler() {
	result, err := eh.Divide(1.0, 0.0)
	if err != nil {
		switch err.(type) {
		case *eh.ErrZeroDivision:
			fmt.Println(err.Error())
		default:
			fmt.Println("What the h* just happened?")
		}
	}
	fmt.Println(result)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
