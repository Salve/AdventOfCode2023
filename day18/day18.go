package day18

import (
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"image"
	"strconv"
	"strings"
)

const day = 18

var input []byte

var example = []byte(`R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
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

	site := digsiteFromInput(input, false)
	area := shoelace(site)
	result := totalPoints(area, len(site))
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	site := digsiteFromInput(input, true)
	area := shoelace(site)
	result := totalPoints(area, len(site))
	fmt.Printf("Part 2: %v\n", result)
}

func digsiteFromInput(input []byte, p2 bool) []image.Point {
	cur := image.Point{0, 0}
	o := []image.Point{}
	for _, line := range inputs.Lines(input) {
		s := strings.Split(string(line), " ")
		dir, dist := dirs[s[0][0]], dropErr(s[1], strconv.Atoi)
		if p2 {
			dir, dist = decode(s[2])
		}
		for i := 0; i < dist; i++ {
			cur = cur.Add(dir)
			o = append(o, cur)
		}
	}
	return o
}

func decode(in string) (image.Point, int) {
	d, _ := strconv.ParseInt(in[2:7], 16, 64)
	return dirs[in[7]], int(d)
}

func shoelace(points []image.Point) int {
	n := len(points)
	area := 0

	for i := 0; i < n-1; i++ {
		area += points[i].X*points[i+1].Y - points[i].Y*points[i+1].X
	}
	area += points[n-1].X*points[0].Y - points[n-1].Y*points[0].X

	if area < 0 {
		area = -area
	}

	return area
}

func totalPoints(area int, numPoints int) int {
	return ((area - numPoints + 2) / 2) + numPoints
}

var dirs = map[byte]image.Point{
	'U': {0, -1},
	'R': {1, 0},
	'D': {0, 1},
	'L': {-1, 0},
	'3': {0, -1},
	'0': {1, 0},
	'1': {0, 1},
	'2': {-1, 0},
}

func num(a string) int {
	v, _ := strconv.Atoi(a)
	return v
}

func dropErr[P, R any](p P, f func(P) (R, error)) R {
	v, _ := f(p)
	return v
}
