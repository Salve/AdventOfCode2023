package day13

import (
	"bytes"
	"fmt"
	"github.com/Salve/AdventOfCode2023/registry"
)

const day = 13

var example = []byte(`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`)
var patterns []pattern

func init() {
	registry.Register(day, Run)
}

func Run() {
	//patterns = patternsFromInput(inputs.Input(day))
	patterns = patternsFromInput(example)
	part1()
	part2()
}

func part1() {
	result := 0
	for _, p := range patterns {
		result += p.summarize()
	}

	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	result := "TODO"
	fmt.Printf("Part 2: %v\n", result)
}

func (p pattern) summarize(smudges int) int {
	total := 0
	for x := 0; x < len(p[0])-1; x++ {
		if p.vertical(x, x+1) {
			total += x + 1
		}
	}
	for y := 0; y < len(p)-1; y++ {
		if p.horizontal(y, y+1) {
			total += 100 * (y + 1)
		}
	}
	return total
}

func (p pattern) vertical(x1, x2 int, smudges int) bool {
	if x1 < 0 || x2 >= len(p[0]) {
		return true
	}
	for y := range p {
		if p[y][x1] != p[y][x2] {
			return false
		}
	}
	return p.vertical(x1-1, x2+1)
}

func (p pattern) horizontal(y1, y2 int, smudges int) bool {
	if y1 < 0 || y2 >= len(p) {
		return true
	}
	for x := range p[0] {
		if p[y1][x] != p[y2][x] {
			return false
		}
	}
	return p.horizontal(y1-1, y2+1)
}

type pattern [][]byte

func patternsFromInput(input []byte) (o []pattern) {
	for _, p := range bytes.Split(bytes.TrimSuffix(input, []byte("\n")), []byte("\n\n")) {
		o = append(o, bytes.Split(p, []byte("\n")))
	}
	return o
}
