package day6

import (
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"strconv"
	"strings"
)

const day = 6

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
	times := nums(strings.Split(lines[0], ":")[1])
	distances := nums(strings.Split(lines[1], ":")[1])

	result := 1
	for race := range times {
		ways := 0
		for hold := 0; hold < times[race]; hold++ {
			distance := hold * (times[race] - hold)
			if distance > distances[race] {
				ways++
			}
		}
		result *= ways
	}

	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	lines := inputs.Lines(input)
	time := num(lines[0])
	record := num(lines[1])

	ways := 0
	for hold := 0; hold < time; hold++ {
		distance := hold * (time - hold)
		if distance > record {
			ways++
		}
	}

	result := ways
	fmt.Printf("Part 2: %v\n", result)
}

func nums(line string) (o []int) {
	for _, f := range strings.Fields(line) {
		v, _ := strconv.Atoi(f)
		o = append(o, v)
	}
	return o
}

func num(line string) int {
	total := 0
	for _, r := range line {
		if r < '0' || r > '9' {
			continue
		}
		total *= 10
		total += int(r - '0')
	}
	return total
}
