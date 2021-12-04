package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

//waarom doe ik dit mezelf aan

type bingo_result struct {
	score int
	count int
	index int
}

func has_bingo(checked_spaces [5][5]bool) bool {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !checked_spaces[y][x] {
				break
			}
			if y == 4 {
				return true
			}
		}
	}
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !checked_spaces[y][x] {
				break
			}
			if x == 4 {
				return true
			}
		}
	}
	return false
}

func calc_score(checked_spaces [5][5]bool, card [5][5]int) int {
	result := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !checked_spaces[y][x] {
				result += card[y][x]
			}
		}
	}
	return result
}

func go_gadget_checkbingo(card [5][5]int, numbers []int, index int, out chan bingo_result) {
	var checked_spaces [5][5]bool
	result := bingo_result{
		score: 0,
		count: -1,
		index: index,
	}
	for i, num := range numbers {
		for y, l := range card {
			for x, cnum := range l {
				if cnum == num {
					checked_spaces[y][x] = true
				}
			}
		}
		if has_bingo(checked_spaces) {
			result.count = i
			result.score = calc_score(checked_spaces, card)
			out <- result
			return
		}
	}
	out <- result
}

func main() {
	data, err := ioutil.ReadFile("day4/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	blocks := strings.Split(string(data), "\n\n")

	num_strings := strings.Split(blocks[0], ",")
	numbers := make([]int, len(num_strings))
	for i, s := range num_strings {
		numbers[i], _ = strconv.Atoi(s)
	}
	fmt.Println(numbers)

	bingo_strings := blocks[1:]

	bingo_cards := make([][5][5]int, len(bingo_strings))

	for i, bs := range bingo_strings {
		//fmt.Println(bs)
		//kutspaties
		if bs[0] == ' ' {
			bs = bs[1:]
		}
		re1 := regexp.MustCompile("\n *")
		bs_lines := re1.Split(string(bs), -1)
		for y, l := range bs_lines {
			re2 := regexp.MustCompile(" +")
			bs_nums := re2.Split(string(l), -1)
			for x, n := range bs_nums {
				bingo_cards[i][x][y], _ = strconv.Atoi(n)
			}
		}
	}
	out := make(chan bingo_result)
	for i, card := range bingo_cards {
		go go_gadget_checkbingo(card, numbers, i, out)
	}

	results := make([]bingo_result, len(bingo_cards))
	worst_count := -1
	worst_score := 0
	worst_index := -1

	best_count := -1
	best_score := 0
	best_index := -1
	empty := bingo_result{score: 0, count: 0, index: 0}
	for i := 0; i < len(bingo_cards); i++ {

		res := <-out

		if results[res.index] != empty {
			fmt.Println("wtf")
			fmt.Println(results[res.index], res.index)
		}
		results[res.index] = res

		if res.count == -1 {
			continue
		}
		if (best_count == -1) || (res.count < best_count) {
			best_count = res.count
			best_score = res.score
			best_index = res.index
		}

		if (worst_count == -1) || (res.count > worst_count) {
			worst_count = res.count
			worst_score = res.score
			worst_index = res.index
		}
	}
	fmt.Println(best_index, best_count, best_score)
	fmt.Println(best_score * numbers[best_count])

	fmt.Println(worst_index, worst_count, worst_score)
	fmt.Println(worst_score * numbers[worst_count])
}
