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

	var number []byte
	var adjacent bool
	var y int
	for i, v := range input {
		if v == byte('\n') {
			y++
		}
		if v >= '0' && v <= '9' {
			number = append(number, v)
			adjacent = adjacent || s.adjacentToSymbol(point{i % 141, y})
			continue
		}
		if len(number) > 0 {
			if adjacent {
				sum += numberToInt(number)
			}
			number = number[:0]
			adjacent = false
		}
	}

	result := sum
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	sum := 0
	s := schematicFromInput(input)

	gears := make(map[point][]int)
	var number []byte
	var gear point
	var adjacent bool
	var y int
	for i, v := range input {
		if v == byte('\n') {
			y++
		}
		if v >= '0' && v <= '9' {
			number = append(number, v)
			if loc, ok := s.adjacentToGear(point{i % 141, y}); ok {
				gear = loc
				adjacent = true
			}
			continue
		}
		if len(number) > 0 {
			if adjacent {
				gears[gear] = append(gears[gear], numberToInt(number))
			}
			number = number[:0]
			adjacent = false
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

func (s schematic) adjacentToSymbol(p point) bool {
	x, y := p.x, p.y
	return s.isSymbol(point{x - 1, y - 1}) || s.isSymbol(point{x - 1, y}) || s.isSymbol(point{x - 1, y + 1}) || s.isSymbol(point{x, y - 1}) || s.isSymbol(point{x, y + 1}) || s.isSymbol(point{x + 1, y - 1}) || s.isSymbol(point{x + 1, y}) || s.isSymbol(point{x + 1, y + 1})
}

func (s schematic) isSymbol(p point) bool {
	switch s[p] {
	case 0, '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
		return false
	default:
		return true
	}
}

func (s schematic) adjacentToGear(p point) (point, bool) {
	x, y := p.x, p.y
	adjacencies := []point{{x - 1, y - 1}, {x - 1, y}, {x - 1, y + 1}, {x, y - 1}, {x, y + 1}, {x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1}}
	for _, a := range adjacencies {
		if s[a] == '*' {
			return a, true
		}
	}
	return p, false
}

func numberToInt(in []byte) int {
	out, _ := strconv.Atoi(string(in))
	return out
}
