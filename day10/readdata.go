package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readdata(input string) (topoMap [][]int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		var intline []int
		for _, c := range line {
			var n int
			if c == '.' {
				n = -10
			} else {
				n, err = strconv.Atoi(string(c))
				if err != nil {
					log.Fatalf("Error converting '%s' to int: %s", string(c), err)
				}
			}
			intline = append(intline, n)
		}
		topoMap = append(topoMap, intline)
	}
	return topoMap
}
