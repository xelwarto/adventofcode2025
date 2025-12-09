package days

import (
	"fmt"
)

type Day5 int

var day5 = Day5(5)

func (d Day5) Part1() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day5) Part2() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day5) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day5)
}
