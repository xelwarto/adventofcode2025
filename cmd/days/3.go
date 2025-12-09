package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
)

type Day3 int

var day3 = Day3(3)

func (d Day3) Part1() (string, error) {
	s, err := util.File2Array("inputs/day3_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	for _, line := range s {
		fmt.Println(line)
		f_num := 0
		s_num := 0
		nums := strings.Split(line, "")
		for x, snum := range nums {
			num, err := strconv.ParseInt(snum, 10, 64)
			if err != nil {
				return "", err
			}
			if x == 0 {
				if int(num) > f_num {
					f_num = int(num)
				}
			} else if x == len(nums)-1 {
				if int(num) > s_num {
					s_num = int(num)
				}
			} else {
				if int(num) > f_num {
					f_num = int(num)
					s_num = 0
				} else if int(num) > s_num {
					s_num = int(num)
				}
			}
		}

		if f_num == 0 || s_num == 0 {
			return "", fmt.Errorf("result ended in 0")
		}

		total = total + ((f_num * 10) + s_num)

		fmt.Printf("%v / %v\n", f_num, s_num)
	}

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
