package zip

import (
	"archive/zip"
	"fmt"

	"example.go/ex10.2/reader"
)

func Reader(filename string) error {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, file := range r.File {
		fmt.Printf("%s\n", file.Name)
	}
	return nil
}

func init() {
	reader.Register("zip", Reader)
}
