package days

import (
	"fmt"
)

type Day11 int

var day11 = Day11(11)

func (d Day11) Part1() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day11) Part2() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day11) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day11)
}
