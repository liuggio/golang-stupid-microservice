package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSingleHostShouldAccessOnlyWithLocalhost(t *testing.T) {
	res, req := makeRequestAndRecordResponse("GET", "http://localhost")
	middleware := single(fakeHandler, "localhost")

	middleware(res, req)
	assertResponse(t, res, 200, "yes")
}

func TestSingleHostShouldDenyAccess(t *testing.T) {
	res, req := makeRequestAndRecordResponse("GET", "http://youtube.com")
	middleware := single(fakeHandler, "localhost")

	middleware(res, req)
	assertResponse(t, res, 401, "Unauthorized")
}

func fakeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "yes")
}
