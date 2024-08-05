package main

import (
	"fmt"
	"time"

	_ "github.com/Salve/AdventOfCode2023/day1"
	_ "github.com/Salve/AdventOfCode2023/day10"
	_ "github.com/Salve/AdventOfCode2023/day11"
	_ "github.com/Salve/AdventOfCode2023/day12"
	_ "github.com/Salve/AdventOfCode2023/day13"
	_ "github.com/Salve/AdventOfCode2023/day14"
	_ "github.com/Salve/AdventOfCode2023/day15"
	_ "github.com/Salve/AdventOfCode2023/day16"
	_ "github.com/Salve/AdventOfCode2023/day17"
	_ "github.com/Salve/AdventOfCode2023/day18"
	_ "github.com/Salve/AdventOfCode2023/day19"
	_ "github.com/Salve/AdventOfCode2023/day2"
	_ "github.com/Salve/AdventOfCode2023/day20"
	_ "github.com/Salve/AdventOfCode2023/day21"
	_ "github.com/Salve/AdventOfCode2023/day3"
	_ "github.com/Salve/AdventOfCode2023/day4"
	_ "github.com/Salve/AdventOfCode2023/day5"
	_ "github.com/Salve/AdventOfCode2023/day6"
	_ "github.com/Salve/AdventOfCode2023/day7"
	_ "github.com/Salve/AdventOfCode2023/day8"
	_ "github.com/Salve/AdventOfCode2023/day9"
	"github.com/Salve/AdventOfCode2023/registry"
)

func main() {
	name, f := registry.Last()
	fmt.Printf("--- Running last day (%d) ---\n", name)
	timeFunc(f)
}

func timeFunc(f func()) time.Duration {
	start := time.Now()
	f()
	d := time.Now().Sub(start)
	fmt.Printf("--- Execution time: %s ---\n", d)
	return d
}
