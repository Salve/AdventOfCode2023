package day7

import (
	"cmp"
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"slices"
	"strconv"
	"strings"
)

const day = 7

var input []byte

var example = []byte(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
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
	var entries []entry
	for _, line := range inputs.Lines(input) {
		entries = append(entries, entryFromLine(line))
	}
	winnings := 0
	slices.SortFunc(entries, rankHands)
	for i, e := range entries {
		winnings += (i + 1) * e.bid
	}
	result := winnings
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	result := "TODO"
	fmt.Printf("Part 2: %v\n", result)
}

func rankHands(b, a entry) int {
	if typeRank := rankType(a, b); typeRank != 0 {
		return typeRank
	}
	return rankCards(a, b)
}

func rankType(a, b entry) int {
	return cmp.Compare(a.handType(), b.handType())
}

func rankCards(a, b entry) int {
	for i := 0; i < 5; i++ {
		if c := cmp.Compare(a.hand[i], b.hand[i]); c != 0 {
			return c * -1
		}
	}
	return 0
}

func (e entry) handType() int {
	counts := make([]byte, 13)
	for _, v := range e.hand {
		counts[v]++
	}
	slices.Sort(counts)
	slices.Reverse(counts)
	switch {
	case counts[0] == 5:
		return 1
	case counts[0] == 4:
		return 2
	case counts[0] == 3 && counts[1] == 2:
		return 3
	case counts[0] == 3:
		return 4
	case counts[0] == 2 && counts[1] == 2:
		return 5
	case counts[0] == 2:
		return 6
	default:
		return 7
	}
}

type entry struct {
	hand [5]byte
	bid  int
}

func entryFromLine(line string) (e entry) {
	split := strings.Split(line, " ")
	e.bid, _ = strconv.Atoi(split[1])
	for i, r := range split[0] {
		e.hand[i] = cardValue[r]
	}
	return e
}

var cardValue = map[rune]byte{'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7, 'T': 8, 'J': 9, 'Q': 10, 'K': 11, 'A': 12}
