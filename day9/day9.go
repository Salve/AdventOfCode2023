package day9

import (
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"slices"
	"strconv"
	"strings"
)

const day = 9

var input []byte

var example = []byte(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
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
	result := 0
	for _, line := range inputs.Lines(input) {
		result += predict(nums(line, false))
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	result := 0
	for _, line := range inputs.Lines(input) {
		result += predict(nums(line, true))
	}
	fmt.Printf("Part 2: %v\n", result)
}

func predict(seq []int) int {
	if allZero(seq) {
		return 0
	}
	return seq[len(seq)-1] + predict(deltas(seq))
}

func deltas(seq []int) (o []int) {
	for i := 0; i < len(seq)-1; i++ {
		o = append(o, seq[i+1]-seq[i])
	}
	return o
}

func nums(line string, reverse bool) (o []int) {
	for _, f := range strings.Fields(line) {
		v, _ := strconv.Atoi(f)
		o = append(o, v)
	}
	if reverse {
		slices.Reverse(o)
	}
	return o
}

func allZero(seq []int) bool {
	for _, v := range seq {
		if v != 0 {
			return false
		}
	}
	return true
}
