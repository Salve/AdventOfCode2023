package day20

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
)

const day = 20

var input []byte

var example1 = []byte(`broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`)

var example2 = []byte(`broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
`)

func init() {
	registry.Register(day, Run)
}

func Run() {
	input = inputs.Input(day)
	part1()
	part2()
}

var re = regexp.MustCompile(`^(?P<Prefix>[%&]?)(?P<Name>[a-zA-Z]+) -> (?P<Destinations>.+)$`)

func initMods(data []byte) (map[string]*module, *pulseTracker) {
	mods := map[string]*module{}
	pt := pulseTracker{}
	for _, line := range inputs.Lines(data) {
		match := re.FindStringSubmatch(line)
		mods[match[2]] = &module{
			prefix:  match[1],
			name:    match[2],
			outputs: strings.Split(match[3], ", "),
			inputs:  make(map[string]bool),
			pt:      &pt,
		}
	}
	for _, mod := range mods {
		for _, output := range mod.outputs {
			if _, ok := mods[output]; !ok {
				mods[output] = &module{name: output, inputs: make(map[string]bool)}
			}
			mods[output].inputs[mod.name] = false
		}
	}
	return mods, &pt
}

func part1() {
	mods, pt := initMods(input)
	for range 1000 {
		pt.add(pulse{src: "button", dst: "broadcaster", v: low})
		for p, ok := pt.pop(); ok; p, ok = pt.pop() {
			mods[p.dst].process(p)
		}
	}

	result := pt.cLow * pt.cHigh
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	mods, pt := initMods(input)
	result := "Multiply the first cycles of the four inputs to LS"
	for range 5_000 {
		pt.push++
		pt.add(pulse{src: "button", dst: "broadcaster", v: low})
		for p, ok := pt.pop(); ok; p, ok = pt.pop() {
			mods[p.dst].process(p)
		}
	}
	fmt.Printf("Part 2: %v\n", result)
}

type module struct {
	prefix  string
	name    string
	v       bool
	outputs []string
	inputs  map[string]bool
	pt      *pulseTracker
}

type pulseTracker struct {
	pending []pulse
	cLow    int
	cHigh   int
	push    int
}

func (pt *pulseTracker) add(p pulse) {
	switch p.v {
	case low:
		pt.cLow++
	case high:
		pt.cHigh++
	}
	pt.pending = append(pt.pending, p)
	if p.dst == "ls" && p.v {
		fmt.Printf("%s -%s-> %s (push %d)\n", p.src, func(v bool) string {
			if v {
				return "high"
			}
			return "low"
		}(p.v), p.dst, pt.push)
	}
}

func (pt *pulseTracker) pop() (pulse, bool) {
	if len(pt.pending) == 0 {
		return pulse{}, false
	}
	p := pt.pending[0]
	pt.pending = pt.pending[1:]
	return p, true
}

type pulse struct {
	src string
	dst string
	v   bool
}

const (
	low  = false
	high = true
)

func (m *module) process(p pulse) {
	switch {
	case m.name == "broadcaster":
		m.transmit(p.v)
	case m.prefix == "%":
		if p.v == high {
			return
		}
		m.v = !m.v
		m.transmit(m.v)
	case m.prefix == "&":
		m.inputs[p.src] = p.v
		for _, v := range m.inputs {
			if v == low {
				m.transmit(high)
				return
			}
		}
		m.transmit(low)
	}
}

func (m *module) transmit(v bool) {
	for _, output := range m.outputs {
		m.pt.add(pulse{src: m.name, dst: output, v: v})
	}
}
