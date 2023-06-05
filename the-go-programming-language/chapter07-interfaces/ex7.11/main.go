// Add additional handlers so that clients can create, read, uppdate, and delete database
// entries. For example, a request of the form /update?item=socks&price=6 will
// update the price of an item in the inventory and report an error if the item does not
// exist or if the price is invalid. (Warning: thich change introduces concurrent variable updates.)
package main

import (
	"fmt"
	"log"
	"net/http"
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
	for item, price := range db.db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
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

func main() {
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
