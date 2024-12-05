package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func readdata(input string) (rules []rule, updates []update) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	part1 := true
	for scanner.Scan() {
		line := scanner.Text()
		if part1 {
			if line != "" {
				stringnums := strings.Split(line, "|")
				nums := safeAtoi(stringnums)
				rule := rule{predecessor: nums[0], successor: nums[1]}
				rules = append(rules, rule)
			} else {
				part1 = false
			}
		} else {
			stringnums := strings.Split(line, ",")
			nums := safeAtoi(stringnums)
			update := update(nums)
			updates = append(updates, update)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading standard input: %v", err)
	}
	return rules, updates
}
