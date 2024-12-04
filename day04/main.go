package main

import (
	"fmt"
	"time"
)

const SEARCHWORD1 = "XMAS"
const SEARCHWORD2 = "MAS"

func outOfBounds(lines []string, x, y int) bool {
	return !(y >= 0 && y < len(lines) && x >= 0 && x < len(lines[y]))
}

func searchRight(lines []string, x, y int, word string) bool {
	for i := 0; i < len(word); i++ {
		if outOfBounds(lines, x+i, y) || lines[y][x+i] != word[i] {
			return false
		}
	}
	return true
}

func searchLeft(lines []string, x, y int, word string) bool {
	for i := 0; i < len(word); i++ {
		if outOfBounds(lines, x-i, y) || lines[y][x-i] != word[i] {
			return false
		}
	}
	return true
}

func searchDown(lines []string, x, y int, word string) bool {
	for i := 0; i < len(word); i++ {
		if outOfBounds(lines, x, y+i) || lines[y+i][x] != word[i] {
			return false
		}
	}
	return true
}

func searchUp(lines []string, x, y int, word string) bool {
	for i := 0; i < len(word); i++ {
		if outOfBounds(lines, x, y-i) || lines[y-i][x] != word[i] {
			return false
		}
	}
	return true
}

func searchDownRight(lines []string, x, y int, word string) bool {
	for i := 0; i < len(word); i++ {
		if outOfBounds(lines, x+i, y+i) || lines[y+i][x+i] != word[i] {
			return false
		}
	}
	return true
}

func searchDownLeft(lines []string, x, y int, word string) bool {
	for i := 0; i < len(word); i++ {
		if outOfBounds(lines, x-i, y+i) || lines[y+i][x-i] != word[i] {
			return false
		}
	}
	return true
}

func searchUpRight(lines []string, x, y int, word string) bool {
	for i := 0; i < len(word); i++ {
		if outOfBounds(lines, x+i, y-i) || lines[y-i][x+i] != word[i] {
			return false
		}
	}
	return true
}

func searchUpLeft(lines []string, x, y int, word string) bool {
	for i := 0; i < len(word); i++ {
		if outOfBounds(lines, x-i, y-i) || lines[y-i][x-i] != word[i] {
			return false
		}
	}
	return true
}

func task1(lines []string) (result int) {
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if searchRight(lines, x, y, SEARCHWORD1) {
				result++
			}
			if searchLeft(lines, x, y, SEARCHWORD1) {
				result++
			}
			if searchDown(lines, x, y, SEARCHWORD1) {
				result++
			}
			if searchUp(lines, x, y, SEARCHWORD1) {
				result++
			}
			if searchDownRight(lines, x, y, SEARCHWORD1) {
				result++
			}
			if searchDownLeft(lines, x, y, SEARCHWORD1) {
				result++
			}
			if searchUpRight(lines, x, y, SEARCHWORD1) {
				result++
			}
			if searchUpLeft(lines, x, y, SEARCHWORD1) {
				result++
			}
		}
	}
	return result
}

func checkMAS(lines []string, x, y int) bool {
	word := SEARCHWORD2
	dist := len(word) % 2
	if dist == 0 {
		fmt.Print("Error - searchword not of uneven length.")
		return false
	}
	if (searchDownRight(lines, x-dist, y-dist, word) || searchUpLeft(lines, x+dist, y+dist, word)) &&
		(searchDownLeft(lines, x+dist, y-dist, word) || searchUpRight(lines, x-dist, y+dist, word)) {
		return true
	}
	return false
}

func task2(lines []string) (result int) {
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if checkMAS(lines, x, y) {
				result++
			}
		}
	}
	return result
}

func main() {
	input := "input.txt"

	data := readfile(input)
	start := time.Now()
	result := task1(data)
	fmt.Printf("Task 1 - elapsed Time: %12s  \t %5s occurence \t= %d \n", time.Since(start), SEARCHWORD1, result)

	start = time.Now()
	result = task2(data)
	fmt.Printf("Task 2 - elapsed Time: %12s  \t %5s occurence \t= %d \n", time.Since(start), SEARCHWORD2, result)

}
