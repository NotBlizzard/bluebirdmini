package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"strings"
	"time"
)

type User struct {
	Username      string
	Password_Hash string
}

type Post struct {
	Content  string
	Username string
	RandomID int
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
	return false
}

func Login(username string, password string) bool {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("bluebirdmini").C("users")
	user := User{}
	err = c.Find(bson.M{"username": strings.ToLower(username)}).One(&user)
	if err == nil {
		pass := bcrypt.CompareHashAndPassword([]byte(user.Password_Hash), []byte(password))
		if pass == nil {
			return true
		} else {
			return false
		}
	}
	return false
}

func MakePost(content string, username interface{}) int {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("bluebirdmini").C("posts")
	user := username.(string)
	r := rand.NewSource(time.Now().UnixNano())
	ra := rand.New(r)
	id := ra.Intn(1000000)
	fmt.Println(id)
	err = c.Insert(&Post{content, user, id})
	if err != nil {
		panic(err)
	}
	return id
}
