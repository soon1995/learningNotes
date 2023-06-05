// Construct a tool that reports the set of all packages in the workspace
// that transitively depend on the packages specified by the arguments.
// Hint: you will need to run go list twice, once for the initial packages and once
// for all packages. You may want to parse its JSON output using the encoding/json package
// mostly copied from torbiak/gopl/ex10.4
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: finddeps <pkg>...")
	}
	pkgs := make(map[string]bool)
	for _, arg := range os.Args[1:] {
		cmd := exec.Command("go", "list", arg)
		if _, err := cmd.Output(); err != nil {
			log.Fatalf("package %s invalid: %v", arg, err)
		}
		pkgs[arg] = true
	}

	cmd := exec.Command("go", "list", "-f", `{{.ImportPath}} {{join .Deps " "}}`, "...")
	b, err := cmd.Output()
	if err != nil {
		log.Fatalf("cannot exec command %v; err: %v", cmd.Args, err)
	}
	sc := bufio.NewScanner(bytes.NewReader(b))
	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		pkg := fields[0]
		deps := fields[1:]
		for _, v := range deps {
			if pkgs[v] {
				fmt.Println(pkg)
			}
		}
	}
}
