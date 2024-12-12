package main

import (
	"flag"
	"fmt"
	"time"
)

func printf(format string, args ...interface{}) {
	if flag.Lookup("test.v") != nil {
		fmt.Printf(format, args...)
	}
}

func print(args ...interface{}) {
	if flag.Lookup("test.v") != nil {
		fmt.Print(args...)
	}
}

func println(args ...interface{}) {
	if flag.Lookup("test.v") != nil {
		fmt.Println(args...)
	}
}

func inspect(board []string) {
	if flag.Lookup("test.v") != nil {
		for _, line := range board {
			println(line)
		}
	}
}

//-------------------------------------------------------------------------

type position struct {
	x       int
	y       int
	str     string
	val     string
	visited bool
	region  *region
}

func newPosition(x, y int, str string) *position {
	p := &position{}
	p.set(x, y, str)
	return p
}

func (p *position) String() string {
	return p.str
}

func (p *position) setPosition(x, y int) {
	p.x = x
	p.y = y
	p.str = fmt.Sprint("[", x, "-", y, "]")
}

func (p *position) setValue(x, y int, str string) {
	p.val = string(p.str)
}

func (p *position) set(x, y int, str string) {
	p.setPosition(x, y)
	p.val = str
}

func (p *position) get() string {
	return p.val
}

func (p *position) getRegion() *region {
	return p.region
}

func (p *position) setRegion(r *region) {
	p.region = r
}

// -------------------------------------------------------------------------
type area struct {
	dimx      int
	dimy      int
	positions [][]*position
	regions   map[*region]*region
}

func newArea(dimx, dimy int) *area {
	a := &area{dimx, dimy, nil, nil}
	a.positions = make([][]*position, dimy)
	a.regions = make(map[*region]*region)
	for y := range a.positions {
		a.positions[y] = make([]*position, dimx)
		for x := 0; x < dimx; x++ {
			a.positions[y][x] = newPosition(x, y, ".")
		}
	}
	return a
}

func (a area) String() string {
	str := ""
	for y := 0; y < a.dimy; y++ {
		for x := 0; x < a.dimx; x++ {
			str += a.get(x, y)
		}
		str += "\n"
	}
	return str
}

func (a area) getPosition(x, y int) *position {
	return a.positions[y][x]
}

func (a area) setPosition(x, y int, p *position) {
	a.positions[y][x] = p
	p.setPosition(x, y)
}

func (a area) get(x, y int) string {
	return a.positions[y][x].val
}

func (a area) set(x, y int, val string) {
	a.positions[y][x].set(x, y, val)
}

func (a area) setVisited(x, y int) {
	a.positions[y][x].visited = true
}

func (a area) getVisited(x, y int) bool {
	return a.positions[y][x].visited
}

func (a area) getPredecessors(x, y int) (neighboors []*position) {
	if x > 0 {
		neighboors = append(neighboors, a.getPosition(x-1, y))
	}
	// if x < a.dimx-1 {
	// 	neighboors = append(neighboors, a.getPosition(x+1, y))
	// }
	if y > 0 {
		neighboors = append(neighboors, a.getPosition(x, y-1))
	}
	// if y < a.dimy-1 {
	// 	neighboors = append(neighboors, a.getPosition(x, y+1))
	// }
	return neighboors
}

// func (a area) getNeighboors(x, y int) (neighboors []*position) {
// 	// if x > 0 {
// 	neighboors = append(neighboors, a.getPosition(x-1, y))
// 	// }
// 	// if x < a.dimx-1 {
// 	neighboors = append(neighboors, a.getPosition(x+1, y))
// 	// }
// 	// if y > 0 {
// 	neighboors = append(neighboors, a.getPosition(x, y-1))
// 	// }
// 	// if y < a.dimy-1 {
// 	neighboors = append(neighboors, a.getPosition(x, y+1))
// 	// }
// 	// if x > 0 && y > 0 {
// 	neighboors = append(neighboors, a.getPosition(x-1, y-1))
// 	// }
// 	// if x < a.dimx-1 && y > 0 {
// 	neighboors = append(neighboors, a.getPosition(x+1, y-1))
// 	// }
// 	// if x > 0 && y < a.dimy-1 {
// 	neighboors = append(neighboors, a.getPosition(x-1, y+1))
// 	// }
// 	// if x < a.dimx-1 && y < a.dimy-1 {
// 	neighboors = append(neighboors, a.getPosition(x+1, y+1))
// 	// }
// 	return neighboors
// }

func (a *area) addRegion(r *region) {
	a.regions[r] = r
}

func getNeighbors(x, y int) [][]int {
	return [][]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
	// {x - 1, y - 1}, {x + 1, y - 1}, {x - 1, y + 1}, {x + 1, y + 1}}
}

