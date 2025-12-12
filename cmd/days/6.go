package days

import (
	"code/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day6 int

var day6 = Day6(6)

func (d Day6) Part1() (string, error) {
	s, err := util.File2Array("inputs/day6_part1.txt")
	if err != nil {
		return "", err
	}

	data := [][]int64{}
	var spaces = regexp.MustCompile(`\s+`)
	for x := 0; x < len(s)-1; x++ {
		data = append(data, []int64{})
		line := spaces.Split(strings.TrimSpace(s[x]), -1)
		for _, l := range line {
			num, err := strconv.ParseInt(l, 10, 64)
			if err != nil {
				return "", err
			}
			data[x] = append(data[x], num)
		}
		fmt.Println(len(data[x]))
	}
	operators := []string{}
	o_line := spaces.Split(strings.TrimSpace(s[len(s)-1]), -1)
	operators = append(operators, o_line...)

	fmt.Println(len(operators))

	total := int64(0)
	for y := 0; y < len(operators); y++ {
		op := operators[y]

		t := data[0][y]
		for x := 1; x < len(s)-1; x++ {
			if op == "+" {
				t = t + data[x][y]
			}
			if op == "*" {
				t = t * data[x][y]
			}
		}
		fmt.Println(t)
		total = total + t
	}

	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day6) Part2() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day6) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day6)
}
