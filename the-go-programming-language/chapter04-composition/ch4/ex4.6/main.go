// Write an in-place function that squashes each run of adjacent
// Unicode spaces (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a
// single ASCII space.
package main

import (
	"unicode"
	"unicode/utf8"
)

func Squashes(encoded []byte) []byte {
	isPrevSpace := false
	readPointer := 0
	writePointer := 0
	for readPointer < len(encoded) {
		readRune, sizeRead := utf8.DecodeRune(encoded[readPointer:])
		if unicode.IsSpace(readRune) {
			isPrevSpace = true
			readPointer += sizeRead
			continue
		} else {
			if isPrevSpace {
				sizeSpace := utf8.EncodeRune(encoded[writePointer:], int32(' '))
				isPrevSpace = false
				writePointer += sizeSpace
			}
			utf8.EncodeRune(encoded[writePointer:], readRune)
			writePointer += sizeRead
			readPointer += sizeRead
		}
	}
	if isPrevSpace {
		sizeSpace := utf8.EncodeRune(encoded[writePointer:], int32(' '))
		isPrevSpace = false
		writePointer += sizeSpace
	}

	return encoded[:writePointer]
}
