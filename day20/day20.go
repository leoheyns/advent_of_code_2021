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

func expandGrid(grid map[coordinate]bool, value bool) map[coordinate]bool {
	newGrid := map[coordinate]bool{}
	for coord := range grid {
		for i := -2; i <= 2; i++ {
			for j := -2; j <= 2; j++ {
				if i != 0 || j != 0 {
					newGrid[coordinate{coord.i + i, coord.j + j}] = value
				}
			}
		}
	}
	for coord, val := range grid {
		newGrid[coord] = val
	}

	return newGrid
}

func updateGrid(grid map[coordinate]bool, algorithm []bool, infValue bool) map[coordinate]bool {
	newGrid := map[coordinate]bool{}
	count := 0
	for coord, _ := range grid {
		count++
		num := 0
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				val, ok := grid[coordinate{coord.i + i, coord.j + j}]
				if !ok {
					val = infValue
				}
				if val {
					num += 1 << (8 - (i+1)*3 - (j + 1))
				}
				newGrid[coord] = algorithm[num]
			}
		}
		if coord.i == 2 && coord.j == 2 {
			fmt.Println(coord, num)
		}
	}
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
	data, err := ioutil.ReadFile("day20/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	parts := strings.Split(string(data), "\n\n")
	algorithmStringLines := strings.Split(parts[0], "\n")
	algorithmString := ""
	for _, line := range algorithmStringLines {
		algorithmString += line
	}
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

	infValue := false

	for i := 0; i < 50; i++ {
		grid = expandGrid(grid, infValue)
		//printGrid(grid, 10, 10)
		//fmt.Println("-------------------------------")
		grid = updateGrid(grid, algorithm, infValue)
		infValue = !infValue

	}
	//printGrid(grid, 10, 10)

	count := 0

	for _, val := range grid {
		if val {
			count++
		}
	}
	fmt.Println(count)

}
