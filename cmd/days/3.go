package days

import (
	"fmt"
)

type Day3 int

var day3 = Day3(3)

func (d Day3) Part1() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day3) Part2() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day3) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day3)
}
