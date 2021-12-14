package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func fold(points map[point]bool, axis string, line int) {
	for p := range points {
		var newpoint point
		if axis == "x" {
			if p.x > line {
				newpoint.x = line - (p.x - line)
				newpoint.y = p.y
				delete(points, p)
				points[newpoint] = true
			}
		} else {
			if p.y > line {
				newpoint.y = line - (p.y - line)
				newpoint.x = p.x
				delete(points, p)
				points[newpoint] = true
			}
		}
	}
}

func print_points(points map[point]bool, max_x int, max_y int) {
	for y := 0; y < max_y; y++ {
		for x := 0; x < max_x; x++ {
			_, exists := points[point{x: x, y: y}]
			if exists {
				print("#")
			} else {
				print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	data, err := ioutil.ReadFile("day13/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	parts := strings.Split(string(data), "\n\n")
	point_strings := strings.Split(parts[0], "\n")
	fold_strings := strings.Split(parts[1], "\n")
	points := make(map[point]bool)
	for _, ps := range point_strings {
		split := strings.Split(ps, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		points[point{x: x, y: y}] = true
	}

	for _, fs := range fold_strings {
		fmt.Println(len(points))
		split := strings.Split(fs, "=")
		line, _ := strconv.Atoi(split[1])
		axis := string(split[0][11])
		fold(points, axis, line)
	}
	print_points(points, 40, 10)
}
