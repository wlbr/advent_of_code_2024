package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readdata(input string) *area {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	var sarea []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		sarea = append(sarea, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	area := newArea(len(sarea[0]), len(sarea))
	for y, line := range sarea {
		for x, c := range line {
			area.set(x, y, string(c))
		}
	}

	return area
}
