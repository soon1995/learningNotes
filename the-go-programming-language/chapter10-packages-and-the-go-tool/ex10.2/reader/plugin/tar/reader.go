package tar

import (
	"archive/tar"
	"fmt"
	"io"
	"os"

	"example.go/ex10.2/reader"
)

func Reader(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	r := tar.NewReader(file)
	for {
		h, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", h.Name)
	}
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}

func init() {
	reader.Register("tar", Reader)
}
