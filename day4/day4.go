package day4

import (
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"slices"
	"strings"
)

const day = 4

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
	var total int
	for _, line := range inputs.Lines(input) {
		if matches := countMatches(line); matches > 0 {
			total += 1 << (matches - 1)
		}
	}

	result := total
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	totals := [199]int{}
	lines := inputs.Lines(input)
	for i, line := range lines {
		totals[i]++ // original card
		matches := countMatches(line)
		for j := 1; j <= matches; j++ {
			totals[i+j] += totals[i]
		}
	}
	var total int
	for _, v := range totals {
		total += v
	}

	result := total
	fmt.Printf("Part 2: %v\n", result)
}

func countMatches(line string) int {
	var total int
	fields := strings.Fields(line)
	for _, v := range fields[13:] {
		if slices.Contains(fields[2:12], v) {
			total++
		}
	}
	return total
}
