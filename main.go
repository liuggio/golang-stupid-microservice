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

	log.Println("DBName" + os.Getenv("DB_NAME"))
	log.Println("Starting at:" + port)
	http.HandleFunc("/", logger(handlerFunc))
	log.Fatal(http.ListenAndServe(":"+port, nil))

	log.Println("exiting")
}
