package days

import (
	"fmt"
)

type Day10 int

var day10 = Day10(10)

func (d Day10) Part1() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day10) Part2() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day10) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day10)
}
