package days

import (
	"code/util"
	"fmt"
	"strings"
)

type Day11 int

var day11 = Day11(11)
var flow map[string][]string
var total int

func proc(list []string) {
	for _, l := range list {
		if l == "out" {
			total++
		}
		if flow[l] != nil {
			proc(flow[l])
		}
	}
}

func (d Day11) Part1() (string, error) {
	total = 0
	s, err := util.File2Array("inputs/day11_part1.txt")
	if err != nil {
		return "", err
	}

	flow = make(map[string][]string)
	for _, line := range s {
		data1 := strings.Split(line, ":")
		if len(data1) != 2 {
			return "", fmt.Errorf("input data error")
		}
		data2 := strings.Split(strings.TrimSpace(data1[1]), " ")
		flow[data1[0]] = data2
	}

	// fmt.Println(flow)
	proc(flow["you"])

	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################
var threads []string

func proc2(list []string, thread string) {
	for _, l := range list {
		new_t := fmt.Sprintf("%s,%s", thread, l)
		if l == "svr" {
			threads = append(threads, new_t)
			fmt.Println(new_t)
		}
		if flow[l] != nil {
			proc2(flow[l], new_t)
		}
	}
}

func (d Day11) Part2() (string, error) {
	total = 0
	s, err := util.File2Array("inputs/day11_part2.txt")
	if err != nil {
		return "", err
	}

	flow = make(map[string][]string)
	for _, line := range s {
		data1 := strings.Split(line, ":")
		if len(data1) != 2 {
			return "", fmt.Errorf("input data error")
		}
		data2 := strings.Split(strings.TrimSpace(data1[1]), " ")
		flow[data1[0]] = data2
	}

	// fmt.Println(flow)
	threads = []string{}
	proc2(flow["ggg"], "out,ggg")

	for _, t := range threads {
		fmt.Println(t)
		if strings.Contains(t, "fft") && strings.Contains(t, "dac") {
			total++
		}
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day11) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day11)
}
