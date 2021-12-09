package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type display struct {
	hints   []string
	outputs []string
}

func string_overlap(s1 string, s2 string) int {
	count := 0
	for _, r := range s1 {
		if strings.ContainsRune(s2, r) {
			count++
		}
	}
	return count
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func get_num(disp display) int {
	string_to_num := make(map[string]int)
	one := ""
	four := ""
	for _, num_string_unsorted := range disp.hints {
		num_string := SortString(num_string_unsorted)
		if len(num_string) == 2 {
			one = num_string
			string_to_num[num_string] = 1
		} else if len(num_string) == 4 {
			four = num_string
			string_to_num[num_string] = 4
		} else if len(num_string) == 3 {
			string_to_num[num_string] = 7
		} else if len(num_string) == 7 {
			string_to_num[num_string] = 8
		}
	}
	for _, num_string_unsorted := range disp.hints {
		num_string := SortString(num_string_unsorted)
		if len(num_string) == 5 {
			if string_overlap(one, num_string) == 2 {
				string_to_num[num_string] = 3
			} else if string_overlap(four, num_string) == 3 {
				string_to_num[num_string] = 5
			} else {
				string_to_num[num_string] = 2
			}
		} else if len(num_string) == 6 {
			if string_overlap(one, num_string) != 2 {
				string_to_num[num_string] = 6
			} else if string_overlap(four, num_string) == 4 {
				string_to_num[num_string] = 9
			} else {
				string_to_num[num_string] = 0
			}
		}
	}
	//fmt.Println(string_to_num)
	//fmt.Println(disp.outputs)
	return 1000*string_to_num[SortString(disp.outputs[0])] + 100*string_to_num[SortString(disp.outputs[1])] + 10*string_to_num[SortString(disp.outputs[2])] + string_to_num[SortString(disp.outputs[3])]
}

func main() {
	data, err := ioutil.ReadFile("day8/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	displays := make([]display, len(lines))
	for i, line := range lines {
		spl := strings.Split(line, " | ")
		//fmt.Println(spl)
		displays[i] = display{hints: strings.Split(spl[0], " "), outputs: strings.Split(spl[1], " ")}
	}
	count_1_4_7_8 := 0
	for _, disp := range displays {
		for _, outp := range disp.outputs {
			if len(outp) == 2 || len(outp) == 3 || len(outp) == 4 || len(outp) == 7 {
				count_1_4_7_8++
			}
		}
	}
	fmt.Println(count_1_4_7_8)
	count_total := 0
	for _, disp := range displays {
		fmt.Println(get_num(disp))
		count_total += get_num(disp)
	}
	fmt.Println(count_total)
}
