package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)
	// $ curl -i -X POST http://localhost:8080/search  -d 'max=100&1=1&1=2&x=true&1=3&email=abc@gmailco.ma&zip=3'
	// HTTP/1.1 400 Bad Request
	// Content-Type: text/plain; charset=utf-8
	// X-Content-Type-Options: nosniff
	// Date: Sat, 03 Jun 2023 11:54:39 GMT
	// Content-Length: 35

	// 3 is not a valid zip portal format
}

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"1"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
		Email      string   `binding:"email"`
		Zip        string   `binding:"usZip"`
		CreditCard string   `binding:"creditCard"`
	}
	data.MaxResults = 10 // set default
	if err := UnpackBinding(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}
