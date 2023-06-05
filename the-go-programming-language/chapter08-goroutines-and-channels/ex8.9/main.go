// Write a version of du that computes and periodically displays seperate
// totals for each root directories
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(root int, dir string, wg *sync.WaitGroup, files chan<- *File) {
	defer wg.Done()
	for _, entry := range dirents1(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(root, subdir, wg, files)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du1: %v\n", err)
				return
			}
			files <- &File{root, info.Size()}
		}
	}
}

type File struct {
	root  int
	bytes int64
}

func WalkDir() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Traverse the file tree.
	files := make(chan *File)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(i, root, &n, files)
	}
	go func() {
		n.Wait()
		close(files)
	}()
	// print the results
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
loop:
	for {
		select {
		case file, ok := <-files:
			if !ok {
				break loop // fileSizes was close
			}
			nfiles[file.root]++
			nbytes[file.root] += file.bytes
		case <-tick:
			printDiskUsageWithRoot(roots, nfiles, nbytes)
		}
	}
	printDiskUsageWithRoot(roots, nfiles, nbytes)
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func printDiskUsageWithRoot(roots []string, nfiles, nbytes []int64) {
	for i, v := range roots {
		fmt.Printf("%s - %d files %f MB\n", v, nfiles[i], float64(nbytes[i])/1000000)
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents1(dir string) []fs.DirEntry {
	sema <- struct{}{}
	entries, err := os.ReadDir(dir)
	<-sema
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	WalkDir()
}
