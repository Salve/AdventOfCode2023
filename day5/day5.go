package day5

import (
	"cmp"
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"slices"
	"strconv"
	"strings"
)

const day = 5

var input []byte
var a almanac

func init() {
	registry.Register(day, Run)
}

func Run() {
	input = inputs.Input(day)
	a = almanacFromInput(strings.TrimSuffix(string(input), "\n"))
	part1()
	part2()
}

func part1() {
	var locs []interval
	for _, seed := range nums(strings.Split(inputs.Lines(input)[0], ": ")[1]) {
		locs = append(locs, a.lookup(interval{seed, seed})...)
	}

	slices.SortFunc(locs, func(a, b interval) int {
		return cmp.Compare(a.start, b.start)
	})
	result := locs[0].start

	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	var locs []interval
	for _, i := range seedIntervalsFromInput(input) {
		locs = append(locs, a.lookup(i)...)
	}
	slices.SortFunc(locs, func(a, b interval) int {
		return cmp.Compare(a.start, b.start)
	})
	result := locs[0].start

	fmt.Printf("Part 2: %v\n", result)
}

func seedIntervalsFromInput(input []byte) (o []interval) {
	seednums := nums(strings.Split(inputs.Lines(input)[0], ": ")[1])
	for i := 0; i < len(seednums); i += 2 {
		start, rnge := seednums[i], seednums[i+1]
		o = append(o, interval{
			start: start,
			end:   start + rnge - 1,
		})
	}
	return o
}

func almanacFromInput(input string) (a almanac) {
	for _, am := range strings.Split(input, "\n\n")[1:] {
		a = append(a, almanacMapFromInput(am))
	}
	return a
}

func almanacMapFromInput(input string) (am almanacMap) {
	for _, r := range strings.Split(input, "\n")[1:] {
		am = append(am, ruleFromInput(r))
	}
	return am
}

func ruleFromInput(input string) rule {
	vals := nums(input)
	dst, src, rnge := vals[0], vals[1], vals[2]
	return rule{
		interval: interval{
			start: src,
			end:   src + rnge - 1,
		},
		offset: dst - src,
	}
}

func nums(line string) (o []int) {
	for _, f := range strings.Fields(line) {
		v, _ := strconv.Atoi(f)
		o = append(o, v)
	}
	return o
}

type almanac []almanacMap

type almanacMap []rule

type rule struct {
	interval
	offset int
}

type interval struct {
	start, end int
}

func (a almanac) lookup(i interval) []interval {
	current := []interval{i}
	for _, am := range a {
		current = am.lookupMulti(current)
	}
	return current
}

func (am almanacMap) lookupMulti(intervals []interval) (o []interval) {
	for _, i := range intervals {
		o = append(o, am.lookup(i)...)
	}
	return o
}

func (am almanacMap) lookup(i interval) []interval {
	mapped := []interval{}
	unmapped := []interval{i}

	for _, r := range am {
		m, u := r.applyToMulti(unmapped)
		mapped = append(mapped, m...)
		unmapped = u
	}
	return append(mapped, unmapped...)
}

func (r rule) applyToMulti(intervals []interval) (mapped []interval, unmapped []interval) {
	for _, i := range intervals {
		m, u := r.applyTo(i)
		mapped = append(mapped, m...)
		unmapped = append(unmapped, u...)
	}
	return
}

func (r rule) applyTo(i interval) (mapped []interval, unmapped []interval) {
	if i.end < r.start || i.start > r.end {
		unmapped = []interval{i}
		return
	}
	if i.start < r.start {
		unmapped = append(unmapped, interval{i.start, r.start - 1})
	}
	common := interval{max(i.start, r.start), min(i.end, r.end)}
	mapped = append(mapped, common.add(r.offset))
	if i.end > r.end {
		unmapped = append(unmapped, interval{r.end + 1, i.end})
	}
	return
}

func (i interval) add(v int) interval {
	return interval{i.start + v, i.end + v}
}
