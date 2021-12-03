package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func count_inc(nums []int) int {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			//fmt.Println(i + 1, "  ", lines[i + 1], " (increased)")
			count++
		}
	}
	return count
}

func main() {
	data, err := ioutil.ReadFile("day1/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	//lines = lines[:len(lines)-1]
	fmt.Println(len(lines))

	//fmt.Println(0, "  ", lines[0])
	depths := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		depth, _ := strconv.Atoi(lines[i])
		depths[i] = depth
	}

	//count := 0
	//count_dec := 0
	//for i := 0; i < len(lines) - 1; i++{
	//	current, _ := strconv.Atoi(lines[i])
	//	next, _ := strconv.Atoi(lines[i + 1])
	//	if next > current{
	//		//fmt.Println(i + 1, "  ", lines[i + 1], " (increased)")
	//		count++
	//	} else{
	//		//fmt.Println(i + 1, "  ", lines[i + 1], " (decreased)")
	//		count_dec++
	//	}
	//}
	//fmt.Println(count, count_dec)

	fmt.Println(count_inc(depths))
	//part 2

	windows := make([]int, len(depths)-2)
	for i := 0; i < len(depths)-2; i++ {
		windows[i] = depths[i] + depths[i+1] + depths[i+2]
	}
	//fmt.Println(windows)
	fmt.Println(count_inc(windows))
}
