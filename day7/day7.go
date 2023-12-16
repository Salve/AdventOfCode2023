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

var jokersEnabled bool

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
	jokersEnabled = true
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
		av, bv := a.hand[i], b.hand[i]
		if jokersEnabled && av == 11 {
			av = 1
		}
		if jokersEnabled && bv == 11 {
			bv = 1
		}
		if c := cmp.Compare(bv, av); c != 0 {
			return c
		}
	}
	fmt.Println("eq")
	return 0
}

func (e entry) handType() int {
	counts := make([]byte, 15)
	for _, v := range e.hand {
		counts[v]++
	}
	var jokerCount byte
	if jokersEnabled {
		jokerCount = counts[11]
		counts[11] = 0 // don't double count them
	}
	slices.Sort(counts)
	slices.Reverse(counts)
	switch {
	case counts[0]+jokerCount == 5:
		return 1
	case counts[0]+jokerCount == 4:
		return 2
	case counts[0]+jokerCount == 3 && counts[1] == 2:
		return 3
	case counts[0]+jokerCount == 3:
		return 4
	case counts[0] == 2 && counts[1]+jokerCount == 2:
		return 5
	case counts[0]+jokerCount == 2:
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

var cardValue = map[rune]byte{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
