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

func printgrid(east map[coordinate]bool, south map[coordinate]bool, boundI int, boundJ int) {
	for i := 0; i < boundI; i++ {
		for j := 0; j < boundJ; j++ {
			_, ok1 := east[coordinate{i, j}]
			_, ok2 := south[coordinate{i, j}]
			if ok1 {
				print(">")
			} else if ok2 {
				print("v")
			} else {
				print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	data, err := ioutil.ReadFile("day25/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	east := map[coordinate]bool{}
	south := map[coordinate]bool{}
	boundI := len(lines)
	boundJ := len(lines[0])

	for i, line := range lines {
		for j, cell := range line {
			if cell == '>' {
				east[coordinate{i, j}] = true
			}
			if cell == 'v' {
				south[coordinate{i, j}] = true
			}
		}
	}
	steps := 0
	for true {
		//fmt.Println(steps)
		//printgrid(east, south, boundI, boundJ)
		changed := false
		newEast := map[coordinate]bool{}
		for coord := range east {
			i := coord.i
			j := coord.j

			_, ok1 := east[coordinate{i, (j + 1) % boundJ}]
			_, ok2 := south[coordinate{i, (j + 1) % boundJ}]
			if !ok1 && !ok2 {
				newEast[coordinate{i, (j + 1) % boundJ}] = true
				changed = true
			} else {
				newEast[coord] = true
			}

		}
		east = newEast

		newSouth := map[coordinate]bool{}
		for coord := range south {
			i := coord.i
			j := coord.j

			_, ok1 := east[coordinate{(i + 1) % boundI, j}]
			_, ok2 := south[coordinate{(i + 1) % boundI, j}]
			if !ok1 && !ok2 {
				newSouth[coordinate{(i + 1) % boundI, j}] = true
				changed = true
			} else {
				newSouth[coord] = true
			}

		}
		south = newSouth
		steps++
		if !changed {
			fmt.Println(steps)
			break
		}
	}

}
