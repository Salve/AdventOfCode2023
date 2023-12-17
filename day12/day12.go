package day12

import (
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"slices"
	"strconv"
	"strings"
)

const day = 12

var input []byte

var example = []byte(`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`)

var example2 = []byte(`?###???????? 3,2,1
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
	sum := 0
	for _, line := range inputs.Lines(input) {
		sum += variationsForLine(line)
	}
	result := sum
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	result := "TODO"
	fmt.Printf("Part 2: %v\n", result)
}

func variationsForLine(line string) int {
	springs, rest, _ := strings.Cut(line, " ")
	groups := nums(rest)
	return variations(springs, groups, 0)
}

func variations(springs string, groups []int, pos int) int {
	if pos == len(springs) {
		if valid(springs, groups) {
			return 1
		}
		return 0
	}
	if springs[pos] != '?' {
		return variations(springs, groups, pos+1)
	}
	vars := 0
	vars += variations(operational(springs, pos), groups, pos+1)
	vars += variations(damaged(springs, pos), groups, pos+1)
	return vars
}

func valid(springs string, groups []int) bool {
	return slices.Equal(groups, calcGroups(springs))
}

func calcGroups(springs string) (o []int) {
	inGroup := false
	for _, spring := range springs {
		switch spring {
		case '.':
			inGroup = false
		case '#':
			if !inGroup {
				o = append(o, 0)
				inGroup = true
			}
			o[len(o)-1]++
		}
	}
	return o
}

func operational(springs string, pos int) string {
	b := []byte(springs)
	b[pos] = '.'
	return string(b)
}

func damaged(springs string, pos int) string {
	b := []byte(springs)
	b[pos] = '#'
	return string(b)
}

func nums(line string) (o []int) {
	for _, f := range strings.Split(line, ",") {
		v, _ := strconv.Atoi(f)
		o = append(o, v)
	}
	return o
}
