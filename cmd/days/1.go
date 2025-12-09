package days

import (
	"code/util"
	"fmt"
)

type Day1 int

var day1 = Day1(1)

func (d Day1) Part1() (string, error) {
	_, err := util.File2Array("inputs/day1_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day1) Part2() (string, error) {
	_, err := util.File2Array("inputs/day1_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day1) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day1)
}
