// Use the html/template package to replace printTracks with a function
// that displays the tracks as an HTML table. Use the solution to the previous
// exercise to arrange that each click on a column head makes an HTTP request
// to sort the table
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Moby", "Boby", 1992, length("3m37s")},
	{"Go", "Moby", "Aoby", 1992, length("3m37s")},
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type MySort struct {
	tracks []*Track
	sortBy []string
}

func (m *MySort) Len() int      { return len(m.tracks) }
func (m *MySort) Swap(i, j int) { m.tracks[i], m.tracks[j] = m.tracks[j], m.tracks[i] }
func (m *MySort) Less(i, j int) bool {
	if len(m.sortBy) == 0 {
		return false
	}
	s := 0
	for k := 0; k < 3 && s < len(m.sortBy); k++ {
		switch m.sortBy[s] {
		case "title":
			if m.tracks[i].Title != m.tracks[j].Title {
				return m.tracks[i].Title < m.tracks[j].Title
			}
		case "artist":
			if m.tracks[i].Artist != m.tracks[j].Artist {
				return m.tracks[i].Artist < m.tracks[j].Artist
			}
		case "album":
			if m.tracks[i].Album != m.tracks[j].Album {
				return m.tracks[i].Album < m.tracks[j].Album
			}
		}
		s++
	}
	return false
}

var templ string

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
	templ = b.String()

	http.HandleFunc("/sort", SortHandler)

	http.ListenAndServe(":8080", nil)
}

func SortHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	p := r.URL.Query().Get("sort")
	temp, err := template.New("sort").Parse(templ)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	my := &MySort{
		tracks: tracks,
		sortBy: []string{p},
	}
	sort.Sort(my)
	resp := &bytes.Buffer{}

	if err := temp.Execute(resp, my.tracks); err != nil {
		log.Printf("cannot execute template: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp.Bytes())
	return

}
