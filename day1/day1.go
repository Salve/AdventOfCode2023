package day1

import (
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"strings"
)

const day = 1

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
	sum := 0
	for _, line := range inputs.Lines(input) {
		sum += calibrationValue(line, digits)
	}
	result := sum
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	sum := 0
	for _, line := range inputs.Lines(input) {
		sum += calibrationValue(line, alphaDigits)
	}
	result := sum
	fmt.Printf("Part 2: %v\n", result)
}

func calibrationValue(line string, digits map[string]int) int {
	combined := first(line, digits)*10 + last(line, digits)
	return combined
}

func first(line string, digits map[string]int) int {
	for pos := 0; pos < len(line); pos++ {
		for k, v := range digits {
			if strings.HasPrefix(line[pos:], k) {
				return v
			}
		}
	}
	panic("eof")
}

func last(line string, digits map[string]int) int {
	for pos := len(line); pos >= 0; pos-- {
		for k, v := range digits {
			if strings.HasSuffix(line[:pos], k) {
				return v
			}
		}
	}
	panic("eof")
}

var digits = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

var alphaDigits = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}
