package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/search", search)
	http.HandleFunc("/packapi", packapi)
	http.ListenAndServe(":8080", nil)
	// $ curl -i -X POST http://localhost:8080/search  -d 'max=100&1=1&1=2&x=true&1=3'
	// HTTP/1.1 200 OK
	// Date: Sat, 03 Jun 2023 05:46:17 GMT
	// Content-Length: 51
	// Content-Type: text/plain; charset=utf-8

	// Search: {Labels:[1 2 3] MaxResults:100 Exact:true}
}

func packapi(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"1"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.Labels = append(data.Labels, "test1")
	data.Labels = append(data.Labels, "testw")
	data.MaxResults = 10 // set default
	data.Exact = true    // set default
	res, err := Pack(&data)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	fmt.Printf("%#v", res)
	// ...rest of handler...
}
