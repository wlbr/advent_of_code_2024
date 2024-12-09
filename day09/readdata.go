package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readdata(input string) (diskLayout []int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		count := 0
		for i, c := range line {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatalf("Error converting '%s' to int: %s", string(c), err)
			}

			unfold := make([]int, n)
			for j := 0; j < n; j++ {
				if i%2 == 0 {
					unfold[j] = count
				} else {
					unfold[j] = -1
				}
			}
			diskLayout = append(diskLayout, unfold...)
			if i%2 == 0 {
				count++
			}
		}
	}

	return diskLayout
}
