package day21

import (
	"fmt"
	"image"
	"maps"
	"strings"

	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
)

const day = 21

var input []byte

var example = []byte(`...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
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
	g := gridFromInput(input)
	for range 64 {
		g = step(g)
	}
	result := count(g)
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	result := "TODO"
	fmt.Printf("Part 2: %v\n", result)
}

type grid map[image.Point]rune

func gridFromInput(input []byte) grid {
	g := make(grid, len(input))
	for y, line := range strings.Split(string(input), "\n") {
		for x, v := range line {
			p := image.Point{x, y}
			g[p] = v
		}
	}
	return g
}

func step(g grid) grid {
	points, wiped := current(g)
	for _, p := range points {
		for _, d := range dirs {
			newpos := p.Add(d)
			if v, ok := wiped[newpos]; ok && v == '.' {
				wiped[newpos] = 'S'
			}
		}
	}
	return wiped
}

func count(g grid) int {
	c := 0
	for _, v := range g {
		if v == 'S' {
			c++
		}
	}
	return c
}

func current(g grid) ([]image.Point, grid) {
	wiped := maps.Clone(g)
	points := []image.Point{}
	for k, v := range g {
		if v != 'S' {
			continue
		}
		points = append(points, k)
		wiped[k] = '.'
	}
	return points, wiped
}

var dirs = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
