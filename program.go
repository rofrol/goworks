package main

import (
	"github.com/eknkc/amber"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	compiler := amber.New()
	compiler.ParseFile("tpl/index.amber")
	tpl, _ := compiler.Compile()
	tpl.Execute(w, map[string]interface{}{"date": time.Now().Format("Mon Jan 2 2006")})
}

func main() {
	http.HandleFunc("/home", http.HandlerFunc(hello))
	http.ListenAndServe(":4000", nil)
}
