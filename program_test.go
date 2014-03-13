package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	w := httptest.NewRecorder()

	resource := "/hello"
	r, err := http.NewRequest("GET", resource, nil)
	if err != nil {
		log.Fatal(err)
	}

	hello(w, r)

	match := fmt.Sprintf("<h1>Hello World</h1>\n<p>Today is %s.</p>\n", time.Now().Format("Mon Jan 2 2006"))

	if got := w.Body.String(); got != match {
		t.Errorf("%s: %q = %v, want %v", resource, w.Body, got, match)
	}
}

func TestForm(t *testing.T) {
	w := httptest.NewRecorder()

	data := url.Values{}
	str := "The quick brown 狐 jumped over the lazy 犬"
	match := "犬 yzal eht revo depmuj 狐 nworb kciuq ehT"
	data.Set("str", str)

	resource := "/form"
	r, err := http.NewRequest("POST", resource, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
	if err != nil {
		serveError(w, err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	form(w, r)

	if got := w.Body.String(); got != match {
		t.Errorf("%s: %q = %v, want %v", resource, w.Body, got, match)
	}
}

func TestStatic(t *testing.T) {
	ts := httptest.NewServer(staticHandler())

	client := &http.Client{}
	resource := "/public/"
	r, err := http.NewRequest("GET", ts.URL + resource, nil)
	if err != nil {
		log.Fatal(err)
	}

	w, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	// TestServeFile http://golang.org/src/pkg/net/http/fs_test.go
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	match :=
		`<html>
  <head>
    <link rel="stylesheet" type="text/css" href="/public/main.css"/>
  </head>
  <body>
    <p>I am red!</p>
  </body>
</html>
`

	if got := string(b); got != match {
		t.Errorf(">>>> %s: \n\n%q = \n\n%v, want \n\n%v", resource, w.Body, got, match)
	}
}
