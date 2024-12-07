package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func safeAtoi(nums []string) (int64s []int64) {

	for _, n := range nums {
		i, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalf("Error converting '%v' coords to int64: %s from %v", n, nums)
		}
		int64s = append(int64s, int64(i))
	}
	return int64s
}

func readdata(input string) (equations []equation) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		eq := equation{}
		line := scanner.Text()
		strresult := strings.Split(line, ": ")
		tmp, err := strconv.Atoi(strresult[0])
		eq.result = int64(tmp)
		if err != nil {
			log.Fatalf("Error converting result '%s' to int64: %s", strresult[0], err)
		}
		stroperands := strings.Split(strresult[1], " ")
		eq.operands = safeAtoi(stroperands)
		equations = append(equations, eq)
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error reading standard input:", err)
	}
	return equations
}
