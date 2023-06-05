package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	htmltemplate "html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"
	"unicode"
	"unicode/utf8"
)

func main() {
	// dedup()
	// charcount()
	// AddEdge("a", "b")
	// fmt.Println(hasEdge("a", "b"))
	// arr := []int{3, 2, 1, 4, 5}
	// Sort(arr)
	// fmt.Println(arr)
	// compare()
	// embed()
	// movie()
	// SearchIssues() // go run main.go repo:golang/go is:open json decoder
	// issuesReport()
	// issuesReportHTML()
  autoescape()
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// append
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		// There is insufficient space, allocate a new arry.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

// append2
func appendInt2(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		// There is insufficient space, allocate a new arry.
		// Grow by doubling, for amortized linear complexity.
		xlen := len(x)
		zcap := zlen
		for zcap < (2 * xlen) {
			xlen = 2 * xlen
			zcap = xlen
		}

		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

// nonempty is an example of an in-place slice algorithm
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func noempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func dedup() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup : %v\n", err)
		os.Exit(1)
	}
}

// Charcount computes counts of Unicode characters
func charcount() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characteres\n", invalid)
	}
}

var graph = make(map[string]map[string]bool)

func AddEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}
func hasEdge(from, to string) bool {
	return graph[from][to]
}

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

type Point struct{ X, Y int }

func compare() {
	p := Point{1, 2}
	q := Point{1, 2}
	fmt.Println(p == q)
}

type circle struct {
	Point
	Radius int
}
type Wheel struct {
	circle
	Spokes int
}

func embed() {
	w := Wheel{circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		circle: circle{
			Point: Point{X: 8, Y: 8},
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w)
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func movie() {
	// data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

const IssueURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // In markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func searchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func SearchIssues() {
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

const templ = `{{.TotalCount}} issues:
{{range.Items}}-----------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}
`

// because templates are usually fixed at compile time,
// failure to parse a template indicates a fatal bug in the program.
// The template.Must helper function makes error handling more convenient
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func issuesReport() {
	// report, err := template.New("report").
	// 	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	// 	Parse(templ)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// result, err := searchIssues(os.Args[1:])
	result, err := searchIssues([]string{"json", "decode"})
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// var issueList = htmltemplate.Must(htmltemplate.New("issuelist").Parse(`
var issueList = htmltemplate.Must(htmltemplate.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func issuesReportHTML() {
	result, err := searchIssues([]string{"repo:golang/go", "3133", "10535"})
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func autoescape() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := htmltemplate.Must(htmltemplate.New("escape").Parse(templ))
	var data struct {
		A string            // untrusted plain text
		B htmltemplate.HTML // trusted HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
