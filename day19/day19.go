package day19

import (
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"maps"
	"strconv"
	"strings"
)

const day = 19

var input []byte

var example = []byte(`px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}
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
	flows, parts := flowsAndPartsFromInput(input)

	result := 0
	for _, part := range parts {
		if accepted(part, "in", flows) {
			result += partValue(part)
		}
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	flows, _ := flowsAndPartsFromInput(input)

	ranges := attrRanges{
		'x': {1, 4000},
		'm': {1, 4000},
		'a': {1, 4000},
		's': {1, 4000},
	}

	result := 0
	for _, r := range acceptedRanges(0, "in", flows, ranges) {
		xRange := r['x'][1] - r['x'][0] + 1
		mRange := r['m'][1] - r['m'][0] + 1
		aRange := r['a'][1] - r['a'][0] + 1
		sRange := r['s'][1] - r['s'][0] + 1

		combinations := xRange * mRange * aRange * sRange
		result += combinations
	}
	fmt.Printf("Part 2: %v\n", result)
}

func acceptedRanges(i int, flow string, flows map[string][]string, ranges attrRanges) []attrRanges {
	if flow == "A" {
		return []attrRanges{ranges}
	}
	if flow == "R" {
		return nil
	}
	step := flows[flow][i]
	condition, jump, found := strings.Cut(step, ":")
	if !found {
		return acceptedRanges(0, step, flows, ranges)
	}
	rangesTrue, rangesFalse := constrain(ranges, condition)
	return append(acceptedRanges(0, jump, flows, rangesTrue), acceptedRanges(i+1, flow, flows, rangesFalse)...)
}

func accepted(part map[byte]int, flow string, flows map[string][]string) bool {
	if flow == "A" {
		return true
	}
	if flow == "R" {
		return false
	}
	for _, step := range flows[flow] {
		condition, jump, found := strings.Cut(step, ":")
		if !found {
			return accepted(part, step, flows)
		}
		if conditionMet(part, condition) {
			return accepted(part, jump, flows)
		}
	}
	return true
}

func constrain(ranges attrRanges, condition string) (cTrue, cFalse attrRanges) {
	cTrue, cFalse = maps.Clone(ranges), maps.Clone(ranges)
	attr, val := ranges[condition[0]], num(condition[2:])
	attrTrue, attrFalse := attr, attr
	switch condition[1] {
	case '<':
		if attr[1] >= val {
			attrTrue[1] = val - 1
		}
		if attr[0] < val {
			attrFalse[0] = val
		}
	case '>':
		if attr[0] <= val {
			attrTrue[0] = val + 1
		}
		if attr[1] > val {
			attrFalse[1] = val
		}
	}
	cTrue[condition[0]] = attrTrue
	cFalse[condition[0]] = attrFalse
	return cTrue, cFalse
}

func conditionMet(part map[byte]int, condition string) bool {
	switch condition[1] {
	case '<':
		return part[condition[0]] < num(condition[2:])
	case '>':
		return part[condition[0]] > num(condition[2:])
	}
	panic("condition")
}

func num(i string) int {
	v, _ := strconv.Atoi(i)
	return v
}

func flowsAndPartsFromInput(input []byte) (flows map[string][]string, parts []map[byte]int) {
	inFlows, inParts, _ := strings.Cut(string(input), "\n\n")
	flows = map[string][]string{}
	for _, l := range strings.Split(inFlows, "\n") {
		flowName, flowText, _ := strings.Cut(l, "{")
		flows[flowName] = strings.Split(flowText[:len(flowText)-1], ",")
	}

	for _, l := range strings.Split(strings.TrimSuffix(inParts, "\n"), "\n") {
		part := map[byte]int{}
		for _, attr := range strings.Split(l[1:len(l)-1], ",") {
			part[attr[0]] = num(attr[2:])
		}
		parts = append(parts, part)
	}
	return flows, parts
}

func partValue(part map[byte]int) (sum int) {
	for _, v := range part {
		sum += v
	}
	return sum
}

type attrRanges map[byte][2]int
