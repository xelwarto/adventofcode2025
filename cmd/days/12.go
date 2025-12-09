package days

import (
	"fmt"
)

type Day12 int

var day12 = Day12(12)

func (d Day12) Part1() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day12) Part2() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day12) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day12)
}
