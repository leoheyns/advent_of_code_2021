package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type vent_line struct {
	origin coord
	len    int
	dir    coord
}

//waarom de fuck moet ik dit zelf schrijven
func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func get_dir(x1 int, y1 int, x2 int, y2 int) coord {
	var dirx int
	var diry int
	if x2 > x1 {
		dirx = 1
	} else if x1 > x2 {
		dirx = -1
	}
	if y2 > y1 {
		diry = 1
	} else if y1 > y2 {
		diry = -1
	}
	return coord{x: dirx, y: diry}
}

func get_overlap(vent_lines []vent_line, max_x int, max_y int, part1 bool) int {
	vent_grid := make([][]int, max_x)
	for i := 0; i < max_x; i++ {
		vent_grid[i] = make([]int, max_y)
	}
	for _, vl := range vent_lines {
		//check if diagonal alleen voor part1
		if part1 && vl.dir.x != 0 && vl.dir.y != 0 {
			continue
		}
		current := vl.origin
		for i := 0; i < vl.len; i++ {
			vent_grid[current.x][current.y] += 1
			current.x += vl.dir.x
			current.y += vl.dir.y
		}
	}
	count_overlap := 0
	for _, vgl := range vent_grid {
		for _, vgp := range vgl {
			if vgp > 1 {
				count_overlap++
			}
		}
	}
	return count_overlap
}

func main() {
	data, err := ioutil.ReadFile("day5/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	vent_lines := make([]vent_line, len(lines))
	max_x := 0
	max_y := 0

	for i, l := range lines {
		re := regexp.MustCompile(",|( -> )")
		num_strings := re.Split(l, -1)
		nums := make([]int, len(num_strings))
		for j, ns := range num_strings {
			nums[j], _ = strconv.Atoi(ns)
		}
		max_x = max(max_x, max(nums[0], nums[2]))
		max_y = max(max_y, max(nums[1], nums[3]))

		vent_lines[i] = vent_line{origin: coord{
			x: nums[0], y: nums[1]},
			len: max(abs(nums[0]-nums[2]), abs(nums[1]-nums[3])) + 1,
			dir: get_dir(nums[0], nums[1], nums[2], nums[3]),
		}
	}
	max_x++
	max_y++
	fmt.Println(get_overlap(vent_lines, max_x, max_y, true))
	fmt.Println(get_overlap(vent_lines, max_x, max_y, false))
}
