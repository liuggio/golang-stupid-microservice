package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func makeRequestAndRecordResponse(method string, path string) (*httptest.ResponseRecorder, *http.Request) {

	req, _ := http.NewRequest(method, path, nil)

	return httptest.NewRecorder(), req
}

func assertResponse(t *testing.T, w *httptest.ResponseRecorder, assertStatus int, assertBody string) {

	bodyBytes, _ := ioutil.ReadAll(w.Body)

	if assertStatus > 0 {
		if w.Code != assertStatus {
			t.Errorf("Status code should be %v, was %d with [%s]", assertStatus, w.Code, bodyBytes)
		}
	}

	if assertBody != "" {
		assertion := []byte(assertBody)
		if !bytes.Contains(bodyBytes, assertion) {
			t.Errorf("body content assertion was [%s] found [%s]", assertBody, bodyBytes)
		}
	}
}
