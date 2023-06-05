// Extend the jpeg program so that it converts any supported input format to
// to any output format, use image.Decode to detect the input format and
// a flag to select the output format
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
  outFmt := flag.String("output", "", "output format")
  flag.Parse()
  if *outFmt == "" {
    log.Fatalf("Usage: -output <output format>")
  }
	if err := convert(os.Stdin, os.Stdout, *outFmt); err != nil {
		fmt.Fprintf(os.Stderr, "image: %v\n", err)
		os.Exit(1)
	}
}

func convert(in io.Reader, out io.Writer, outFmt string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "input format =", kind)
	switch outFmt {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, nil)
	}
  return fmt.Errorf("not supported output %v", outFmt)
}
