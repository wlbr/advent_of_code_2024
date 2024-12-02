package main

import (
	"fmt"
	"time"
)

const MINDELTA = 1
const MAXDELTA = 3

func remove(slice []int, s int) []int {
	tmp := make([]int, len(slice))
	copy(tmp, slice)
	return append(tmp[:s], tmp[s+1:]...)
}

func checkRising(a, b int) bool {
	if b-a < MINDELTA || b-a > MAXDELTA {
		return false
	}
	return true
}

func checkFalling(a, b int) bool {
	if a-b < MINDELTA || a-b > MAXDELTA {
		return false
	}
	return true
}

func checkReport(report []int) bool {
	var isRising, isFalling bool = true, true
	for i := 1; i < len(report); i++ {
		r := checkRising(report[i-1], report[i])
		f := checkFalling(report[i-1], report[i])
		if !r {
			isRising = false
		}
		if !f {
			isFalling = false
		}

	}
	if isRising || isFalling {
		return true
	}
	return false
}

func checkFixes(report []int, damping bool) bool {
	if checkReport(report) {
		return true
	}
	if damping {
		for j := 0; j < len(report); j++ {
			fix := remove(report, j)
			if checkReport(fix) {
				return true
			}
		}
	}
	return false
}

func task1(reports [][]int) (result int) {
	for _, report := range reports {
		//if checkStetigSteigend(report, false) || checkStetigFallend(report, false) {
		if checkFixes(report, false) {
			result++
		}
	}

	return result
}

func task2(reports [][]int) (result int) {
	for _, report := range reports {
		if checkFixes(report, true) {
			result++
		}
	}
	return result
}

func main() {
	input := "input.txt"

	data := readdata(input)
	start := time.Now()
	result := task1(data)
	fmt.Printf("Task 1 - elapsed Time: %14s \t - number of safe reports \t = %d \n", time.Since(start), result)

	start = time.Now()
	result = task2(data)
	fmt.Printf("Task 2 - elapsed Time: %14s \t - number of safe reports \t = %d \n", time.Since(start), result)

}
