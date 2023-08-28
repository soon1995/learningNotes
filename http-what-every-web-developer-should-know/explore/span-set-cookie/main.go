package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/span", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Set-Cookie", "k=v; domain=.google.com;")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Println("hello")
		fmt.Fprintln(w, "hello")
	})
	http.HandleFunc("/not-span", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Set-Cookie", "k=v;")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Println("hello")
		fmt.Fprintln(w, "hello")
	})
	http.ListenAndServe(":8080", nil)
}
