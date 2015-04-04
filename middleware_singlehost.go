package main

import (
	"net/http"
)

func single(handlerFunc http.HandlerFunc, allowedHost string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		if host == allowedHost {

			handler := http.HandlerFunc(handlerFunc)
			handler.ServeHTTP(w, r)
			return
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
