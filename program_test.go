package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"fmt"
)

func TestHello(t *testing.T) {
	response := httptest.NewRecorder()

	request, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		log.Fatal(err)
	}

	hello(response, request)

	match := fmt.Sprintf("<h1>Hello World</h1>\n<p>Today is %s.</p>\n", time.Now().Format("Mon Jan 2 2006"))

	if got := response.Body.String(); got != match {
		t.Errorf("%s: %q = %v, want %v", "hello", response.Body, got, match)
	}
}
