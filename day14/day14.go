package day14

import (
	"bytes"
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
)

const day = 14

var input []byte

var example = []byte(`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
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
	p := platform(bytes.Split(bytes.TrimSuffix(input, []byte("\n")), []byte("\n")))
	p.tiltNorth()
	result := p.weight()

	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	p := platform(bytes.Split(bytes.TrimSuffix(input, []byte("\n")), []byte("\n")))

	cache := map[string]int{}
	cycleFound := false
	cycle := 0
	for cycle < 1_000_000_000 {
		for i := 0; i < 4; i++ {
			p.tiltNorth()
			p.rotate90Clockwise()
		}
		cycle++
		if cycleFound {
			continue
		}
		if prev, found := cache[string(bytes.Join(p, []byte{}))]; found && cycle < 1000 {
			cycleLength := cycle - prev
			remainingCycles := 1_000_000_000 - cycle
			skip := remainingCycles - (remainingCycles % cycleLength)
			cycle += skip
			cycleFound = true
			continue
		}

		cache[string(bytes.Join(p, []byte{}))] = cycle
	}
	result := p.weight()

	fmt.Printf("Part 2: %v\n", result)
}

func (p platform) tiltNorth() {
	for y, row := range p {
		for x, v := range row {
			if v != 'O' {
				continue
			}
			stoneY := y
			for {
				if stoneY > 0 && p[stoneY-1][x] == '.' {
					p[stoneY-1][x], p[stoneY][x] = p[stoneY][x], p[stoneY-1][x]
					stoneY--
					continue
				}
				break
			}
		}
	}
}

type platform [][]byte

func (p platform) weight() int {
	total := 0
	maxY := len(p) - 1
	for y, row := range p {
		for _, v := range row {
			if v != 'O' {
				continue
			}
			total += maxY - y + 1
		}
	}
	return total
}

func (p platform) rotate90Clockwise() {
	size := len(p)
	for layer := 0; layer < size/2; layer++ {
		first := layer
		last := size - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := p[first][i]
			p[first][i] = p[last-offset][first]
			p[last-offset][first] = p[last][last-offset]
			p[last][last-offset] = p[i][last]
			p[i][last] = top
		}
	}
}
