// Define a generic archive file-reading function capable of reading ZIP files (archive/zip)
// and POSIX tar files (archive/tar). Use a registration mechanism similar to the one described above
// so that support for each file format can be plugged in using blank imports.
package main

import (
	"example.go/ex10.2/reader"
	_ "example.go/ex10.2/reader/plugin/tar"
	_ "example.go/ex10.2/reader/plugin/zip"
)

func main() {
	reader.ArchiveReader("arc.zip")
}
