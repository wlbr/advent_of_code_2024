package main

import (
	"fmt"
	"time"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func checksum(board []int) (sum int) {
	for i, c := range board {
		if c != -1 {
			sum += c * i
		}
	}
	return sum
}

func compact(board []int) []int {
	for i := len(board) - 1; i >= 0; i-- {
		if board[i] != -1 {
			for j := 0; j < i; j++ {
				if board[j] == -1 {
					board[j] = board[i]
					board[i] = -1
					break
				}
			}
		}
	}
	return board
}

func findGap(board []int, comparator, length int) int {
	for i := 0; i < len(board)-length+1; i++ {
		found := true
		for j := 0; j < length; j++ {
			if board[i+j] != comparator {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}

func findNextSubSliceFromEnd(board []int, pos int) (elem, start, length int) {
	if pos < 0 {
		pos = int(len(board) - 1)
	}
	for i := pos; i >= 0; i-- {
		comparator := board[i]
		if comparator != -1 {
			for j := 1; j < len(board); j++ {
				if i-int(j) < 0 || board[i-int(j)] != comparator {
					length = int(j)
					start = i
					return comparator, start, length
				}
			}
		}
	}
	return -1, -1, -1
}

func compactCompleteFiles(board []int) []int {
	moved := make(map[int]bool)
	var comparator, pos, length int = 0, -1, 0
	for {
		comparator, pos, length = findNextSubSliceFromEnd(board, pos-length)
		if comparator != -1 {
			if _, ok := moved[comparator]; !ok {
				gstart := findGap(board[:max(pos-(length-1), 0)], -1, length)
				if gstart != -1 {
					for i := pos; i >= pos-length; i-- {
						if board[i] != -1 {
							moved[comparator] = true
							for j := gstart; j < gstart+length; j++ {
								if board[j] == -1 {
									board[j] = board[i]
									board[i] = -1
									i--
								}
							}
						}
					}
				}
			}
		}
		if comparator == -1 || pos-length < 0 {
			break
		}
	}
	return board
}

func task1(input []int) int {
	compacted := compact(input)
	return checksum(compacted)
}

func task2(input []int) int {
	compacted := compactCompleteFiles(input)
	return checksum(compacted)
}

func main() {
	input := "input.txt"

	board := readdata(input)
	start := time.Now()
	result := task1(board)
	fmt.Printf("Task 1 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)

	board = readdata(input)
	start = time.Now()
	result = task2(board)
	fmt.Printf("Task 1 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)

}
