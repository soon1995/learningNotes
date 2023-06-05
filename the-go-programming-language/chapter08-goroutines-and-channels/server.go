package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"gopl.io/ch5/links"
	"gopl.io/ch8/thumbnail"
)

// gopl.io/ch8/spinner
func Spinner() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibboncci(%d) %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\"/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

// gopl.io/ch8/clock1
// Clock 1 is a TCP server that periodically writes the time
func Clock1() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err) // e.g. connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time
	}
}

// gopl.io/ch8/clock2
func Clock2() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err) // e.g. connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnecter
		}
		time.Sleep(1 * time.Second)
	}
}

// gopl.io/ch8/reverb1
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn2(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func Echo1() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn2(conn)
	}
}

// gopl.io/ch8/pipeline1
func Pipeline1() {
	naturals := make(chan int)
	squares := make(chan int)
	//Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
	}()
	//Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()
	// Printer (in main go routing)
	for {
		fmt.Println(<-squares)
	}
}

func Pipeline1_5() {
	naturals := make(chan int)
	squares := make(chan int)
	//Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	//Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()
	// Printer (in main go routing)
	for {
		x, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(x)
	}
}

// gopl.io/ch8/pipeline2
func Pipeline2() {
	naturals := make(chan int)
	squares := make(chan int)
	//Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	//Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	// Printer (in main go routing)
	for x := range squares {
		fmt.Println(x)
	}
}

// gopl.io/ch8/pipeline3
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func Pipeline3() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

// return the first
func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses
}

func request(hostname string) (response string) { /*...*/ return "aa" }

// gopl.io/ch8/thumbnail
// makeThumbnails makes thumbnails of the specified files.
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// TOO FAST! makeThumbnails returns before it has finished doing what it was supposed to do.
// Doesn't wait for them to finish...
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // NOTE: ignoring errors
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
}

// It returns an error if any step failed
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: incorrect: goroutine leak!
		}
	}
	return nil
}

// Solution to makeThumbnails4, using a sufficient capacity buffered channel
// It returns the generated file names in an arbitrary order.
// or an error if any step failed
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

// makeThumbnail6 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}
	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

// gopl.io/ch8/crawl1
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// It has two problems.
//  1. The program created so many network connections at once that it exceeded the per-process limit on
//     on the number or open files, causing operation such as DNS lookups and calls to net.Dial to start failing.
//  2. The program never terminates.
func Crawl() {
	worklist := make(chan []string)

	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

// gopl.io/ch8/crawl2
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent request.
var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

// n for break the program to terminate
func Crawl2() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist, that are yet to occured

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

// Alternative for Crawl2
// To save space, we have not addressed the problem of termination in this example.
func Crawl3() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	var n int // number of pending sends to worklist, that are yet to occured

	// Add command-line arguments to worklist.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}
	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				unseenLinks <- link
			}
		}
	}
}

// old
func CrawlWithoutConcurrent() {
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// gopl.io/ch8/countdown1
func Countdown1() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("launch!")
}

// gopl.io/ch8/countdown2
func Countdown2() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
	// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

// more subtle
func Countdown2_5() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}

// gopl.io/ch8/countdown3
// goroutine leak! ticker goroutine is still there trying in vain to send on
// a channel from which no goroutine is receiving
func Countdown3() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown. Press return to abord.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		//Do nothing
		case <-abort:
			fmt.Println("Launch aborted")
			return
		}
	}
	launch()
}

func Countdown3a() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown. Press return to abord.")
	ticker := time.NewTicker(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-ticker.C:
		//Do nothing
		case <-abort:
			ticker.Stop() // cause the ticker's goroutine to terminate
			fmt.Println("Launch aborted")
			return
		}
	}
	launch()
}

// gopl.io/ch8/du1
// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du1: %v\n", err)
				return
			}
			fileSizes <- info.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []fs.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func WalkDir() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Traverse the file tree.
	fileSize := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSize)
		}
		close(fileSize)
	}()
	// print the results
	var nfiles, nbytes int64
	for size := range fileSize {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %f MB\n", nfiles, float64(nbytes)/1000000)
}

// gopl.io/ch8/du2
// It still takes too long to finish. There's no reason why all the calls to walkDir
// can't be done concurrently.
var verbose = flag.Bool("v", false, "show verbose progress messages")

func WalkDir2() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Traverse the file tree.
	fileSize := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSize)
		}
		close(fileSize)
	}()
	// print the results
	var nfiles, nbytes int64
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
loop:
	for {
		select {
		case size, ok := <-fileSize:
			if !ok {
				break loop // fileSizes was close
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

// gopl.io/ch8/du3
// Since this program creates many thousands of goroutines at its peak,
// we have to change dirents to use a counting semophore to prevent it from
// opening too many files at once
func walkDir3(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	for _, entry := range dirents1(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir3(subdir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func WalkDir3() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Traverse the file tree.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir3(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	// print the results
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was close
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents1(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

// gopl.io/ch8/du4
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// cancel travesal when input is detected
func WalkDirDu4() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Traverse the file tree.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDirDu4(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// print the results
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			// for range fileSizes {
			// 	// Do nothing.
			// }
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was close
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
}

// walkDir pools the cancellation status when it begins
// It might be profitable to poll the cancellation status again within walkDir's loop,
// to avoid creating goroutines after cancellation event.
// Cancellation involves a trade-off; a quicker response often requires more intrusive changes
// to program logic.
// Ensuring that no expensive operations ever occur after the cancellation event may require
// updating many places in your code, but often most of the benefit can be obtained by checking for cancellation
// in a few important places.
func walkDirDu4(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range direntsDu4(dir) {
		// ...
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir3(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func direntsDu4(dir string) []fs.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du4: %v\n", err)
		return nil
	}
	return entries
}

// gopl.io/ch8/chat
func Chat() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnChat(conn)
	}
}

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client message
)

func broadcaster() {
	clients := make(map[client]bool) // all connected client
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConnChat(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch chan string) {
  for msg := range ch {
    fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
  }
}

func main() {
	// Spinner()
	// Clock1()
	// Echo1()
	// Pipeline1()
	// Pipeline1_5()
	// Pipeline2()
	// Pipeline3()
	// fmt.Println(mirroredQuery())
	// Crawl2()
	// Crawl3()
	// CrawlWithoutConcurrent()
	// Countdown1()
	// Countdown2()
	// Countdown3a()
	// WalkDir()
	// t1 := time.Now()
	// WalkDir2()
	// WalkDir3()
	// fmt.Println("diff", time.Now().Sub(t1))
	// WalkDirDu4()
	// panic("asd")
  Chat()

}
