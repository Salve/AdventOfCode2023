package day15

import (
	"bytes"
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
)

const day = 15

var input []byte

var example = []byte(`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`)

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
	for _, step := range bytes.Split(bytes.TrimSuffix(input, []byte("\n")), []byte(",")) {
		result += HASH(step)
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	result := "TODO"
	fmt.Printf("Part 2: %v\n", result)
}

func HASH(in []byte) int {
	cv := 0
	for _, c := range in {
		cv += int(c)
		cv *= 17
		cv %= 256
	}
	return cv
}
