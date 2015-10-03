package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type User struct {
	Username      string
	Password_Hash string
}

func Register(username string, password string) bool {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("bluebirdmini").C("users")
	user := User{}
	err = c.Find(bson.M{"username": strings.ToLower(username)}).One(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println("it is '" + user.Username + "'")
	if user.Username != "" {
		return false
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	err = c.Insert(&User{strings.ToLower(username), string(pass)})

	if err != nil {
		panic(err)
	}
	return true
}
