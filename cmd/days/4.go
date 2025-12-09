package days

import (
	"fmt"
)

type Day4 int

var day4 = Day4(4)

func (d Day4) Part1() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day4) Part2() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day4) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day4)
}
