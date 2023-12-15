package day5

import (
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"slices"
	"strconv"
	"strings"
)

const day = 5

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
	lines := inputs.Lines(input)
	var locations []int
	for _, seed := range nums(strings.Split(lines[0], ": ")[1]) {
		locations = append(locations, locationForSeed(seed, lines[1:]))
	}
	result := slices.Min(locations)
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	lines := inputs.Lines(input)
	var locations []int
	seednums := nums(strings.Split(lines[0], ": ")[1])
	for i := 0; i < len(seednums); i += 2 {
		start, rnge := seednums[i], seednums[i+1]
		for seed := start; seed < start+rnge; seed++ {
			locations = append(locations, locationForSeed(seed, lines[1:]))
		}
	}
	result := slices.Min(locations)
	fmt.Printf("Part 2: %v\n", result)
}

func locationForSeed(seed int, almanac []string) int {
	cur, next := seed, seed
	for _, line := range almanac {
		v := nums(line)
		if len(v) < 3 { // We're at a new section, next is the new value to map
			cur = next
			continue
		}
		if newV := lookup(cur, v[0], v[1], v[2]); newV != 0 {
			next = newV
		}
	}
	return next
}

func lookup(value, dst, src, rnge int) int {
	if value < src || value > src+rnge {
		return 0
	}
	return value + dst - src
}

func nums(line string) (o []int) {
	for _, v := range strings.Fields(line) {
		nv, _ := strconv.Atoi(v)
		o = append(o, nv)
	}
	return o
}
