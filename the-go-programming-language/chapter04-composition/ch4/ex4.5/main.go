// Write an in-place function to eliminate adjacent duplicates
// in a []string slice
package main

// func deleteDuplicate(s []string) []string {
// 	i := len(s) - 1
// 	finalLen := len(s)
// 	for i > 0 {
// 		if s[i] == s[i-1] {
// 			copy(s[i-1:], s[i:])
// 			finalLen--
// 		}
// 		i--
// 	}
// 	return s[:finalLen]
// }

func deleteDuplicate(s []string) []string {
  if len(s) == 0 {
    return s
  }
	w := 0
	for _, v := range s {
		if s[w] == v {
			continue
		}
		w++
		s[w] = v
	}
	return s[:w+1]
}
