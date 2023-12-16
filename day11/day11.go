package day11

import (
	"bytes"
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
)

const day = 11

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
	g := galaxiesFromInput(input)
	sum := 0
	for i := 0; i < len(g)-1; i++ {
		for j := i + 1; j < len(g); j++ {
			sum += distance(g[i], g[j])
		}
	}
	fmt.Printf("Part 1: %v\n", sum)
}

func part2() {
	result := "TODO"
	fmt.Printf("Part 2: %v\n", result)
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

func galaxiesFromInput(input []byte) (o []point) {
	lines := bytes.Split(input[:len(input)-1], []byte("\n"))
	lines = doubleEmptyRows(lines)
	lines = rotate(lines)
	lines = doubleEmptyRows(lines)

	for y, line := range lines {
		for x, v := range line {
			if v == '.' {
				continue
			}
			o = append(o, point{x, y})
		}
	}
	return o
}

func rotate(in [][]byte) [][]byte {
	m, n := len(in), len(in[0])
	out := make([][]byte, n)
	for i := range out {
		out[i] = make([]byte, m)
		for j := range out[i] {
			out[i][j] = in[m-j-1][i]
		}
	}
	return out
}

func doubleEmptyRows(in [][]byte) (out [][]byte) {
	for _, line := range in {
		if bytes.IndexByte(line, '#') == -1 {
			out = append(out, line)
		}
		out = append(out, line)
	}
	return out
}
