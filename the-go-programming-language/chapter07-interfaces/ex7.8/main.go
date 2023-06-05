// Many GUIs provide a table widget with a stateful multi-tier sort:
// the primary sort key is the most recently clicked column head, the secondary sort key
// is the second-most recently clicked column head, and so on. Define an implementation
// of sort.Interface for use by such a table. Compare that approach with repeated sorting using
// sort.Stable
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

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

func main() {
	sortBy := os.Args[1:]
	my := &MySort{tracks, sortBy}
	sort.Sort(my)
	printTracks(my.tracks)
}
