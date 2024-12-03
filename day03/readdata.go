package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readfile(input string) (lines []string) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return lines
}

func safeAtoi(nums []string) (ints []int) {

	for _, n := range nums {
		i, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalf("Error converting command '%v' coords to int: %s", nums, err)
		}
		ints = append(ints, i)
	}
	return ints
}

var rex *regexp.Regexp = regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d+),(\d+)\))`)

func readdata(filename string) (program []*command) {
	lines := readfile(filename)
	for _, line := range lines {
		matches := rex.FindAllStringSubmatch(line, -1)
		for _, m := range matches {
			if m[1] == "do()" {
				program = append(program, &command{cmd: "do"})
				continue
			} else if m[1] == "don't()" {
				program = append(program, &command{cmd: "don't"})
				continue
			} else if m[1][:3] == "mul" {
				nums := safeAtoi(m[2:])
				program = append(program, &command{cmd: "mul", x: nums[0], y: nums[1]})
				continue
			}
		}
	}

	return program
}
