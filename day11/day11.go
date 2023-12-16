package day11

import (
	"bytes"
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
)

const day = 11

var input []byte

var example = []byte(`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
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
	g := galaxiesFromInput(input, 2)
	sum := 0
	for i := 0; i < len(g)-1; i++ {
		for j := i + 1; j < len(g); j++ {
			sum += distance(g[i], g[j])
		}
	}
	fmt.Printf("Part 1: %v\n", sum)
}

func part2() {
	g := galaxiesFromInput(input, 1_000_000)
	sum := 0
	for i := 0; i < len(g)-1; i++ {
		for j := i + 1; j < len(g); j++ {
			sum += distance(g[i], g[j])
		}
	}
	fmt.Printf("Part 2: %v\n", sum)
}

func distance(a, b point) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	return abs(a.x-b.x) + abs(a.y-b.y)
}

type point struct {
	x, y int
}

func galaxiesFromInput(input []byte, expansion int) (o []point) {
	lines := bytes.Split(input[:len(input)-1], []byte("\n"))

	xLen, yLen := len(lines[0]), len(lines)
	xMap, yMap := make([]int, xLen), make([]int, yLen)
	newX := 0
	for i := 0; i < xLen; i++ {
		if emptyColumn(lines, i) {
			newX += expansion - 1
		}
		xMap[i] = newX
		newX++
	}
	newY := 0
	for i := 0; i < yLen; i++ {
		if emptyRow(lines, i) {
			newY += expansion - 1
		}
		yMap[i] = newY
		newY++
	}

	for y, line := range lines {
		for x, v := range line {
			if v == '.' {
				continue
			}
			o = append(o, point{xMap[x], yMap[y]})
		}
	}
	return o
}

func emptyRow(g [][]byte, i int) bool {
	return bytes.IndexByte(g[i], '#') == -1
}

func emptyColumn(g [][]byte, j int) bool {
	for i := 0; i < len(g); i++ {
		if g[i][j] == '#' {
			return false
		}
	}
	return true
}
