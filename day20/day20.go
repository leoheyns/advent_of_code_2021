package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type coordinate struct {
	i int
	j int
}

func expandGrid(grid map[coordinate]bool) map[coordinate]bool {
	newGrid := map[coordinate]bool{}
	for coord, _ := range grid {
		for i := -2; i <= 2; i++ {
			for j := -2; j <= 2; j++ {
				if i != 0 || j != 0 {
					newGrid[coordinate{coord.i + i, coord.j + j}] = false
				}
			}
		}
	}
	for coord, val := range grid {
		newGrid[coord] = val
	}

	return newGrid
}

func updateGrid(grid map[coordinate]bool, algorithm []bool) map[coordinate]bool {
	newGrid := map[coordinate]bool{}
	count := 0
loop:
	for coord, _ := range grid {
		count++
		num := 0
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				val, ok := grid[coordinate{coord.i + i, coord.j + j}]
				if !ok {
					continue loop
				}
				if val {
					num += 1 << ((i+1)*3 + (j + 1))
				}
				newGrid[coord] = algorithm[num]
			}
		}
	}
	fmt.Println(count)
	return newGrid
}

func printGrid(grid map[coordinate]bool, maxI int, maxJ int) {
	for i := -maxI; i < maxI; i++ {
		for j := -maxJ; j < maxJ; j++ {
			val, ok := grid[coordinate{i, j}]
			if ok {
				if val {
					print("#")
				} else {
					print(".")
				}
			} else {
				print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	data, err := ioutil.ReadFile("day20/example_input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	parts := strings.Split(string(data), "\n\n")
	algorithmString := parts[0]
	algorithm := make([]bool, len(algorithmString))
	for i, char := range algorithmString {
		algorithm[i] = char == '#'
	}

	grid_string := parts[1]
	lines := strings.Split(grid_string, "\n")
	grid := map[coordinate]bool{}
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				grid[coordinate{i, j}] = true
			}
		}
	}

	for i := 0; i < 2; i++ {
		grid = expandGrid(grid)
		printGrid(grid, 6, 6)
		grid = updateGrid(grid, algorithm)
	}
	printGrid(grid, 6, 6)

	count := 0

	for _, val := range grid {
		if val {
			count++
		}
	}
	fmt.Println(count)

}
