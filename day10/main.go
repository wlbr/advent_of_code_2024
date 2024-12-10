package main

import (
	"fmt"
	"time"
)

type position struct {
	x int
	y int
}

func (p position) String() string {
	return fmt.Sprintf("[%d,%d]", p.x, p.y)
}

func searchAllStartPositions(board [][]int) []position {
	var startPositions []position

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				startPositions = append(startPositions, position{j, i})
			}
		}
	}

	return startPositions
}

func getNeighbors(board [][]int, pos position) []position {
	neighbors := []position{
		{pos.x - 1, pos.y},
		{pos.x + 1, pos.y},
		{pos.x, pos.y - 1},
		{pos.x, pos.y + 1},
	}

	var validNeighbors []position
	for _, n := range neighbors {
		if n.x >= 0 && n.x < len(board) && n.y >= 0 && n.y < len(board[0]) {
			validNeighbors = append(validNeighbors, n)
		}
	}

	return validNeighbors
}

func traverseBoard(board [][]int, pos position, visited map[string]int) (score int) {

	neighbors := getNeighbors(board, pos)
	p := board[pos.y][pos.x]
	if p == 9 {
		if visited != nil {
			if visited[pos.String()] > 0 {
				return 0
			}
			visited[pos.String()]++
		}
		return 1
	}
	for _, n := range neighbors {
		m := board[n.y][n.x]
		t := p + 1
		if m == t {
			score += traverseBoard(board, n, visited)
		}
	}

	return score
}

func task1(input [][]int) (score int) {
	starts := searchAllStartPositions(input)
	for _, start := range starts {
		score += traverseBoard(input, start, make(map[string]int))
	}
	return score
}

func task2(input [][]int) (score int) {
	starts := searchAllStartPositions(input)
	for _, start := range starts {
		score += traverseBoard(input, start, nil)
	}
	return score
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
	fmt.Printf("Task 2 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)

}
