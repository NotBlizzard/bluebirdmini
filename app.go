package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.ListenAndServe(":8000", nil)
}
