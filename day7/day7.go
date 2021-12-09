package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}
func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
func sumTo(n int) int {
	return n * (n + 1) / 2
}

func main() {
	data, err := ioutil.ReadFile("day7/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	init_state := strings.Split(lines[0], ",")
	positions := make([]int, len(init_state))
	sum := 0
	m := 0
	for i, num_string := range init_state {
		x, _ := strconv.Atoi(num_string)
		positions[i] = x
		sum += x
		m = max(m, x)
	}
	sort.Ints(positions)
	middle := positions[len(positions)/2]
	deviation := 0
	for _, pos := range positions {
		deviation += abs(pos - middle)
	}
	fmt.Println(deviation)
	min_deviation := -1
	min_dev_pos := 0
	for i := 0; i < m; i++ {
		other_deviation := 0
		for _, pos := range positions {
			other_deviation += sumTo(abs(pos - i))
		}
		if min_deviation == -1 || other_deviation < min_deviation {
			min_deviation = other_deviation
			min_dev_pos = i
		}
	}
	fmt.Println(min_deviation)
	fmt.Println(min_dev_pos)

}
