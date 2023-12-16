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
	start := t.start()
	var result int
	for _, p := range []point{start.north(), start.east(), start.south(), start.west()} {
		if t.connectsTo(start, p) {
			result = (1 + t.distance(start, p, start)) / 2
			break
		}
	}

	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	result := "TODO"
	fmt.Printf("Part 2: %v\n", result)
}

func (t tiles) distance(prev, cur, destination point) int {
	for _, next := range t.connectionsFrom(cur) {
		switch next {
		case prev:
			continue
		case destination:
			return 1
		}
		return 1 + t.distance(cur, next, destination)
	}
	panic("eof distance")
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
