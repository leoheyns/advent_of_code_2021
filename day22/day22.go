package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func cubeOverlaps(c1 cuboid, c2 cuboid) bool {
	xOverlaps := c1.minx <= c2.maxx && c1.maxx >= c2.minx
	yOverlaps := c1.miny <= c2.maxy && c1.maxy >= c2.miny
	zOverlaps := c1.minz <= c2.maxz && c1.maxz >= c2.minz
	return xOverlaps && yOverlaps && zOverlaps
}

func subtrackt(c1 cuboid, c2 cuboid) []cuboid {
	result := []cuboid{}
	//cube which will be split
	splitCube := c1
	if cubeOverlaps(c1, c2) {
		if c2.maxx <= splitCube.maxx && c2.maxx >= splitCube.minx {
			split1 := splitCube
			split1.maxx = c2.maxx
			split2 := splitCube
			split2.minx = c2.maxx
			splitCube = split1
			result = append(result, split2)
		}

		if c2.minx <= splitCube.maxx && c2.minx >= splitCube.minx {
			split1 := splitCube
			split1.maxx = c2.minx
			split2 := splitCube
			split2.minx = c2.minx
			splitCube = split2
			result = append(result, split1)
		}

		if c2.maxy <= splitCube.maxy && c2.maxy >= splitCube.miny {
			split1 := splitCube
			split1.maxy = c2.maxy
			split2 := splitCube
			split2.miny = c2.maxy
			splitCube = split1
			result = append(result, split2)
		}

		if c2.miny <= splitCube.maxy && c2.miny >= splitCube.miny {
			split1 := splitCube
			split1.maxy = c2.miny
			split2 := splitCube
			split2.miny = c2.miny
			splitCube = split2
			result = append(result, split1)
		}

		if c2.maxz <= splitCube.maxz && c2.maxz >= splitCube.minz {
			split1 := splitCube
			split1.maxz = c2.maxz
			split2 := splitCube
			split2.minz = c2.maxz
			splitCube = split1
			result = append(result, split2)
		}

		if c2.minz <= splitCube.maxz && c2.minz >= splitCube.minz {
			split1 := splitCube
			split1.maxz = c2.minz
			split2 := splitCube
			split2.minz = c2.minz
			splitCube = split2
			result = append(result, split1)
		}
		return result

	} else {
		return []cuboid{c1}
	}
}

type cuboid struct {
	minx int
	maxx int
	miny int
	maxy int
	minz int
	maxz int
}

func main() {
	data, err := ioutil.ReadFile("day22/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	modes := make([]string, len(lines))
	cuboids := make([]cuboid, len(lines))

	stopAt := -1
	for i, l := range lines {
		modes[i] = strings.Split(l, " ")[0]
		cube := cuboid{}

		cuboidStrings := strings.Split(strings.Split(l, " ")[1], ",")

		cube.minx, _ = strconv.Atoi(strings.Split(cuboidStrings[0], "..")[0][2:])
		cube.maxx, _ = strconv.Atoi(strings.Split(cuboidStrings[0], "..")[1])
		cube.maxx++

		cube.miny, _ = strconv.Atoi(strings.Split(cuboidStrings[1], "..")[0][2:])
		cube.maxy, _ = strconv.Atoi(strings.Split(cuboidStrings[1], "..")[1])
		cube.maxy++

		cube.minz, _ = strconv.Atoi(strings.Split(cuboidStrings[2], "..")[0][2:])
		cube.maxz, _ = strconv.Atoi(strings.Split(cuboidStrings[2], "..")[1])
		cube.maxz++

		cuboids[i] = cube

		if stopAt == -1 {
			if cube.minx > 50 || cube.maxx < -50 || cube.miny > 50 || cube.maxy < -50 || cube.minz > 50 || cube.maxz < -50 {
				fmt.Println(cube)
				stopAt = i
			}
		}
	}

	result := []cuboid{}

	for i, cube := range cuboids {
		newResult := []cuboid{}

		for _, resultCube := range result {
			newResult = append(newResult, subtrackt(resultCube, cube)...)
		}
		if modes[i] == "on" {
			newResult = append(newResult, cube)
		}
		result = newResult
	}
	onCount := 0
	for _, cube := range result {
		onCount += (cube.maxx - cube.minx) * (cube.maxy - cube.miny) * (cube.maxz - cube.minz)
	}
	fmt.Println(onCount)

	//part1

	result = []cuboid{}
	for i := 0; i < stopAt; i++ {
		cube := cuboids[i]
		//fmt.Println(cube)
		newResult := []cuboid{}

		for _, resultCube := range result {
			newResult = append(newResult, subtrackt(resultCube, cube)...)
		}
		if modes[i] == "on" {
			newResult = append(newResult, cube)
		}
		result = newResult
	}
	onCount = 0
	for _, cube := range result {
		onCount += (cube.maxx - cube.minx) * (cube.maxy - cube.miny) * (cube.maxz - cube.minz)
	}
	fmt.Println(onCount)

}
