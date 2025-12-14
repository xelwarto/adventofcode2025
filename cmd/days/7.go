package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
)

type Day7 int

var day7 = Day7(7)

func (d Day7) Part1() (string, error) {
	total := 0
	s, err := util.File2Array("inputs/day7_part1.txt")
	if err != nil {
		return "", err
	}

	field := [][]string{}
	for _, line := range s {
		field = append(field, strings.Split(line, ""))
	}

	// fmt.Println(len(field))

	for y := 0; y < len(field[0]); y++ {
		pos := field[0][y]
		if pos == "S" {
			field[1][y] = "|"
		}
	}

	for x := 0; x < len(field); x++ {
		for y := 0; y < len(field[x]); y++ {
			pos := field[x][y]
			if pos == "^" {
				if field[x-1][y] == "|" {
					field[x][y-1] = "|"
					field[x][y+1] = "|"
					total++
				}
			} else if x != 0 && field[x-1][y] == "|" {
				field[x][y] = "|"
			}
		}
		// fmt.Println(field[x])
	}

	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day7) Part2() (string, error) {
	s, err := util.File2Array("inputs/day7_part2.txt")
	if err != nil {
		return "", err
	}

	field := [][]string{}
	for _, line := range s {
		field = append(field, strings.Split(line, ""))
	}

	for y := 0; y < len(field[0]); y++ {
		pos := field[0][y]
		if pos == "S" {
			field[1][y] = "1"
		}
	}

	for x := 0; x < len(field); x++ {
		// fmt.Printf("%v\n", field[x])
		for y := 0; y < len(field[x]); y++ {
			pos := field[x][y]
			if pos == "^" {
				if field[x-1][y] != "." && field[x-1][y] != "^" {
					num1, err := strconv.ParseInt(field[x-1][y], 10, 64)
					if err != nil {
						return "", err
					}

					if field[x][y-1] == "." {
						// fmt.Printf("HERE2 ")
						if field[x-1][y-1] != "." && field[x-1][y-1] != "^" {
							num3, err := strconv.ParseInt(field[x-1][y-1], 10, 64)
							if err != nil {
								return "", err
							}
							field[x][y-1] = fmt.Sprintf("%v", num1+num3)
						} else {
							field[x][y-1] = fmt.Sprintf("%v", num1)
						}

					} else {
						num2, err := strconv.ParseInt(field[x][y-1], 10, 64)
						if err != nil {
							return "", err
						}
						field[x][y-1] = fmt.Sprintf("%v", num1+num2)

					}

					if field[x][y+1] == "." {
						// fmt.Printf("HERE-%v ", y)
						if field[x-1][y+1] != "." && field[x-1][y+1] != "^" {
							num3, err := strconv.ParseInt(field[x-1][y+1], 10, 64)
							if err != nil {
								return "", err
							}
							field[x][y+1] = fmt.Sprintf("%v", num1+num3)
						} else {
							field[x][y+1] = fmt.Sprintf("%v", num1)
						}
					} else {
						num2, err := strconv.ParseInt(field[x][y+1], 10, 64)
						if err != nil {
							return "", err
						}
						field[x][y+1] = fmt.Sprintf("%v", num1+num2)
					}

					// field[x][y-1] = "|"
					// field[x][y+1] = "|"

				}
			} else if x != 0 && field[x-1][y] != "." && field[x-1][y] != "^" {
				if field[x-1][y] == "S" {
					field[x][y] = "1"
				} else {
					if field[x][y] == "." {
						num, err := strconv.ParseInt(field[x-1][y], 10, 64)
						if err != nil {
							return "", err
						}
						field[x][y] = fmt.Sprintf("%v", num)
					}
				}

			}
		}
		// fmt.Printf("%v - %v\n", field[x], line_t)
		// fmt.Printf("%v\n", field[x])
		// fmt.Println(total)
	}

	total := int64(0)
	for y := 0; y < len(field[len(field)-1]); y++ {
		if field[len(field)-1][y] != "." && field[len(field)-1][y] != "^" {
			num, err := strconv.ParseInt(field[len(field)-1][y], 10, 64)
			if err != nil {
				return "", err
			}
			total = total + num
		}
	}

	return fmt.Sprintf("%v", total), nil

}

func (d Day7) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day7)
}
