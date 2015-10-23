package main

import (
	"fmt"
	"github.com/notblizzard/bluebirdmini/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/notblizzard/bluebirdmini/Godeps/_workspace/src/github.com/gorilla/sessions"
	"html/template"
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
			session, _ := store.Get(r, "user")

			session.Values["logged_in"] = true
			session.Save(r, w)
			http.Redirect(w, r, "/home", 302)
			return
		}
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/login.html")
		t.Execute(w, nil)
		return
	} else {
		r.ParseForm()
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		login := Login(username, password)
		if login == true {
			session, _ := store.Get(r, "user")
			session.Values["logged_in"] = true
			session.Values["username"] = username
			session.Save(r, w)
			http.Redirect(w, r, "/home", 302)
			return
		} else {
			http.Redirect(w, r, "/login", 302)
			return
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	session.Values["logged_in"] = nil
	session.Values["username"] = nil
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
	return
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if session.Values["logged_in"] == nil {
		http.Redirect(w, r, "/login", 302)
		return
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/new.html")
		t.Execute(w, nil)
		return
	} else if r.Method == "POST" {
		r.ParseForm()
		content := r.PostFormValue("content")
		username := session.Values["username"]
		fmt.Println(username)

		st := MakePost(content, username)
		user := username.(string)
		fmt.Println(st)
		http.Redirect(w, r, "/"+user+"/"+st, 302)
		return
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/post.html")
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	post := GetPost(id)
	t.Execute(w, post)
	return

}
