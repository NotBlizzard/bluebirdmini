package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	r.HandleFunc("/new", NewHandler)

	r.HandleFunc("/home", HomeHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/logout", LogoutHandler)
	r.HandleFunc("/register", RegisterHandler)
	r.HandleFunc("/{user}/{id}", PostHandler)
	http.Handle("/", r)

	http.ListenAndServe(":8000", nil)
}
