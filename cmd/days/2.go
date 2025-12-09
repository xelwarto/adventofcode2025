package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
)

type Day2 int

var day2 = Day2(2)

func (d Day2) Part1() (string, error) {
	s, err := util.File2Array("inputs/day2_part1.txt")
	if err != nil {
		return "", err
	}
	total := 0
	line := s[0]
	products := strings.Split(line, ",")
	for _, p := range products {
		if strings.Contains(p, "-") {
			nums := strings.Split(p, "-")
			if len(nums) == 2 {
				start, err := strconv.ParseInt(nums[0], 10, 64)
				if err != nil {
					return "", err
				}

				end, err := strconv.ParseInt(nums[1], 10, 64)
				if err != nil {
					return "", err
				}

				// fmt.Printf("%v - %v\n", start, end)

				for x := start; x <= end; x++ {
					sx := strconv.Itoa(int(x))
					if (len(sx) % 2) == 0 {
						hl := len(sx) / 2

						s_start := sx[0:hl]
						count := strings.Count(sx, s_start)

						if count == 2 {
							total = total + int(x)
						}

						if count > 2 {
							// fmt.Printf("%s - %v / %s / %v\n", sx, hl, s_start, count)
						}

						// fmt.Printf("%s - %v / %s / %v\n", sx, hl, s_start, count)
					}
				}
			} else {
				return "", fmt.Errorf("invalid input")
			}
		} else {
			return "", fmt.Errorf("invalid input")
		}
	}

	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day2) Part2() (string, error) {
	total := 0
	return fmt.Sprintf("%v", total), nil
}

func (d Day2) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day2)
}
