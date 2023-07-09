package main

import (
	"log"
	"os"
	"strconv"

	"example.com/model"
)

func main() {
	var count int64
	if len(os.Args) < 3 {
		log.Fatalf("GumballMachine <name> <inventory>")
	}

	count, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatalf("GumballMachine <name> <inventory>, inventory must be number")
	}

	machine := model.NewGumballMachine(os.Args[1], int(count))
	monitor := model.NewGumballMonitor(machine)

	monitor.Report()
}
