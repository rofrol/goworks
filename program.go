package main

import (
	"fmt"
	"github.com/eknkc/amber"
	"html"
	"net/http"
	"time"
)

func form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		compiler := amber.New()
		compiler.ParseFile("tpl/form.amber")
		tpl, err := compiler.Compile()
		if err != nil {
			serveError(w, err)
		}
		tpl.Execute(w, map[string]interface{}{})

	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			serveError(w, err)
		}
		str := html.EscapeString(r.Form.Get("str"))
		fmt.Fprint(w, Reverse(str))
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	compiler := amber.New()
	compiler.ParseFile("tpl/index.amber")
	tpl, err := compiler.Compile()
	if err != nil {
		fmt.Fprint(w, "Error: %v", err)
		return
	}
	tpl.Execute(w, map[string]interface{}{"date": time.Now().Format("Mon Jan 2 2006")})
}

func main() {
	http.HandleFunc("/home", http.HandlerFunc(hello))
	http.HandleFunc("/form", http.HandlerFunc(form))
	http.ListenAndServe(":4000", nil)
}
