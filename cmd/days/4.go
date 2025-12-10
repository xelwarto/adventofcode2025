package days

import (
	"code/util"
	"fmt"
	"strings"
)

type Day4 int

var day4 = Day4(4)

func (d Day4) Part1() (string, error) {
	s, err := util.File2Array("inputs/day4_part1.txt")
	if err != nil {
		return "", err
	}

	grid := [][]string{}
	for _, line := range s {
		points := strings.Split(line, "")
		grid = append(grid, points)
	}

	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			count := 0
			if grid[y][x] == "@" {
				if y != 0 {
					if grid[y-1][x] == "@" {
						count++
					}
					if x != 0 {
						if grid[y-1][x-1] == "@" {
							count++
						}
					}
					if x != len(grid[y])-1 {
						if grid[y-1][x+1] == "@" {
							count++
						}
					}
				}

				if x != 0 {
					if grid[y][x-1] == "@" {
						count++
					}
				}
				if x != len(grid[y])-1 {
					if grid[y][x+1] == "@" {
						count++
					}
				}

				if y != len(grid)-1 {
					if grid[y+1][x] == "@" {
						count++
					}
					if x != 0 {
						if grid[y+1][x-1] == "@" {
							count++
						}
					}
					if x != len(grid[y])-1 {
						if grid[y+1][x+1] == "@" {
							count++
						}
					}
				}
				if count < 4 {
					total++
				}
			}

		}
	}

	// for _, line1 := range grid {
	// 	fmt.Println(line1)
	// }

	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func (d Day4) Part2() (string, error) {
	s, err := util.File2Array("inputs/day4_part2.txt")
	if err != nil {
		return "", err
	}

	grid := [][]string{}

	for _, line := range s {
		points := strings.Split(line, "")
		grid = append(grid, points)
	}

	total := 0
	for {
		tracker := [][]int{}
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				count := 0
				if grid[y][x] == "@" {
					if y != 0 {
						if grid[y-1][x] == "@" {
							count++
						}
						if x != 0 {
							if grid[y-1][x-1] == "@" {
								count++
							}
						}
						if x != len(grid[y])-1 {
							if grid[y-1][x+1] == "@" {
								count++
							}
						}
					}

					if x != 0 {
						if grid[y][x-1] == "@" {
							count++
						}
					}
					if x != len(grid[y])-1 {
						if grid[y][x+1] == "@" {
							count++
						}
					}

					if y != len(grid)-1 {
						if grid[y+1][x] == "@" {
							count++
						}
						if x != 0 {
							if grid[y+1][x-1] == "@" {
								count++
							}
						}
						if x != len(grid[y])-1 {
							if grid[y+1][x+1] == "@" {
								count++
							}
						}
					}
					if count < 4 {
						pos := []int{y, x}
						tracker = append(tracker, pos)
						total++
					}
				}
			}
		}

		if len(tracker) == 0 {
			break
		} else {
			for _, t := range tracker {
				posy := t[0]
				posx := t[1]
				grid[posy][posx] = "x"
			}
		}

		// for _, line1 := range grid {
		// 	fmt.Println(line1)
		// }
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day4) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day4)
}
