package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Echo4()
	// Exception()
	// Cf()
  Popcount()
}

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "seperator")

// Echo4 prints its command-line arguments
func Echo4() {
	// MUST called before flags are used
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

// go run main/main.go -s "/" -n  asdfsadf asdfasdf asdfasdf asdfasdf adsf asdf adsf adf asdf
// Result:
// asdfsadf/asdfasdf/asdfasdf/asdfasdf/adsf/asdf/adsf/adf/asdf%

// go run main/main.go -h or --help
// Usage of /tmp/go-build4258876248/b001/exe/main:
//  -n    omit trailing newline
//  -s string
//        seperator (default " ")

func Exception() {
	fmt.Println([1]int{1} == [1]int{1})
	fmt.Println(struct{}{} == struct{}{})
	fmt.Println(struct{ hello string }{} == struct{ hello string }{})
}

// Package tempconv performs Celsius and Fahrenheit temperature computations.
type Celsius float64

type Fahrenheit float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

const (
	AbsoulateZeroC Celsius = -273.15
	FreezingC      Celsius = 0
	BoilingC       Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Cf converts its numeric argument to Celsius and Fahrenheit.
func Cf() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf:  %v\n", err)
			os.Exit(1)
		}
		f := Fahrenheit(t)
		c := Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, FtoC(f), c, CToF(c))
	}
}

// pc[i] is the population count of i
var pc [256]byte

func Popcount() {
  fmt.Println(1>>0*8)
  fmt.Println(byte(1>>0*8))
	for i := range pc {
    // fmt.Println("pc i/2")
    // fmt.Println(pc[i/2])
    // fmt.Println("i&1")
    // fmt.Println(byte(i&1))
		pc[i] = pc[i/2] + byte(i&1)
	}
	fmt.Println(pc)
  fmt.Println(popCount(256))
}

func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
