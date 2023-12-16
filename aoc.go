package main

import (
	"fmt"
	"time"

	_ "github.com/Salve/AdventOfCode2023/day1"
	_ "github.com/Salve/AdventOfCode2023/day10"
	_ "github.com/Salve/AdventOfCode2023/day2"
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
