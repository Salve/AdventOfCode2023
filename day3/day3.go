package day3

import (
	"bytes"
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"strconv"
)

const day = 3

var input []byte

func init() {
	registry.Register(day, Run)
}

func Run() {
	input = inputs.Input(day)
	part1()
	part2()
}

func part1() {
	sum := 0
	s := schematicFromInput(input)
	nums := numbersFromInput(input)

	for _, num := range nums {
		if num.adjacentToSymbol(s) {
			sum += num.value()
		}
	}

	result := sum
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	sum := 0
	s := schematicFromInput(input)
	nums := numbersFromInput(input)
	gears := make(map[point][]int)

	for _, num := range nums {
		if gear, adj := num.adjacentToGear(s); adj {
			gears[gear] = append(gears[gear], num.value())
		}
	}
	for _, v := range gears {
		if len(v) != 2 {
			continue
		}
		sum += v[0] * v[1]
	}

	result := sum
	fmt.Printf("Part 2: %v\n", result)
}

type point struct {
	x, y int
}
type number struct {
	digits []byte
	points []point
}
type schematic map[point]byte

func schematicFromInput(input []byte) schematic {
	s := make(schematic, len(input))
	for y, line := range bytes.Split(input, []byte("\n")) {
		for x, v := range line {
			s[point{x, y}] = v
		}
	}
	return s
}

func numbersFromInput(input []byte) (o []number) {
	y := 0
	cur := number{}
	for i, v := range input {
		if v == byte('\n') {
			y++
		}
		if v >= '0' && v <= '9' {
			cur.digits = append(cur.digits, v)
			cur.points = append(cur.points, point{i % 141, y})
			continue
		}
		if len(cur.digits) > 0 {
			o = append(o, cur)
			cur = number{}
		}
	}
	return o
}

func (n number) adjacentToSymbol(s schematic) bool {
	for _, p := range n.points {
		if s.adjacentToSymbol(p) {
			return true
		}
	}
	return false
}

func (n number) adjacentToGear(s schematic) (point, bool) {
	for _, p := range n.points {
		if gear, adj := s.adjacentToGear(p); adj {
			return gear, true
		}
	}
	return point{-1, -1}, false
}

func (s schematic) adjacentToSymbol(p point) bool {
	for _, a := range adjacentPoints(p) {
		if v := s[a]; v == 0 || v == '.' || (v >= '0' && v <= '9') {
			continue
		}
		return true
	}
	return false
}

func (s schematic) adjacentToGear(p point) (point, bool) {
	for _, a := range adjacentPoints(p) {
		if s[a] == '*' {
			return a, true
		}
	}
	return p, false
}

func adjacentPoints(p point) []point {
	x, y := p.x, p.y
	return []point{{x - 1, y - 1}, {x - 1, y}, {x - 1, y + 1}, {x, y - 1}, {x, y + 1}, {x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1}}
}

func (n number) value() int {
	out, _ := strconv.Atoi(string(n.digits))
	return out
}
