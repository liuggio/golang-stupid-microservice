package main

import (
	"log"
	"net/http"
)

func main() {
	listenAndServe(helloHandler)
}

func listenAndServe(handlerFunc http.HandlerFunc) {
	log.Println("Starting 80")
	http.HandleFunc("/", logger(handlerFunc))
	http.ListenAndServe(":80", nil)
}
