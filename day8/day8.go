package day8

import (
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"strings"
)

const day = 8

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
	m := readMap(input)
	result := m.stepsToZ(0, "AAA", inputs.Lines(input)[0])
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	m := readMap(input)
	directions := inputs.Lines(input)[0]
	var cycles []int
	for pos := range m {
		if pos[2] == 'A' {
			cycles = append(cycles, m.stepsToZ(0, pos, directions))
		}
	}
	result := cycles[0]
	for _, v := range cycles[1:] {
		result = lcm(result, v)
	}
	fmt.Printf("Part 2: %v\n", result)
}

func (m camelMap) stepsToZ(stepsTaken int, position, directions string) int {
	if position[2] == 'Z' {
		return stepsTaken
	}
	dir := 0
	if directions[stepsTaken%len(directions)] == 'R' {
		dir = 1
	}
	return m.stepsToZ(stepsTaken+1, m[position][dir], directions)
}

type camelMap map[string][2]string

func readMap(input []byte) camelMap {
	o := make(camelMap)
	for _, line := range strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")[2:] {
		n1, n2, n3 := line[0:3], line[7:10], line[12:15]
		o[n1] = [2]string{n2, n3}
	}
	return o
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
