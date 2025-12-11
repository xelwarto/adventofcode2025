package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
)

type Day5 int

var day5 = Day5(5)

func (d Day5) Part1() (string, error) {
	ingredients := []int64{}
	fresh := [][]int64{}
	s, err := util.File2Array("inputs/day5_part1.txt")
	if err != nil {
		return "", err
	}

	list := true
	for _, line := range s {
		if line == "" {
			list = false
		} else if list {
			r := strings.Split(line, "-")
			if len(r) != 2 {
				return "", fmt.Errorf("invalid ranges")
			}

			s_num, err := strconv.ParseInt(r[0], 10, 64)
			if err != nil {
				return "", err
			}
			l_num, err := strconv.ParseInt(r[1], 10, 64)
			if err != nil {
				return "", err
			}

			if s_num > l_num {
				return "", fmt.Errorf("invalid ranges")
			}
			fresh = append(fresh, []int64{s_num, l_num})
		} else {
			num, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				return "", err
			}
			ingredients = append(ingredients, num)
		}
	}

	// fmt.Println(len(ingredients))

	total := 0
	for _, i := range ingredients {
		// fmt.Println(i)
		found := false
		for _, rang := range fresh {
			if i >= rang[0] && i <= rang[1] {
				found = true
				break
			}
		}
		// fmt.Println(found)
		if found {
			total++
		}
	}

	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day5) Part2() (string, error) {
	ingredients := []int64{}
	fresh := [][]int64{}
	s, err := util.File2Array("inputs/day5_part2.txt")
	if err != nil {
		return "", err
	}

	list := true
	for _, line := range s {
		if line == "" {
			list = false
		} else if list {
			r := strings.Split(line, "-")
			if len(r) != 2 {
				return "", fmt.Errorf("invalid ranges")
			}

			s_num, err := strconv.ParseInt(r[0], 10, 64)
			if err != nil {
				return "", err
			}
			l_num, err := strconv.ParseInt(r[1], 10, 64)
			if err != nil {
				return "", err
			}

			if s_num > l_num {
				return "", fmt.Errorf("invalid ranges")
			}
			fresh = append(fresh, []int64{s_num, l_num})
		} else {
			num, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				return "", err
			}
			ingredients = append(ingredients, num)
		}
	}

	// fmt.Println(len(ingredients))

	f := [][]int64{}
	for _, i := range ingredients {
		// fmt.Println(i)
		// found := false
		for _, rang := range fresh {
			if i >= rang[0] && i <= rang[1] {
				f = append(f, rang)
			}
		}

	}

	fmt.Println(len(f))
	total := []int64{}

	for _, fu := range f {
		fmt.Println(fu)
		// for x := fu[0]; x <= fu[1]; x++ {

		// }
	}

	return fmt.Sprintf("%v", len(total)), nil
}

func (d Day5) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day5)
}
