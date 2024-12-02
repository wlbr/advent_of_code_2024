package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readdata(input string) (reports [][]int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		var report []int
		words := strings.Split(line, " ")
		for _, w := range words {
			var c int
			count, err := fmt.Sscan(w, &c)
			if err != nil || count == 0 {
				break
			}
			report = append(report, c)
		}
		reports = append(reports, report)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return reports
}
