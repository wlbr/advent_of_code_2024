package main

import (
	"bufio"
	"log"
	"os"
)

func readdata(input string) (terra terrain) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		terra = append(terra, runes)
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading standard input: %v", err)
	}
	return terra
}
