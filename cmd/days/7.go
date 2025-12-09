package days

import (
	"fmt"
)

type Day7 int

var day7 = Day7(7)

func (d Day7) Part1() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day7) Part2() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day7) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day7)
}
