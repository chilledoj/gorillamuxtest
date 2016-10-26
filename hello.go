package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func helloRoutes(s *mux.Router) *mux.Router {
	//r := mux.NewRouter().PathPrefix("/hello").Subrouter().StrictSlash(false)
	r := s.PathPrefix("/hello").Subrouter()
	r.HandleFunc("/", root)
	r.HandleFunc("/{name}", param)
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
