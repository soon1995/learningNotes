// The sort.Interface type can be adapted to other uses. Write a function
// IsPalindrome(s sort.Interface) bool that reports whether the sequence s is a
// palindrome, in other words, reversing the sequence would not change it. Assume
// that the elements at indices i and j are equal if !s.Less(i, j) && !s.Less(j, i).
package main

import "sort"

func IsPalindrome(s sort.Interface) bool {
	n := s.Len() - 1
	if n <= 0 {
		return true
	}
	for i := 0; i < n>>1; i++ {
		if s.Less(i, n-i) || s.Less(n-i, i) {
			return false
		}
	}
	return true
}
