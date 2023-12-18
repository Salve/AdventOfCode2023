package day12

import (
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"strconv"
	"strings"
)

const day = 12

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
	result := 0
	for _, line := range inputs.Lines(input) {
		result += variationsForLine(line, 1)
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	result := 0
	for _, line := range inputs.Lines(input) {
		result += variationsForLine(line, 5)
	}
	fmt.Printf("Part 2: %v\n", result)
}

func variationsForLine(line string, expand int) int {
	s, g, _ := strings.Cut(line, " ")
	var springs, groups string
	for i := 0; i < expand; i++ {
		springs, groups = springs+s+"?", groups+g+","
	}
	springs, groups = strings.TrimSuffix(springs, "?"), strings.TrimSuffix(groups, ",")
	return variations(springs+".", nums(groups)) // avoid a special case involving the last spring
}

var memo = map[string]int{}

func memoVariations(springs string, groups []int) int {
	key := fmt.Sprintf("%s%v", springs, groups)
	if result, found := memo[key]; found {
		return result
	}
	result := variations(springs, groups)
	memo[key] = result
	return result
}

func variations(springs string, groups []int) int {
	if len(groups) == 0 {
		if strings.Contains(springs, "#") {
			return 0 // all groups are consumed but broken springs remain - not valid
		}
		return 1 // this is a valid variation, all groups have been accounted for, no broken springs remain
	}
	// try placing the next group of damaged springs at each possible starting position that leaves room for all the
	// known remaining groups of broken springs.
	result := 0
	for i := 0; i <= len(springs)-sum(groups)-len(groups)+1; i++ {
		next := i + groups[0]
		if strings.Contains(springs[:i], "#") {
			break // if we've passed a known damaged spring, we know the next group has started, no need to keep looking
		}

		// skip this position if it overlaps with a known good spring, or runs straight into a known damaged spring
		if !strings.Contains(springs[i:next], ".") && springs[next:next+1] != "#" {
			// it's not impossible to place a group here, try to recurse
			result += memoVariations(springs[next+1:], groups[1:])
		}
	}
	return result
}

func nums(line string) (o []int) {
	for _, f := range strings.Split(line, ",") {
		v, _ := strconv.Atoi(f)
		o = append(o, v)
	}
	return o
}

func sum(a []int) (o int) {
	for _, v := range a {
		o += v
	}
	return o
}
