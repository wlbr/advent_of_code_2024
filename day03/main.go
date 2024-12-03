package main

import (
	"fmt"
	"time"
)

type command struct {
	x   int
	y   int
	cmd string
}

func task1(fname string) (result int) {
	program := readdata(fname)
	for _, cmd := range program {
		result += cmd.x * cmd.y
	}
	return result
}

func task2(fname string) (result int) {
	program := readdata(fname)
	active := true
	for _, cmd := range program {
		switch cmd.cmd {
		case "do":
			active = true
		case "don't":
			active = false
		case "mul":
			if active {
				result += cmd.x * cmd.y
			}
		}
	}
	return result
}

func main() {
	input := "input.txt"

	start := time.Now()
	result := task1(input)
	fmt.Printf("Task 1 - elapsed Time: %12s   - result \t = %10d \n", time.Since(start), result)

	start = time.Now()
	result = task2(input)
	fmt.Printf("Task 2 - elapsed Time: %12s   - result \t = %10d \n", time.Since(start), result)

}
