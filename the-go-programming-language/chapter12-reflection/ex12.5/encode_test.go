package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func TestEncode(t *testing.T) {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	got, err := Marshal(strangelove)
	if err != nil {
		t.Errorf("got got err: %s", err)
	}
	movie := Movie{}
	err = json.Unmarshal(got, &movie)
	if err != nil {
		t.Errorf("unmarshal json err: %s", err)
	}
	if !reflect.DeepEqual(movie, strangelove) {
		t.Errorf("not equal")
	}
}
