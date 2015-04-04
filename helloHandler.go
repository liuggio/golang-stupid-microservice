package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Hello struct {
	Id   bson.ObjectId `bson:"_id"`
	Time time.Time     `bson:"time"`
	Name string        `bson:"string"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "" || r.URL.Path == "/" {
		http.Error(w, "Must Contain a Name", http.StatusBadRequest)
		return
	}

	session, _ := mgo.Dial(os.Getenv("DATABASE_PORT_27017_TCP_ADDR"))
	db := session.DB(os.Getenv("DB_NAME"))
	defer session.Close()

	// insert new record
	Hello := Hello{
		Id:   bson.NewObjectId(),
		Time: time.Now(),
		Name: r.URL.Path[1:],
	}
	db.C("Hellos").Insert(Hello)

	c, _ := db.C("Hellos").Count()
	fmt.Fprintf(w, "Hello %s,\nnumber of visits <b>%d</b>.\n", Hello.Name, c)
}
