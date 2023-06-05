// Change the chat server's network protocol so that each client provides its
// name on entering. Use that name instead of the network address when prefixing
// each message with its sender's identify.
package main

import (
	"log"
	"net/http"
)

func main() {
	go broadcaster()

	http.HandleFunc("/", SignInHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = r.RemoteAddr
	}
	hi, ok := w.(http.Hijacker)
	if !ok {
		log.Print("cannot Hijack.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	conn, _, err := hi.Hijack()
	if err != nil {
		log.Print("Hijack error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	go handleConnChat(conn, name)
}

// Tips: use 
// nc localhost 8000
// GET /?name=yourname HTTP/1.0
// <Enter>
// Start using this server
