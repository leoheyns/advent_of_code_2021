package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day6/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	init_state := strings.Split(lines[0], ",")
	geilheid := [9]int{}
	for _, num_string := range init_state {
		x, _ := strconv.Atoi(num_string)
		geilheid[x]++
	}
	var geilheid_new [9]int
	for d := 0; d < 256; d++ {
		geilheid_new = [9]int{}
		for i, g := range geilheid {
			if i == 0 {
				geilheid_new[6] += g
				geilheid_new[8] += g
			} else {
				geilheid_new[i-1] += g
			}
		}
		geilheid = geilheid_new
	}
	sum := 0
	for _, g := range geilheid {
		sum += g
	}
	fmt.Println(sum)
}
