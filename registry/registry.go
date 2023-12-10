package registry

import (
	"fmt"
	"sort"
)

var days map[int]func()

func Register(day int, f func()) {
	if days == nil {
		days = make(map[int]func())
	}
	days[day] = f
}

func Last() (int, func()) {
	if len(days) == 0 {
		return 0, func() {
			fmt.Println("uninitialized")
		}
	}

	keys := make([]int, 0, len(days))
	for k := range days {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	last := keys[len(keys)-1]
	return last, days[last]
}
