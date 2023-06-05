package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {

	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Color           bool
		Oscars          []string
		Sequel          *string
	}
	s := `((Title "Dr. Strangelove") (Subtitle "How I Learned to Stop Worrying and Love the Bomb") (Year 1964) (Color false) (Actor (("Gen. Buck Turgidson" "George C. Scott") ("Brig. Gen. Jack D. Ripper" "Sterling Hayden") ("Maj. T.J. \"King\" Kong" "Slim Pickens") ("Dr. Strangelove" "Peter Sellers") ("Grp. Capt. Lionel Mandrake" "Peter Sellers") ("Pres. Merkin Muffley" "Peter Sellers"))) (Oscars ("Best Actor (Nomin.)" "Best Adapted Screenplay (N omin.)" "Best Director (Nomin.)" "Best Picture (Nomin.)")) (Sequel nil))`
	r := bytes.NewReader([]byte(s))
	dec := NewDecoder(r)

	movie := &Movie{}
	err := dec.Decode(movie)
	if err != nil {
		t.Error(err)
	}
  fmt.Println(movie)
}
