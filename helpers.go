package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math/rand"
	"strconv"
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
	RandomID string
}

func Register(username string, password string) bool {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	c := session.DB("bluebirdmini").C("users")
	user := User{}
	err = c.Find(bson.M{"username": strings.ToLower(username)}).One(&user)
	if err != nil {

		pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		err = c.Insert(&User{strings.ToLower(username), string(pass)})

		if err != nil {
			log.Fatal(err)
		}
		return true
	}
	return false
}

func Login(username string, password string) bool {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
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

func MakePost(content string, username interface{}) string {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	c := session.DB("bluebirdmini").C("posts")
	user := username.(string)
	r := rand.NewSource(time.Now().UnixNano())
	ra := rand.New(r)
	i := ra.Intn(1000000)
	id := strconv.Itoa(i)
	fmt.Println(id)
	err = c.Insert(&Post{content, user, id})
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func GetPost(id string) Post {
	session, _ := mgo.Dial("localhost")
	defer session.Close()
	c := session.DB("bluebirdmini").C("posts")
	post := Post{}
	err := c.Find(bson.M{"randomid": id}).One(&post)
	if err != nil {
		log.Fatal(err)
	}
	return post
}
