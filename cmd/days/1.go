package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
)

type Day1 int

var day1 = Day1(1)

func (d Day1) Part1() (string, error) {
	s, err := util.File2Array("inputs/day1_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	pos := 50

	for _, line := range s {
		dir := string(line[0])
		snum := strings.TrimLeft(line, "LR")
		num, err := strconv.ParseInt(snum, 10, 64)
		if err != nil {
			return "", err
		}

		// fmt.Printf("%s / %s / %v\n", line, dir, int(num))

		if dir == "L" {
			for x := 0; x < int(num); x++ {
				pos--
				if pos < 0 {
					pos = 99
				}
			}
		}

		if dir == "R" {
			for x := 0; x < int(num); x++ {
				pos++
				if pos > 99 {
					pos = 0
				}
			}
		}

		// fmt.Println(pos)

		if pos == 0 {
			total++
		}
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day1) Part2() (string, error) {
	s, err := util.File2Array("inputs/day1_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	pos := 50

	for _, line := range s {
		dir := string(line[0])
		snum := strings.TrimLeft(line, "LR")
		num, err := strconv.ParseInt(snum, 10, 64)
		if err != nil {
			return "", err
		}

		// fmt.Printf("%s / %s / %v\n", line, dir, int(num))

		if dir == "L" {
			for x := 0; x < int(num); x++ {
				pos--
				if pos < 0 {
					pos = 99
				}
				if pos == 0 {
					total++
				}
			}
		}

		if dir == "R" {
			for x := 0; x < int(num); x++ {
				pos++
				if pos > 99 {
					pos = 0
				}
				if pos == 0 {
					total++
				}
			}
		}

		// fmt.Println(pos)
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day1) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day1)
}
