package main

import (
	"bufio"
	"log"
	"os"
)

func readdata(input string) (board [][]rune) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		runeLine := make([]rune, len(line))
		for i, c := range line {
			runeLine[i] = c
		}
		board = append(board, runeLine)
	}

	return board
}
