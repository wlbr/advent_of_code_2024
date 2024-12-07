package main

import (
	"fmt"
	"math"
	"time"
)

type equation struct {
	result   int64
	operands []int64
}

var permutationCache map[int][][]func(int64, int64) int64 = make(map[int][][]func(int64, int64) int64)

func permutate(operators []func(int64, int64) int64, length int) (result [][]func(int64, int64) int64) {
	if p, ok := permutationCache[length]; ok {
		return p
	}
	if length == 1 {
		for _, op := range operators {
			result = append(result, []func(int64, int64) int64{op})
		}
	} else {
		for _, op := range operators {
			for _, p := range permutate(operators, length-1) {
				result = append(result, append(p, op))
			}
		}
	}
	permutationCache[length] = result
	return result
}

func add(a, b int64) int64 {
	return a + b
}

func mul(a, b int64) int64 {
	return a * b
}

func concat(a, b int64) int64 {
	p := int(math.Floor(math.Log10(float64(b))) + 1)
	return a*int64(math.Pow10(p)) + b
}

func task1(binaries []equation) (result int64) {
	for _, eq := range binaries {
		tries := permutate([]func(int64, int64) int64{add, mul}, len(eq.operands)-1)
		for _, t := range tries {
			r := eq.operands[0]
			for i, op := range t {
				r = op(r, eq.operands[i+1])
			}
			if r == eq.result {
				result += eq.result
				break
			}
		}
	}
	return result
}

func task2(binaries []equation) (result int64) {
	permutationCache = make(map[int][][]func(int64, int64) int64)
	for _, eq := range binaries {
		tries := permutate([]func(int64, int64) int64{add, mul, concat}, len(eq.operands)-1)
		for _, t := range tries {
			r := eq.operands[0]
			for i, op := range t {
				r = op(r, eq.operands[i+1])
			}
			if r == eq.result {
				result += eq.result
				break
			}
		}
	}
	return result
}

func main() {
	input := "input.txt"

	data := readdata(input)
	start := time.Now()
	result := task1(data)
	fmt.Printf("Task 1 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)

	start = time.Now()
	result = task2(data)
	fmt.Printf("Task 2 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)

}
