package main

import (
	"fmt"
	"time"
)

// Minf64 returns the minimum of a slice of float64
func Min(v []int) int {
	m := v[0]
	for _, e := range v {
		if e < m {
			m = e
		}
	}
	return m
}

// Maxf64 returns the maximum of a slice of float64
func Max(v []int) int {
	m := v[0]
	for _, e := range v {
		if e > m {
			m = e
		}
	}
	return m
}

func checkStetigSteigend(report []int, mindelta, maxdelta int) bool {
	for i := 1; i < len(report); i++ {
		if report[i]-report[i-1] < mindelta || report[i]-report[i-1] > maxdelta {
			return false
		}
	}
	return true
}

func checkStetigFallend(report []int, mindelta, maxdelta int) bool {
	for i := 1; i < len(report); i++ {
		if report[i-1]-report[i] < mindelta || report[i-1]-report[i] > maxdelta {
			return false
		}
	}
	return true
}

func task1(reports [][]int) (result int) {
	for _, report := range reports {
		if checkStetigSteigend(report, 1, 3) || checkStetigFallend(report, 1, 3) {
			result++
		}
	}

	return result
}

func task2(reports [][]int) (result int) {
	return result
}

func main() {
	input := "input.txt"

	data := readdata(input)
	start := time.Now()
	result := task1(data)
	fmt.Printf("Task 1 - elapsed Time: %s \t - result \t = %d \n", time.Since(start), result)

	start = time.Now()
	result = task2(data)
	fmt.Printf("Task 2 - elapsed Time: %s \t - result \t = %d \n", time.Since(start), result)

}
