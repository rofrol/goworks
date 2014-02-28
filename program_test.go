package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	response := httptest.NewRecorder()

	request, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		log.Fatal(err)
	}

	hello(response, request)

	match := "Hello World"

	if got := response.Body.String(); got != match {
		t.Errorf("%s: %q = %v, want %v", "hello", response.Body, got, match)
	}
}