// func (a *area) getRegion(r *region) *region {
// 	return a.regions[r]
// }

func (a *area) delRegion(r *region) bool {
	if _, ok := a.regions[r]; ok {
		delete(a.regions, r)
		return true
	}
	return false
}

func (a *area) getRegionCount() int {
	return len(a.regions)
}

// -------------------------------------------------------------------------

type region struct {
	index     map[*position]*position
	positions [][]*position
	val       string
	capx      int
	capy      int
}

func outOfBounds(x, y, dimx, dimy int) bool {
	return x < 0 || y < 0 || x >= dimx || y >= dimy
}

func newRegion(val string, capx, capy int) *region {
	r := &region{val: val, capx: capx, capy: capy}
	r.positions = make([][]*position, capy)
	r.index = make(map[*position]*position)
	for y := 0; y < capy; y++ {
		r.positions[y] = make([]*position, capx)
	}
	return r
}

func (r *region) String() string {
	str := ""
	for y := 0; y < r.capy; y++ {
		for x := 0; x < r.capx; x++ {
			str += r.get(x, y)
		}
		str += "\n"
	}
	return str
}

func (r *region) add(p *position) {
	// if _, ok := r.index[p]; !ok {
	r.index[p] = p
	r.positions[p.y][p.x] = p
	p.setRegion(r)
	// } else {
	// 	log.Print("Position already in region")
	// }
}

func (r *region) get(x, y int) string {
	if r.positions[y][x] == nil {
		return "."
	}
	return r.positions[y][x].val
}

func (r *region) merge(r2 *region) {
	for _, p := range r2.index {
		r.add(p)
		//p.setRegion(r)
	}
	r2.index = make(map[*position]*position)
	r2.positions = make([][]*position, r2.capy)
	for y := 0; y < r2.capy; y++ {
		r2.positions[y] = make([]*position, r2.capx)
	}
}

func (r *region) getArea() int {
	return len(r.index)
}

func (r *region) getPeremiter(board *area) int {
	var result int
	border := make(map[string]int)
	if r.val == "C" {
		fmt.Println("region ", r.val)
	}
	for _, p := range r.index {
		neighboors := getNeighbors(p.x, p.y)
		for _, npos := range neighboors {
			if outOfBounds(npos[0], npos[1], board.dimx, board.dimy) {
				border[fmt.Sprintf("[%d,%d]", npos[0], npos[1])]++
			} else {
				n := board.getPosition(npos[0], npos[1])
				if n.val != p.val {
					border[fmt.Sprintf("[%d,%d]", npos[0], npos[1])]++
				}
			}
		}
		// if border {
		// 	result++
		// }
	}
	for _, v := range border {
		result += v
	}
	return result
}

//-------------------------------------------------------------------------

func searchSubArea(board *area, x, y int) {
	current := board.getPosition(x, y)
	if !current.visited {
		current.visited = true
		//if current.getRegion() == nil {
		r := newRegion(current.val, board.dimx, board.dimy)
		r.add(current)
		board.addRegion(r)
		//}

		neighboors := board.getPredecessors(x, y)
		for _, n := range neighboors {
			if n.val == current.val {
				if n.visited {
					oldRegion := current.getRegion()
					if n.getRegion() != current.getRegion() {
						n.getRegion().merge(oldRegion)
						board.delRegion(oldRegion)
					}
				} else {
					n.setRegion(current.getRegion())
					//searchSubArea(board, n.x, n.y)
				}
			}
		}
	}
}

func inspectRegions(x, y int, regions map[*region]*region) {
	printf("[%d,%d] -------\n", x, y)
	for _, r := range regions {
		println(r)
	}

}

func searchSubAreas(board *area) {
	for y := 0; y < board.dimy; y++ {
		for x := 0; x < board.dimx; x++ {
			searchSubArea(board, x, y)
			inspectRegions(x, y, board.regions)
		}
	}
}

func task1(board *area) (result int) {
	println(board)
	println()
	searchSubAreas(board)

	for _, r := range board.regions {
		p := r.getPeremiter(board)
		a := r.getArea()
		result += p * a
	}

	return result
}

func task2(board *area) (result int) {
	println(board)
	println()
	searchSubAreas(board)

	for _, r := range board.regions {
		p := r.getPeremiter2(board)
		a := r.getArea()
		result += p * a
		println("------")
		println(r)
		println(a, " - ", p, "=", p*a, " sum ", result)
	}
	return result
}

func main() {
	input := "input.txt"

	data := readdata(input)
	start := time.Now()
	result := task1(data)
	fmt.Printf("Task 1 - elapsed Time: %12s   - result \t = %d \n", time.Since(start), result)

	start = time.Now()
	result = task2(data)
	fmt.Printf("Task 2 - elapsed Time: %12s   - result \t = %d \n", time.Since(start), result)

}
