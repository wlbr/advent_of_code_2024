package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readdata(input string) (binaries []int64) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		snums := strings.Split(line, " ")
		for _, snum := range snums {
			num, err := strconv.ParseInt(snum, 10, 0)
			if err != nil {
				log.Fatalf("Error converting string '%s' to number: %s", snum, err)
			}
			binaries = append(binaries, num)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return binaries
}
