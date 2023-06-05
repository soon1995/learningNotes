// The JSON-based web service of the Open Movie Database lets you search
// https:// omdbapi.com/ for a movie by name and download its poster image.
// Write a tool poster that downloads the poster image for the movie named on the command line.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/metal3d/go-slugify"
)

const (
	APIURL = "https://omdbapi.com/?"
)

var (
	key = "" // 391c33d4
)

type Movie struct {
	Title  string
	Year   string
	Poster string
}

func (m Movie) posterFilename() string {
	ext := filepath.Ext(m.Poster)
	title := slugify.Marshal(m.Title)
	return fmt.Sprintf("%s_(%s)%s", title, m.Year, ext)
}

func getMovie(title string) (movie Movie, err error) {
	url_ := fmt.Sprintf("%st=%s&apikey=%s", APIURL, url.QueryEscape(title), url.QueryEscape(key))
	resp, err := http.Get(url_)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%d response from %s", resp.StatusCode, url_)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		return
	}
	return
}

func (m Movie) writePoster() error {
	url_ := m.Poster
  if url_ == "N/A" {
    return fmt.Errorf("No poster for %s", m.Title)
  }
	resp, err := http.Get(url_)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d response from %s", resp.StatusCode, url_)
	}
	file, err := os.Create(m.posterFilename())
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: poster MOVIE_TITLE")
		os.Exit(1)
	}
	key = os.Getenv("OPEN_MOVIE_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "api key is required")
		os.Exit(1)
	}
	title := os.Args[1]
	movie, err := getMovie(title)
	if err != nil {
		log.Fatal(err)
	}
	if zero := new(Movie); movie == *zero {
		fmt.Fprintf(os.Stderr, "No results for %q\n", title)
		os.Exit(2)
	}

	err = movie.writePoster()
	if err != nil {
		log.Fatal(err)
	}
}

// ./ex4.13 spider
