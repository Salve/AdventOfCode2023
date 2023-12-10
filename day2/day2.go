package day2

import (
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"regexp"
	"strconv"
	"strings"
)

const day = 2

var input []byte

func init() {
	registry.Register(day, Run)
}

func Run() {
	input = inputs.Input(day)
	part1()
	part2()
}

var reGame = regexp.MustCompile(`\d+ (?:red|green|blue)`)

func part1() {
	sum := 0
	for game, line := range inputs.Lines(input) {
		if legal(line) {
			sum += game + 1
		}
	}
	result := sum
	fmt.Printf("Part 1: %v\n", result)
}

func legal(line string) bool {
	for _, g := range reGame.FindAllString(line, -1) {
		gs := strings.Split(g, " ")
		v, _ := strconv.Atoi(gs[0])
		switch gs[1] {
		case "red":
			if v > 12 {
				return false
			}
		case "green":
			if v > 13 {
				return false
			}
		case "blue":
			if v > 14 {
				return false
			}
		}
	}
	return true
}

func part2() {
	sum := 0
	for _, line := range inputs.Lines(input) {
		sum += power(line)
	}
	result := sum
	fmt.Printf("Part 2: %v\n", result)
}

func power(line string) int {
	colors := map[string]int{"red": 0, "green": 0, "blue": 0}
	for _, g := range reGame.FindAllString(line, -1) {
		split := strings.Split(g, " ")
		v, _ := strconv.Atoi(split[0])
		color := split[1]
		if v > colors[color] {
			colors[color] = v
		}
	}
	return colors["red"] * colors["green"] * colors["blue"]
}
