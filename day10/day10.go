package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var bracket_friend = map[rune]rune{
	'<': '>',
	'[': ']',
	'(': ')',
	'{': '}',
}

var opening_brackets = map[rune]bool{
	'<': true,
	'[': true,
	'(': true,
	'{': true,
}

var error_points = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var auto_points = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func check_valid(line string) (bool, rune, rune, []rune) {
	chunk_stack := make([]rune, 0)
	for _, c := range line {
		_, is_opening := opening_brackets[c]
		if is_opening {
			chunk_stack = append(chunk_stack, c)
		} else {
			if c == bracket_friend[chunk_stack[len(chunk_stack)-1]] {
				chunk_stack = chunk_stack[:len(chunk_stack)-1]
			} else {
				return false, bracket_friend[chunk_stack[len(chunk_stack)-1]], c, chunk_stack
			}
		}
	}
	return true, '0', '0', chunk_stack
}

func auto_complete(chunk_stack []rune) int {
	score := 0
	for i := range chunk_stack {
		score *= 5
		score += auto_points[bracket_friend[chunk_stack[len(chunk_stack)-i-1]]]
	}
	return score
}

func main() {
	data, err := ioutil.ReadFile("day10/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	score := 0
	autocomp_scores := make([]int, 0)
	for _, l := range lines {
		is_valid, _, err_rune, chunk_stack := check_valid(l)
		if !is_valid {
			score += error_points[err_rune]
		} else {
			autocomp_scores = append(autocomp_scores, auto_complete(chunk_stack))
		}
	}
	fmt.Println(score)
	sort.Ints(autocomp_scores)
	fmt.Println(autocomp_scores[len(autocomp_scores)/2])
}
