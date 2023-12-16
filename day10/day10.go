package day10

import (
	"bytes"
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
)

const day = 10

var input []byte

var example = []byte(`..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`)

func init() {
	registry.Register(day, Run)
}

func Run() {
	input = inputs.Input(day)
	part1()
	part2()
}

func part1() {
	t := tilesFromInput(input)
	l := t.getLoop()
	result := len(l) / 2

	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	t := tilesFromInput(input)
	l := t.getLoop()
	result := enclosedPoints(shoelace(l), len(l))
	fmt.Printf("Part 2: %v\n", result)
}

func (t tiles) getLoop() []point {
	start := t.start()
	loop := []point{start}
	for _, p := range []point{start.north(), start.east(), start.south(), start.west()} {
		if t.connectsTo(start, p) {
			t.path(start, p, start, &loop)
			break // we don't need to traverse in both directions
		}
	}
	return loop
}

func (t tiles) path(prev, cur, destination point, acc *[]point) {
	*acc = append(*acc, cur)
	for _, next := range t.connectionsFrom(cur) {
		switch next {
		case prev:
			continue
		case destination:
			return
		}
		t.path(cur, next, destination, acc)
	}
}

func (t tiles) connectionsFrom(p point) (o []point) {
	switch t[p] {
	case '|':
		return append(o, p.north(), p.south())
	case '-':
		return append(o, p.west(), p.east())
	case 'L':
		return append(o, p.north(), p.east())
	case 'J':
		return append(o, p.north(), p.west())
	case '7':
		return append(o, p.south(), p.west())
	case 'F':
		return append(o, p.south(), p.east())
	case 'S':
		panic("start")
	default:
		return o
	}
}

func (t tiles) connectsTo(a, b point) bool {
	for _, p := range t.connectionsFrom(b) {
		if p == a {
			return true
		}
	}
	return false
}

func (p point) north() point {
	return point{p.x, p.y - 1}
}
func (p point) east() point {
	return point{p.x + 1, p.y}
}
func (p point) south() point {
	return point{p.x, p.y + 1}
}
func (p point) west() point {
	return point{p.x - 1, p.y}
}

func (t tiles) start() point {
	for p, v := range t {
		if v == 'S' {
			return p
		}
	}
	panic("eof")
}

type point struct {
	x, y int
}

type tiles map[point]byte

func tilesFromInput(input []byte) tiles {
	t := make(tiles, len(input))
	for y, line := range bytes.Split(input, []byte("\n")) {
		for x, v := range line {
			t[point{x, y}] = v
		}
	}
	return t
}

// apparently there's something called a shoelace formula. news to me.
func shoelace(points []point) int {
	n := len(points)
	area := 0

	for i := 0; i < n-1; i++ {
		area += points[i].x*points[i+1].y - points[i].y*points[i+1].x
	}
	area += points[n-1].x*points[0].y - points[n-1].y*points[0].x

	if area < 0 {
		area = -area
	}

	return area
}

// apparently there's something called picks theorem. news to me.
func enclosedPoints(area int, numPoints int) int {
	return (area - numPoints + 2) / 2
}
