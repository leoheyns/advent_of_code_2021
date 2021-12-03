package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func filter(pils []string, mc bool) string {
	if len(pils[0]) == 0 {
		return ""
	}
	count := 0
	fcount := 0
	for _, code := range pils {
		if code[0] == '1' {
			count += 1
		}
	}
	var t_bit uint8 = '0'
	if count >= len(pils)/2 {
		if mc {
			t_bit = '1'
		} else {
			t_bit = '0'
		}
		fcount = count
	} else {
		if mc {
			t_bit = '0'
		} else {
			t_bit = '1'
		}
		fcount = len(pils) - count
	}
	var remainder []string
	if mc {
		remainder = make([]string, fcount)
	} else {
		remainder = make([]string, len(pils)-fcount)
	}
	if count == len(pils) {
		remainder = make([]string, len(pils))
		t_bit = '1'
	}
	if count == 0 {
		remainder = make([]string, len(pils))
		t_bit = '0'
	}

	i := 0
	for _, code := range pils {
		if code[0] == t_bit {
			remainder[i] = code[1:]
			i++
		}
	}
	return string(t_bit) + filter(remainder, mc)
}

func main() {
	data, err := ioutil.ReadFile("day3/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	//lines = lines[:len(lines)-1]
	//fmt.Println(len(lines))
	nums := make([]int64, len(lines))
	for i, s := range lines {
		n, _ := strconv.ParseInt(s, 2, 64)
		nums[i] = n
	}
	counts := [12]int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, line := range lines {
		for j := 0; j < 12; j++ {
			if line[j] == '1' {
				counts[j]++
			}
		}
	}
	gamma := 0
	for i := 0; i < 12; i++ {
		//fmt.Println("num: ", counts[i])
		if counts[i] > int32(len(lines))/2 {
			//fmt.Println("1", 1 << (11 - i))
			gamma += 1 << (11 - i)
		}
	}

	epsilon := gamma ^ 0b111111111111
	//part1
	//fmt.Println(counts,int32(len(lines)) / 2, gamma, epsilon)
	fmt.Println(gamma * epsilon)

	ox, _ := strconv.ParseInt(filter(lines, true), 2, 64)
	co, _ := strconv.ParseInt(filter(lines, false), 2, 64)
	fmt.Println(ox * co)

}
