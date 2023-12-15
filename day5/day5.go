package day5

import (
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
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
	result := "TODO"
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	result := "TODO"
	fmt.Printf("Part 2: %v\n", result)
}

func lookup(value, dst, src, rnge int) int {
	if value < src || value > src+rnge {
		return value
	}

}
