package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func init() {
	log.Println("Initialising")
}

func main() {

	router := buildRoutes()

	n := negroni.Classic() // Includes the default middlewares
	n.UseHandler(router)
	// Taken from Gorilla/mux example
	srv := &http.Server{
		Handler: n,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func buildRoutes() *mux.Router {
	r := mux.NewRouter()

	home := homeHandlers{}
	r.NotFoundHandler = http.HandlerFunc(home.errPage404)
	r.HandleFunc("/", home.homePage)
	r.HandleFunc("/about", home.aboutPage)

	r.PathPrefix("/hello").Handler(helloRoutes())
	return r
}
