// Write a program that prints the SHA256 hash of its standard input by
// default but supports a command-line flag to print the SHA384 or
// SHA512 hash instead.
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var (
	sha = flag.Int("sha", 256, "Algorithm of hash")
)

func main() {
	flag.Parse()
	var function func([]byte) []byte
	switch *sha {
	case 384:
		function = func(b []byte) []byte {
			h := sha512.Sum384(b)
			return h[:]
		}
	case 512:
		function = func(b []byte) []byte {
			h := sha512.Sum512(b)
			return h[:]
		}
	default:
		function = func(b []byte) []byte {
			h := sha256.Sum256(b)
			return h[:]
		}
	}
	stdin := bufio.NewScanner(os.Stdin)
  var b []byte
  for stdin.Scan() {
    b = stdin.Bytes()
    break
  }
	fmt.Printf("%x\n", function(b))
}
