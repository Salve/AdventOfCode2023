package day17

import (
	"container/heap"
	"fmt"
	"github.com/Salve/AdventOfCode2023/inputs"
	"github.com/Salve/AdventOfCode2023/registry"
	"image"
	"math"
	"strings"
)

const day = 17

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
	g, end := gridFromInput(input)
	result := g.navigate(1, 3, end)
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	g, end := gridFromInput(input)
	result := g.navigate(4, 10, end)
	fmt.Printf("Part 2: %v\n", result)
}

func (g grid) navigate(min, max int, end image.Point) int {
	down, right := image.Point{0, 1}, image.Point{1, 0}
	q, visited := PQ[State]{}, map[State]struct{}{}
	q.GPush(State{image.Point{0, 0}, down}, 0)
	q.GPush(State{image.Point{0, 0}, right}, 0)

	for len(q) > 0 {
		state, heat := q.GPop()
		if state.Pos == end {
			return heat
		}
		if _, seen := visited[state]; seen {
			continue
		}
		visited[state] = struct{}{}

		for i := -max; i <= max; i++ {
			n := state.Pos.Add(state.Dir.Mul(i))
			if _, ok := g[n]; !ok || i > -min && i < min {
				continue
			}
			addHeat, s := 0, int(math.Copysign(1, float64(i)))
			for j := s; j != i+s; j += s {
				addHeat += g[state.Pos.Add(state.Dir.Mul(j))]
			}
			q.GPush(State{n, image.Point{state.Dir.Y, state.Dir.X}}, heat+addHeat)
		}
	}
	return -1
}

type grid map[image.Point]int

func gridFromInput(input []byte) (grid, image.Point) {
	g, end := grid{}, image.Point{}
	for y, line := range strings.Split(string(input), "\n") {
		for x, v := range line {
			p := image.Point{x, y}
			g[p] = int(v - '0')
			end = p
		}
	}
	return g, end
}

type pqi[T any] struct {
	v T
	p int
}

type PQ[T any] []pqi[T]

func (q PQ[_]) Len() int           { return len(q) }
func (q PQ[_]) Less(i, j int) bool { return q[i].p < q[j].p }
func (q PQ[_]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *PQ[T]) Push(x any)        { *q = append(*q, x.(pqi[T])) }
func (q *PQ[_]) Pop() (x any)      { x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]; return x }
func (q *PQ[T]) GPush(v T, p int)  { heap.Push(q, pqi[T]{v, p}) }
func (q *PQ[T]) GPop() (T, int)    { x := heap.Pop(q).(pqi[T]); return x.v, x.p }

type State struct {
	Pos image.Point
	Dir image.Point
}
