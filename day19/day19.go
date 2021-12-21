package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
	z int
}

func pointCloudFromString(s string) map[coordinate]bool {
	result := map[coordinate]bool{}
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		if i == 0 {
			continue
		}
		x, _ := strconv.Atoi(strings.Split(l, ",")[0])
		y, _ := strconv.Atoi(strings.Split(l, ",")[1])
		z, _ := strconv.Atoi(strings.Split(l, ",")[2])
		result[coordinate{x, y, z}] = true
	}
	return result
}

func permutateCloud(cloud map[coordinate]bool, x_axis int, y_axis int, x_sign int, y_sign int) map[coordinate]bool {
	result := map[coordinate]bool{}
	for coord := range cloud {
		new_coord := coordinate{
			x: 0,
			y: 0,
			z: 0,
		}
		switch x_axis {
		case 0:
			new_coord.x = coord.x * x_sign
		case 1:
			new_coord.y = coord.x * x_sign
		case 2:
			new_coord.z = coord.x * x_sign
		}

		switch y_axis {
		case 0:
			new_coord.x = coord.y * y_sign
		case 1:
			new_coord.y = coord.y * y_sign
		case 2:
			new_coord.z = coord.y * y_sign
		}
		z_sign := x_sign * y_sign

		if (x_axis == 0 && y_axis == 2) || (x_axis == 1 && y_axis == 0) || (x_axis == 2 && y_axis == 1) {
			z_sign = -z_sign
		}

		switch 3 - x_axis - y_axis {
		case 0:
			new_coord.x = coord.z * z_sign
		case 1:
			new_coord.y = coord.z * z_sign
		case 2:
			new_coord.z = coord.z * z_sign
		}
		result[new_coord] = true
	}
	return result
}

func generatePossibleClouds(cloud map[coordinate]bool) []map[coordinate]bool {
	result := []map[coordinate]bool{}
	for x_axis := 0; x_axis < 3; x_axis++ {
		for y_axis := 0; y_axis < 3; y_axis++ {
			if x_axis != y_axis {
				result = append(result, permutateCloud(cloud, x_axis, y_axis, 1, 1))
				result = append(result, permutateCloud(cloud, x_axis, y_axis, -1, 1))
				result = append(result, permutateCloud(cloud, x_axis, y_axis, 1, -1))
				result = append(result, permutateCloud(cloud, x_axis, y_axis, -1, -1))
			}
		}
	}
	return result
}

func checkOverlap(cloud1 map[coordinate]bool, cloud2 map[coordinate]bool) (bool, coordinate) {
	for anchor1 := range cloud1 {
		for anchor2 := range cloud2 {
			offset := coordinate{
				x: anchor1.x - anchor2.x,
				y: anchor1.y - anchor2.y,
				z: anchor1.z - anchor2.z,
			}

			overlap_count := 0
			for point := range cloud2 {
				offsetPoint := coordinate{
					x: point.x + offset.x,
					y: point.y + offset.y,
					z: point.z + offset.z,
				}
				_, ok := cloud1[offsetPoint]
				if ok {
					overlap_count++
				} else {
					break
				}
			}
			if overlap_count >= 12 {
				return true, offset
			}
		}
	}
	return false, coordinate{}
}

func resultPointCloud(pointClouds []*map[coordinate]bool) (map[coordinate]bool, map[*map[coordinate]bool]coordinate) {

	added := map[*map[coordinate]bool]bool{pointClouds[0]: true}
	total_offsets := map[*map[coordinate]bool]coordinate{pointClouds[0]: {
		x: 0,
		y: 0,
		z: 0,
	}}

	result := map[coordinate]bool{}

	for point := range *pointClouds[0] {
		result[point] = true
	}
	queue := []*map[coordinate]bool{pointClouds[0]}
	for len(added) < len(pointClouds) {
		currentReferenceScanner := queue[0]
		queue = queue[1:]

		for _, cloud := range pointClouds {
			_, ok := added[cloud]
			if !ok {
				var cloudRotated map[coordinate]bool
				for _, cloudPerm := range generatePossibleClouds(*cloud) {
					overlaps, offset := checkOverlap(*currentReferenceScanner, cloudPerm)
					if overlaps {
						cloudRotated = cloudPerm
						queue = append(queue, &cloudRotated)
						added[cloud] = true

						total_offsets[&cloudRotated] = coordinate{
							x: offset.x + total_offsets[currentReferenceScanner].x,
							y: offset.y + total_offsets[currentReferenceScanner].y,
							z: offset.z + total_offsets[currentReferenceScanner].z,
						}
						for point := range cloudPerm {
							offsetPoint := coordinate{
								x: point.x + total_offsets[&cloudRotated].x,
								y: point.y + total_offsets[&cloudRotated].y,
								z: point.z + total_offsets[&cloudRotated].z,
							}
							result[offsetPoint] = true
						}
						break
					}
				}
			}
		}
	}
	return result, total_offsets
}

func manhattan(coord1 coordinate, coord2 coordinate) int {
	x := coord1.x - coord2.x
	if x < 0 {
		x = -x
	}

	y := coord1.y - coord2.y
	if y < 0 {
		y = -y
	}

	z := coord1.z - coord2.z
	if z < 0 {
		z = -z
	}

	return x + y + z
}

func main() {
	data, err := ioutil.ReadFile("day19/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	scanner_strings := strings.Split(string(data), "\n\n")
	clouds := []*map[coordinate]bool{}
	for _, scannerString := range scanner_strings {
		cloud := pointCloudFromString(scannerString)
		clouds = append(clouds, &cloud)
	}
	allBeacons, scannerOffsets := resultPointCloud(clouds)
	fmt.Println(len(allBeacons))
	max := 0
	for _, coord1 := range scannerOffsets {
		for _, coord2 := range scannerOffsets {
			dist := manhattan(coord1, coord2)
			if dist > max {
				max = dist
			}
		}
	}
	fmt.Println(max)
}
