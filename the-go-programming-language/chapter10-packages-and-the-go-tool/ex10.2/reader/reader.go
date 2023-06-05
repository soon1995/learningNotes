package reader

import (
	"fmt"
	"strings"
)

var registry = make(map[string]func(string) error)

func Register(name string, f func(string) error) {
	registry[name] = f
}

func ArchiveReader(s string) error {
	dot := strings.LastIndex(s, ".")
	if dot == -1 || dot == len(s)-1 {
		return fmt.Errorf("unknown archive extension: %s", s)
	}
	ext := s[dot+1:]
	f := registry[ext]
	if f == nil {
		return fmt.Errorf("reader: %v not registered.", ext)
	}
	return f(s)
}
