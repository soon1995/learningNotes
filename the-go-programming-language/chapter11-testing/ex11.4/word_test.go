// Modify randomPalindrome to exercise IsPalindrome's handling of punctuation and spaces
package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

// randomPalindrome returns a parindrome whose length and contents
// are derived from the pseudo-random number generator rng/
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	i, j := 0, n-1
	for i < j {
		prob := rng.Float64()
		if prob < 0.1 {
			runes[i] = Punct[int(rng.Float64()*float64(len(Punct)-1))]
			i++
		} else if prob < 0.3 {
			runes[i] = Space[int(rng.Float64()*float64(len(Space)-1))]
			i++
		} else {
			r := rune(rng.Intn(0x1000)) // random rne up to '\u0999'
			runes[i] = r
			runes[j] = r
			i++
			j--
		}
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrme(%q) = false", p)
		}
	}
}

var Punct []rune
var Space = [...]rune{'\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0}

func init() {
	for i := 0x21; i <= 0x7E; i++ {
		r := rune(i)
		if unicode.IsPunct(r) {
			Punct = append(Punct, r)
		}
	}
}
