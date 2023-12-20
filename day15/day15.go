package day15

import (
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"slices"
	"strings"
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
	for _, step := range strings.Split(strings.TrimSuffix(string(input), "\n"), ",") {
		result += HASH(string(step))
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	hm := HASHMAP{}
	for _, op := range strings.Split(strings.TrimSuffix(string(input), "\n"), ",") {
		l := len(op)
		if strings.Contains(op, "-") {
			hm.rm(op[0 : l-1])
			continue
		}
		hm.add(lens{op[0 : l-2], int(op[l-1] - '0')})
	}
	result := hm.focusPower()
	fmt.Printf("Part 2: %v\n", result)
}

func HASH(in string) int {
	cv := 0
	for _, c := range in {
		cv += int(c)
		cv *= 17
		cv %= 256
	}
	return cv
}

type lens struct {
	label       string
	focalLength int
}

type HASHMAP [256][]lens

func (hm *HASHMAP) focusPower() int {
	total := 0
	for box, slots := range hm {
		for slot, lens := range slots {
			total += (box + 1) * (slot + 1) * lens.focalLength
		}
	}
	return total
}

func (hm *HASHMAP) rm(label string) {
	box := HASH(label)
	for slot, lens := range hm[box] {
		if lens.label == label {
			hm[box] = slices.Delete(hm[box], slot, slot+1)
			return
		}
	}
}

func (hm *HASHMAP) add(nl lens) {
	box := HASH(nl.label)
	for i, ol := range hm[box] {
		if ol.label == nl.label {
			hm[box][i] = nl
			return
		}
	}
	hm[box] = append(hm[box], nl)
}
