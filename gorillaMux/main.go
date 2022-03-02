package main

import (
	"net/http"

	"github.com/gorilla/mux"
	a "github.com/shivanisharma200/learningGo/gorillaMux/restAPIs"
)

func main() {
	// creating new request router
	r := mux.NewRouter()

	// REST-API
	// creating edpoints
	r.HandleFunc("/posts", a.GetPosts).Methods("GET")
	r.HandleFunc("/posts", a.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", a.GetPost).Methods("GET")
	r.HandleFunc("/posts/{id}", a.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", a.DeletePost).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
