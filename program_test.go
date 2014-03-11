package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
	"strconv"
)

func TestHello(t *testing.T) {
	w := httptest.NewRecorder()

	resource := "/hello"
	request, err := http.NewRequest("GET", resource, nil)
	if err != nil {
		log.Fatal(err)
	}

	hello(w, request)

	match := fmt.Sprintf("<h1>Hello World</h1>\n<p>Today is %s.</p>\n", time.Now().Format("Mon Jan 2 2006"))

	if got := w.Body.String(); got != match {
		t.Errorf("%s: %q = %v, want %v", resource, w.Body, got, match)
	}
}

func TestForm(t *testing.T) {
	w := httptest.NewRecorder()

	data := url.Values{}
	str := "The quick brown 狐 jumped over the lazy 犬"
	data.Set("str", str)

	resource := "/form"
	r, err := http.NewRequest("POST", resource, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
	if err != nil {
		serveError(w, err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))


	form(w, r)

	match := "犬 yzal eht revo depmuj 狐 nworb kciuq ehT"

	if got := w.Body.String(); got != match {
		t.Errorf("%s: %q = %v, want %v", resource, w.Body, got, match)
	}
}
