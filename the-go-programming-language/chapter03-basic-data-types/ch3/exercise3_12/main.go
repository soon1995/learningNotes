// Write a function that reports whether two strings are anagrams of each
// other, that is, they contain the same letters in a different order
package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("fried", "fired"))
	fmt.Println(isAnagram("dreads", "sadder"))
	fmt.Println(isAnagram("abc", "der"))
	fmt.Println(isAnagram("我很好", "你好吗"))
	fmt.Println(isAnagram("我很好", "很好我"))
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

  freq := make(map[rune]int)
  for _, c := range a {
    freq[c]++
  }
  for _, c := range b {
    freq[c]--
  }
  for _,v := range freq {
    if v != 0 {
      return false
    }
  }
  return true
}

func isAnagramExample(a, b string) bool {
	aFreq := make(map[rune]int)
	for _, c := range a {
		aFreq[c]++
	}
	bFreq := make(map[rune]int)
	for _, c := range b {
		bFreq[c]++
	}
	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}
	for k, v := range bFreq {
		if aFreq[k] != v {
			return false
		}
	}
	return true
}
