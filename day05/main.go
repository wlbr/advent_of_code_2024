package main

import (
	"fmt"
	"log"
	"slices"
	"time"
)

type rule struct {
	predecessor int
	successor   int
}

type update []int

func compressRules(rules []rule) map[int][]int {
	compressed := make(map[int][]int)
	for _, r := range rules {
		compressed[r.predecessor] = append(compressed[r.predecessor], r.successor)
	}
	return compressed
}

func evenp(n int) bool {
	return n%2 == 0
}

func sumMiddleNumber(nums []update) (sum int) {
	for _, n := range nums {
		if evenp(len(n)) {
			log.Printf("Error:  even number of elements in slice")
		}
		if len(n) == 1 {
			sum += n[0]
		} else {
			sum += n[len(n)/2]
		}
	}

	return sum
}

func task1(rules []rule, updates []update) (result int) {
	var validresults []update
	crules := compressRules(rules)

	for _, update := range updates {
		valid := true
		for pos, n := range update {
			if pos == 0 || pos == len(update) {
				continue
			}
			affectedRules := crules[n]
			for _, r := range affectedRules {
				if slices.Contains(update[:pos], r) {
					valid = false
					break
				}
			}
		}
		if valid {
			validresults = append(validresults, update)
		}
	}

	return sumMiddleNumber(validresults)
}

func task2(rules []rule, updates []update) (result int) {
	var inValidResults []update
	crules := compressRules(rules)

	for _, update := range updates {
		inValid := false
		pos := 0

		for {
			n := update[pos]
			affectedRules := crules[n]
			for _, r := range affectedRules {
				violationpos := slices.Index(update[:pos], r)
				if violationpos != -1 {
					inValid = true
					tmp := update[pos]
					update[pos] = update[violationpos]
					update[violationpos] = tmp
					pos = 0
				}
			}
			pos++
			if pos >= len(update) {
				break
			}
		}
		if inValid {
			inValidResults = append(inValidResults, update)
		}
	}
	return sumMiddleNumber(inValidResults)

}

func main() {
	input := "input.txt"

	rules, updates := readdata(input)
	start := time.Now()
	result := task1(rules, updates)
	fmt.Printf("Task 1 - elapsed Time: %12s   - sum of  middle page numbers      \t = %d \n", time.Since(start), result)

	start = time.Now()
	result = task2(rules, updates)
	fmt.Printf("Task 2 - elapsed Time: %12s   - sum of fixed middle page numbers \t = %d \n", time.Since(start), result)

}
