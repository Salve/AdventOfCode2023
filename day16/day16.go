package day16

import (
	"bytes"
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"slices"
)

const day = 16

var input []byte

var example = []byte(`.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
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
	c := contraptionFromInput(input)
	c.beam(point{0, 0}, right)
	result := c.energized()
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	c := contraptionFromInput(input)
	result := 0
	for i := 0; i < 110; i++ {
		for _, dir := range []point{up, down, left, right} {
			c.reset()
			switch dir {
			case up:
				c.beam(point{0 + i, 109}, dir)
			case down:
				c.beam(point{0 + i, 0}, dir)
			case left:
				c.beam(point{109, 0 + i}, dir)
			case right:
				c.beam(point{0, 0 + i}, dir)
			}
			if e := c.energized(); e > result {
				result = e
			}
		}
	}

	fmt.Printf("Part 2: %v\n", result)
}

type point struct {
	x, y int
}

type tile struct {
	t       byte
	visited []point
}

type contraption map[point]tile

func contraptionFromInput(input []byte) contraption {
	c := make(contraption, len(input))
	for y, line := range bytes.Split(input, []byte("\n")) {
		for x, v := range line {
			c[point{x, y}] = tile{t: v}
		}
	}
	return c
}

func (c contraption) beam(loc point, dir point) {
	if c.visited(loc, dir) {
		return
	}
	switch c[loc].t {
	case '.':
		c.beam(loc.add(dir), dir)
	case '|':
		switch dir {
		case up, down:
			c.beam(loc.add(dir), dir)
		case left, right:
			c.beam(loc.add(up), up)
			c.beam(loc.add(down), down)
		}
	case '-':
		switch dir {
		case up, down:
			c.beam(loc.add(left), left)
			c.beam(loc.add(right), right)
		case left, right:
			c.beam(loc.add(dir), dir)
		}
	case '/':
		switch dir {
		case up:
			c.beam(loc.add(right), right)
		case down:
			c.beam(loc.add(left), left)
		case left:
			c.beam(loc.add(down), down)
		case right:
			c.beam(loc.add(up), up)
		}
	case '\\':
		switch dir {
		case up:
			c.beam(loc.add(left), left)
		case down:
			c.beam(loc.add(right), right)
		case left:
			c.beam(loc.add(up), up)
		case right:
			c.beam(loc.add(down), down)
		}

	}
}

func (c contraption) visited(loc, dir point) bool {
	if slices.Contains(c[loc].visited, dir) {
		return true
	}
	if t, found := c[loc]; found {
		t.visited = append(t.visited, dir)
		c[loc] = t
	}
	return false
}

func (c contraption) energized() int {
	total := 0
	for _, t := range c {
		if len(t.visited) > 0 {
			total++
		}
	}
	return total
}

func (a point) add(b point) point {
	return point{a.x + b.x, a.y + b.y}
}

var (
	up    = point{0, -1}
	right = point{1, 0}
	down  = point{0, 1}
	left  = point{-1, 0}
)

func (c contraption) reset() {
	for k, p := range c {
		p.visited = p.visited[:0]
		c[k] = p
	}
}
