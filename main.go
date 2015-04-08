package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	listenAndServe(helloHandler)
}

func listenAndServe(handlerFunc http.HandlerFunc) {

	port := os.Getenv("PORT")
	if "" == port {
		port = "80"
	}
	if "" == os.Getenv("DB_NAME") {
		os.Setenv("DB_NAME", "DB")
	}
	if "" == os.Getenv("mongo_PORT") {
		os.Setenv("mongo_PORT", "mongodb://mongo:27017")
	}

	log.Println("DBName:     \t" + os.Getenv("DB_NAME"))
	log.Println("Mongo at:   \t" + os.Getenv("mongo_PORT"))
	log.Println("Starting at:\t" + port)
	http.HandleFunc("/", logger(handlerFunc))
	log.Fatal(http.ListenAndServe(":"+port, nil))

	log.Println("exiting")
}
