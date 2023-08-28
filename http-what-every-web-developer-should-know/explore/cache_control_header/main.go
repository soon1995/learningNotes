package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func main() {

	nowtime := time.Now()
	now := nowtime.UTC().Format(http.TimeFormat)
	http.HandleFunc("/cache", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=6")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Println("hello")
		fmt.Fprintln(w, "hello")
	})
	http.HandleFunc("/cache-with-last-modified", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Last-Modified", now)
		w.Header().Set("Cache-Control", "max-age=6")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		t, err := time.Parse(http.TimeFormat, r.Header.Get("If-Modified-Since"))
		if err == nil && nowtime.Before(t.Add(time.Second)) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
		fmt.Println("hello")
		fmt.Fprintln(w, "hello")
	})
	etag := uuid.New().String()
	http.HandleFunc("/cache-with-etag", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("ETag", etag)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if etag == r.Header.Get("If-None-Match") {
			w.WriteHeader(http.StatusNotModified)
			return
		}
		fmt.Println("hello")
		fmt.Fprintln(w, "hello")
	})
	http.ListenAndServe(":8080", nil)
}
