package main

import (
	"log"
	"net/http"
	"time"
)

func logger(handlerFunc http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("Access [%s] %s", r.Method, r.URL.Path)

		handler := http.HandlerFunc(handlerFunc)
		handler.ServeHTTP(w, r)

		log.Printf("Completed %q\n", time.Since(start))

		return
	}

}
