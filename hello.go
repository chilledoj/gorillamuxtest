package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func helloRoutes() *mux.Router {
	r := mux.NewRouter() //.PathPrefix("/api").Subrouter().StrictSlash(true)
	//r := s.PathPrefix("/hello").Subrouter()
	r.HandleFunc("/hello", root)
	r.HandleFunc("/hello/{name}", param)
	return r
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func param(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.Write([]byte("Hello " + name))
}
