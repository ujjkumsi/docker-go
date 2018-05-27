package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	eh "github.com/ujjkumsi/docker-go/best-practices"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is the RESTful api")
	errorHandler()
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)

	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	http.ListenAndServe(":8080", router)
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
