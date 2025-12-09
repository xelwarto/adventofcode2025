package days

import (
	"fmt"
)

type Day6 int

var day6 = Day6(6)

func (d Day6) Part1() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day6) Part2() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day6) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day6)
}
