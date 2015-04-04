package main

import (
	"testing"
)

func TestShouldResponse(t *testing.T) {

	res, req := makeRequestAndRecordResponse("GET", "/liuggio")

	helloHandler(res, req)
	assertResponse(t, res, 200, "liuggio")
}

func TestResponseShouldNotBeEmpty(t *testing.T) {

	res, req := makeRequestAndRecordResponse("GET", "/")

	helloHandler(res, req)
	assertResponse(t, res, 400, "Must Contain a Name")
}
