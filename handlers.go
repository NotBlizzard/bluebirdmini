package main

import (
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("hello"))

func RootHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if session.Values["logged_in"] != nil {
		http.Redirect(w, r, "/home", 302)
	}

	t, _ := template.ParseFiles("views/root.html")
	t.Execute(w, nil)
	return

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if session.Values["logged_in"] == nil {
		http.Redirect(w, r, "/", 302)
	}
	t, _ := template.ParseFiles("views/home.html")
	t.Execute(w, nil)
	return
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/register.html")
		t.Execute(w, nil)
		return
	} else if r.Method == "POST" {
		r.ParseForm()
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		reg := Register(username, password)
		if reg == false {
			http.Redirect(w, r, "/", 302)
			return
		} else {
			session, err := store.Get(r, "user")
			if err != nil {
				log.Fatal(err)
			}
			session.Values["logged_in"] = true
			session.Save(r, w)
			http.Redirect(w, r, "/home", 302)
			return
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "user")
	if err != nil {
		log.Fatal(err)
	}

	if session.Values["logged_in"] != nil {
		session.Values["logged_in"] = nil
	}
	http.Redirect(w, r, "/", 302)
	return

}
