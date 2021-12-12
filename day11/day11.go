package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func cleanup(energy_levels [][]int, has_flashed [][]bool) int {
	flash_count := 0
	for i, l := range has_flashed {
		for j, flashed := range l {
			if flashed {
				has_flashed[i][j] = false
				energy_levels[i][j] = 0
				flash_count++
			}
		}
	}
	return flash_count
}

//func print_grid(energy_levels [][]int) {
//	for _, l := range energy_levels {
//		for _, n := range l {
//			print(n)
//		}
//		fmt.Println()
//	}
//}

func increment(energy_levels [][]int, has_flashed [][]bool, i int, j int) {
	energy_levels[i][j]++
	if energy_levels[i][j] > 9 && !has_flashed[i][j] {
		has_flashed[i][j] = true
		valid_is := []int{i}
		valid_js := []int{j}

		if i > 0 {
			valid_is = append(valid_is, i-1)
		}
		if i < len(energy_levels)-1 {
			valid_is = append(valid_is, i+1)
		}
		if j > 0 {
			valid_js = append(valid_js, j-1)
		}
		if j < len(energy_levels[0])-1 {
			valid_js = append(valid_js, j+1)
		}
		for _, i_adj := range valid_is {
			for _, j_adj := range valid_js {
				if !(i_adj == i && j_adj == j) {
					increment(energy_levels, has_flashed, i_adj, j_adj)
				}
			}
		}
	}
}

func main() {
	data, err := ioutil.ReadFile("day11/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	energy_levels := make([][]int, len(lines))
	has_flashed := make([][]bool, len(lines))
	for i, l := range lines {
		energy_levels[i] = make([]int, len(l))
		has_flashed[i] = make([]bool, len(l))
		for j, n_string := range l {
			n, _ := strconv.Atoi(string(n_string))
			energy_levels[i][j] = n
		}
	}
	total_flashes := 0
	for step := 0; step < 1000; step++ {
		for i := range energy_levels {
			for j := range energy_levels[i] {
				increment(energy_levels, has_flashed, i, j)
			}
		}
		flash_count := cleanup(energy_levels, has_flashed)
		if step < 100 {
			total_flashes += flash_count
		}
		if flash_count == 100 {
			fmt.Println("done", step+1)
			break
		}
		//print_grid(energy_levels)

	}
	fmt.Println(total_flashes)
}
