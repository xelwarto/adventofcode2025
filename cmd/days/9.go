package days

import (
	"fmt"
)

type Day9 int

var day9 = Day9(9)

func (d Day9) Part1() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day9) Part2() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day9) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day9)
}
