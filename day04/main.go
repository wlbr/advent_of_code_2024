package main

import (
	"fmt"
	"strings"
	"time"
)

func reverse(s string) string {
	n := 0
	rune := make([]rune, len(s))
	for _, r := range s {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)
	return output
}

func flipHorizontally(lines []string) (result []string) {
	for _, line := range lines {
		result = append(result, reverse(line))
	}
	return result
}

func flipvertically(lines []string) (result []string) {
	for i := len(lines) - 1; i >= 0; i-- {
		result = append(result, lines[i])
	}
	return result
}

func rotate90(lines []string) (result []string) {
	for i := 0; i < len(lines[0]); i++ {
		var line string
		for j := len(lines) - 1; j >= 0; j-- {
			line += string(lines[j][i])
		}
		result = append(result, line)
	}
	return result
}

func rotate270(lines []string) (result []string) {
	for i := len(lines[0]) - 1; i >= 0; i-- {
		var line string
		for j := 0; j < len(lines); j++ {
			line += string(lines[j][i])
		}
		result = append(result, line)
	}
	return result
}

func rotate45(lines []string) (result []string) {
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			var line string
			rx := x
			ry := y
			for {
				//fmt.Printf("ii: %d - j: %d  line: '%s' lines[%d]: %s. \n", ii, j, line, ii, lines[ii])
				if ry >= len(lines) || rx >= len(lines[ry]) {
					break
				}
				line = line + string(lines[ry][rx])
				rx++
				ry++
			}
			result = append(result, line)
		}
	}
	return result
}

func rotate315(lines []string) (result []string) {
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			var line string
			rx := x
			ry := y
			for {
				//fmt.Printf("ii: %d - j: %d  line: '%s' lines[%d]: %s. \n", ii, j, line, ii, lines[ii])
				if ry >= len(lines) || rx < 0 {
					break
				}
				line = line + string(lines[ry][rx])
				rx--
				ry++
			}
			result = append(result, line)
		}
	}
	return result
}

func count(lines []string) (result int) {
	for _, line := range lines {
		result += strings.Count(line, "XMAS")
	}
	return result
}

func show(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func compare(left, right []string) {
	m := max(len(left), len(right))
	for i := 0; i < m; i++ {
		if i < len(left) {
			fmt.Printf("  %s", left[i])
		} else {
			fmt.Printf(" \t")
		}
		if i < len(right) {
			fmt.Printf("\t %s  \n", right[i])
		}
	}
	fmt.Println()
}

func task1(cmds []string) (result int) {
	show(cmds)
	result = count(cmds)

	compare(cmds, flipHorizontally(cmds))
	result += count(flipHorizontally(cmds))

	compare(cmds, rotate90(cmds))
	result += count(rotate90(cmds))

	compare(cmds, rotate270(cmds))
	result += count(rotate270(cmds))

	compare(cmds, rotate45(cmds))
	result += count(rotate45(cmds))

	compare(cmds, rotate315(cmds))
	result += count(rotate315(cmds))

	return result
}

func task2(cmds []string) (result int) {

	return 12
}

func main() {
	input := "example1.txt"

	data := readfile(input)
	start := time.Now()
	result := task1(data)
	fmt.Printf("Task 1 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)

	start = time.Now()
	result = task2(data)
	fmt.Printf("Task 2 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)

}
