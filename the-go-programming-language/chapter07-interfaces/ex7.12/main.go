// Change the handler for /list to print its output as an HTML table,
// not text.
// You may find html/template package useful
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	db map[string]dollars
	sync.RWMutex
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	db.RLock()
	if err := templ.Execute(w, db.db); err != nil {
		db.RUnlock()
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "cannot execute template")
		return
	}
	db.RUnlock()
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	db.RLock()
	price, ok := db.db[item]
	db.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)

}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	p := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(p, 32)
	if err != nil || name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad request")
		return
	}
	db.Lock()
	db.db[name] = dollars(price)
	db.Unlock()
	w.WriteHeader(http.StatusOK)
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	p := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(p, 32)
	if err != nil || name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad request")
		return
	}
	db.RLock()
	if _, ok := db.db[name]; !ok {
		db.RUnlock()
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item %s not exist", name)
		return
	}
	db.RUnlock()

	db.Lock()
	db.db[name] = dollars(price)
	db.Unlock()
	w.WriteHeader(http.StatusOK)
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	db.Lock()
	delete(db.db, name)
	db.Unlock()
}

var templ *template.Template

func main() {
	file, err := os.Open("template.html")
	if err != nil {
		log.Fatal(err)
	}
	b := &bytes.Buffer{}
	io.Copy(b, file)
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
	templ, err = template.New("list").Parse(b.String())
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]dollars)
	m["shoes"] = 50
	m["socks"] = 5
	db := &database{db: m}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
