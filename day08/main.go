package main

import (
	"fmt"
	"time"
)

type Point struct {
	x int
	y int
}

func getSymetricOfARelativeToB(a Point, b Point) Point {
	// (a.x + s.x) / 2 = b.x => a.x + s.x = 2*b.x => s.x = 2*b.x - a.x
	// (a.y + s.y) / 2 = b.y => a.y + s.y = 2*b.y => s.y = 2*b.y - a.y

	return Point{2*b.x - a.x, 2*b.y - a.y}
}

func testPointIsOnLineAB(a Point, b Point, c Point) bool {
	/*
	   y - yA = m(x - xA)
	   m = (yB - yA)/(xB - xA)

	   => y - yA - m(x - xA) = 0
	   => y - yA - (yB - yA)(x - xA)/(xB-xA) = 0
	*/

	ybya := b.y - a.y
	xxa := c.x - a.x
	xbxa := b.x - a.x

	return float64(c.y-a.y)-float64(ybya)*float64(xxa)/float64(xbxa) == 0
}

func outOfBounds(board [][]rune, p Point) bool {
	return p.x < 0 || p.x >= len(board) || p.y < 0 || p.y >= len(board[0])
}

func collectAntenas(board [][]rune) map[rune][]Point {
	res := map[rune][]Point{}
	for idx := 0; idx < len(board); idx++ {
		for jdx := 0; jdx < len(board[0]); jdx++ {
			charAt := board[idx][jdx]
			if charAt == '.' {
				continue
			}
			if val, ok := res[charAt]; ok {
				val = append(val, Point{idx, jdx})
				res[charAt] = val
			} else {
				res[charAt] = []Point{{idx, jdx}}
			}
		}
	}
	return res
}

func task1(board [][]rune) int {
	antenas := collectAntenas(board)
	seen := map[Point]bool{}
	for _, points := range antenas {
		for idx := 0; idx < len(points)-1; idx++ {
			for jdx := idx + 1; jdx < len(points); jdx++ {
				iSimJ := getSymetricOfARelativeToB(points[idx], points[jdx])
				jSimI := getSymetricOfARelativeToB(points[jdx], points[idx])

				if !outOfBounds(board, iSimJ) {
					if _, ok := seen[iSimJ]; !ok {
						seen[iSimJ] = true
					}
				}

				if !outOfBounds(board, jSimI) {
					if _, ok := seen[jSimI]; !ok {
						seen[jSimI] = true
					}
				}
			}
		}
	}

	return len(seen)
}

func task2(board [][]rune) int {
	antenas := collectAntenas(board)
	seen := map[Point]bool{}
	for _, points := range antenas {
		for idx := 0; idx < len(points)-1; idx++ {
			for jdx := idx + 1; jdx < len(points); jdx++ {
				ax, bx := points[idx].x, points[jdx].x
				ay, by := points[idx].y, points[jdx].y

				cx := ax
				cy := ay
				for !outOfBounds(board, Point{cx, cy}) && testPointIsOnLineAB(points[idx], points[jdx], Point{cx, cy}) {
					if _, ok := seen[Point{cx, cy}]; !ok {
						seen[Point{cx, cy}] = true
					}
					cx -= bx - ax
					cy -= by - ay
				}

				dx := bx
				dy := by
				for !outOfBounds(board, Point{dx, dy}) && testPointIsOnLineAB(points[idx], points[jdx], Point{dx, dy}) {
					if _, ok := seen[Point{dx, dy}]; !ok {
						seen[Point{dx, dy}] = true
					}
					dx += bx - ax
					dy += by - ay
				}
			}
		}
	}

	return len(seen)
}

func main() {
	input := "input.txt"

	board := readdata(input)
	start := time.Now()
	result := task1(board)
	fmt.Printf("Task 1 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)

	start = time.Now()
	result = task2(board)
	fmt.Printf("Task 2 - elapsed Time: %s - result \t = %d \n", time.Since(start), result)

}
