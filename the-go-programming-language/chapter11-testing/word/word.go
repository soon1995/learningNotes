// Package word provides utilities for word games.
package word

import "unicode"

// gopl.io/ch11/word1
// IsPalindrome reports whether s reads the same forward and backward.
// func IsPalindrome(s string) bool {
// 	for i := range s {
// 		if s[i] != s[len(s)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// gopl.io/ch11/word2
// IsPalindrome reports whether s reads the same forward and backward.
// Letter case is ignored, as are non-letters.
func IsPalindrome(s string) bool {
  // pre-locate a sufficient space
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

func IsPalindromeOld2(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

func IsPalindromeOld1(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
