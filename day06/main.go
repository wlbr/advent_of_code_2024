package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/alitto/pond"
)

type terrain [][]rune

func (t terrain) String() string {
	var s string
	for _, line := range t {
		s += string(line) + "\n"
	}
	return s
}

func searchStartPosition(terra terrain) (x, y int) {
	for y, line := range terra {
		for x, field := range line {
			if strings.Contains("^<>v", string(field)) {
				return x, y
			}
		}
	}
	log.Fatal("No start position found")
	return -1, -1
}

func turn(field rune) rune {
	switch field {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	default:
		log.Fatalf("Invalid turn from direction: %c", field)
		return '-'
	}
}

func outOfBounds(lines terrain, x, y int) bool {
	return !(y >= 0 && y < len(lines) && x >= 0 && x < len(lines[y]))
}

func walk(terra terrain) (visitedPositions map[string]int, loop bool) {

	visitedPositions = make(map[string]int)
	startx, starty := searchStartPosition(terra)
	visitedPositions[fmt.Sprintf("%d-%d", startx, starty)] = 1

	x := startx
	y := starty

	for {

		direction := terra[y][x]
		lastx, lasty := x, y
		switch direction {
		case '^':
			y--
		case '>':
			x++
		case 'v':
			y++
		case '<':
			x--
		}
		if outOfBounds(terra, x, y) {
			break
		}
		terra[lasty][lastx] = '.' //'x'

		if terra[y][x] == '.' || terra[y][x] == 'x' {
			terra[y][x] = direction
			visitedPositions[fmt.Sprintf("%d-%d", x, y)]++
			if visitedPositions[fmt.Sprintf("%d-%d", x, y)] > 4 {
				loop = true
				break
			}
		}
		if terra[y][x] == '#' || terra[y][x] == 'O' {
			x, y = lastx, lasty
			terra[y][x] = turn(direction)
		}
	}
	return visitedPositions, loop
}

func copyTerrain(terra terrain) (copy terrain) {
	for _, line := range terra {
		copy = append(copy, append([]rune{}, line...))
	}
	return
}

func task1(terra terrain) (result int) {
	visitedPositions, _ := walk(terra)
	return len(visitedPositions)
}

func task2(terra terrain) (result int) {
	maxproc := runtime.NumCPU()
	runtime.GOMAXPROCS(maxproc)

	pool := pond.New(maxproc, len(terra))
	mux := &sync.RWMutex{}

	for y, line := range terra {
		for x, field := range line {
			if field == '.' || field == 'x' {
				ct := copyTerrain(terra)
				ct[y][x] = 'O'
				pool.Submit(func() {
					_, loop := walk(copyTerrain(ct))
					if loop {
						mux.Lock()
						result++
						mux.Unlock()
					}
				})
			}
		}
	}
	pool.StopAndWait()
	return result
}

func main() {
	input := "input.txt"

	terrain := readdata(input)
	start := time.Now()
	result := task1(terrain)
	fmt.Printf("Task 1 - elapsed Time: %12s   - count of unique positons                      = %d \n", time.Since(start), result)

	terrain = readdata(input)
	start = time.Now()
	result = task2(terrain)
	fmt.Printf("Task 2 - elapsed Time: %12s   - count of different positions for obstructions = %d \n", time.Since(start), result)

}
