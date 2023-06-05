// Write a program that reads a single expression from the standard input,
// prompts the user to provide values for any variables, then evaluates the expression
// in the resulting environment. Handle all errors gracefully.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const assignment_error = 2

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Printf("enter expression: ")
	sc.Scan()
	exprStr := sc.Text()
	fmt.Printf("variable <var>=<val> eg. x=12 y=13: ")
	sc.Scan()
	envStr := sc.Text()
	if sc.Err() != nil {
		log.Fatal(sc.Err())
	}

	env := Env{}
	exitCode := 0
	assignments := strings.Fields(envStr)
	for _, a := range assignments {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad assignment: %s\n", a)
			exitCode = assignment_error
		}
		if len(fields) == 1 {
			fields = append(fields, "0")
		}
		ident, valStr := fields[0], fields[1]
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bad value for %s, using zero: %s\n", ident, err)
			exitCode = assignment_error
		}
		env[Var(ident)] = val
	}

	expr, err := Parse(exprStr)
	if err != nil {
		log.Fatalf("bad expression: %s", err)
	}
	fmt.Println(expr.Eval(env))
	os.Exit(exitCode)
}
