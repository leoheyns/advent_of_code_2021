package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type coordinate struct {
	i int
	j int
}

func get_surrounding_coord(grid [][]int, coord coordinate) []coordinate {
	result := make([]coordinate, 0)
	i := coord.i
	j := coord.j
	if coord.j > 0 {
		result = append(result, coordinate{i: i, j: j - 1})
	}
	if j+1 < len(grid[0]) {
		result = append(result, coordinate{i: i, j: j + 1})
	}
	if i+1 < len(grid) {
		result = append(result, coordinate{i: i + 1, j: j})
	}
	if i > 0 {
		result = append(result, coordinate{i: i - 1, j: j})
	}
	return result
}

func basin_bfs(grid [][]int, low coordinate) int {
	seen := map[coordinate]bool{low: true}
	frontier := []coordinate{low}
	basin_size := 0
	var new_coord coordinate
	for len(frontier) > 0 {
		new_coord, frontier = frontier[0], frontier[1:]
		//fmt.Println(new_coord, frontier)
		if grid[new_coord.i][new_coord.j] != 9 {
			basin_size++
			adjacent := get_surrounding_coord(grid, new_coord)
			for _, adj_coord := range adjacent {
				_, is_seen := seen[adj_coord]
				if !is_seen {
					frontier = append(frontier, adj_coord)
					seen[adj_coord] = true
				}
			}
		}
	}
	return basin_size
}
func main() {
	data, err := ioutil.ReadFile("day9/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	grid := make([][]int, len(lines))
	for i, l := range lines {
		split_line := strings.Split(l, "")
		grid[i] = make([]int, len(split_line))
		for j, n := range split_line {
			grid[i][j], _ = strconv.Atoi(n)
		}
	}
	low_points := make([]coordinate, 0)
	low_sum := 0
	for i, row := range grid {
		for j, height := range row {
			low := true
			adj_points := get_surrounding_coord(grid, coordinate{i: i, j: j})

			for _, adj_point := range adj_points {
				if grid[adj_point.i][adj_point.j] <= height {
					low = false
				}
			}
			if low {
				//fmt.Println(j,i)
				low_sum += height + 1
				low_points = append(low_points, coordinate{i: i, j: j})
			}
		}

	}
	fmt.Println(low_sum)

	basin_sizes := make([]int, len(low_points))
	for i, lp := range low_points {
		basin_sizes[i] = basin_bfs(grid, lp)
		//fmt.Println(lp)
		//fmt.Println(basin_bfs(grid, lp))
	}
	sort.Ints(basin_sizes)
	largest := basin_sizes[len(basin_sizes)-3:]
	fmt.Println(largest[0] * largest[1] * largest[2])
}
