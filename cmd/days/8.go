package days

import (
	"fmt"
)

type Day8 int

var day8 = Day8(8)

func (d Day8) Part1() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day8) Part2() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day8) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day8)
}
