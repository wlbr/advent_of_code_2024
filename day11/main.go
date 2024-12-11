package main

import (
	"fmt"
	"strconv"
	"time"
)

func copyMap(m map[int64]int64) map[int64]int64 {
	n := make(map[int64]int64)
	for k, v := range m {
		n[k] = v
	}
	return n
}

func oneBlink(stones map[int64]int64) map[int64]int64 {
	stonework := copyMap(stones)
	for stone, count := range stonework {
		if count == 0 {
			continue
		}
		if stone == 0 {
			stones[1] += count
			stones[0] -= count
		} else if stone_str := fmt.Sprint(stone); len(stone_str)%2 == 0 {
			new_len := int(len(stone_str) / 2)
			stone_1, _ := strconv.ParseInt(stone_str[:new_len], 10, 0)
			stone_2, _ := strconv.ParseInt(stone_str[new_len:], 10, 0)
			stones[stone_1] += count
			stones[stone_2] += count
			stones[stone] -= count
		} else {
			stones[stone*2024] += count
			stones[stone] -= count
		}
	}
	return stones
}

func allBlinks(stones map[int64]int64, max int) map[int64]int64 {
	for i := 0; i < max; i++ {
		stones = oneBlink(stones)
	}
	return stones
}

func tasks(binaries []int64, max int) (result int64) {
	stones := make(map[int64]int64)
	for _, s := range binaries {
		stones[s]++
	}
	stones = allBlinks(stones, max)
	for _, v := range stones {
		result += v
	}
	return result
}

func main() {
	input := "input.txt"

	data := readdata(input)
	start := time.Now()
	result := tasks(data, 25)
	fmt.Printf("Task 1 - elapsed Time: %12s   - result \t = %d \n", time.Since(start), result)

	data = readdata(input)
	start = time.Now()
	result = tasks(data, 75)
	fmt.Printf("Task 2 - elapsed Time: %12s   - result \t = %d \n", time.Since(start), result)

}
