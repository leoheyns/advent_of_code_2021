package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func merge_counters(c1 map[string]int, c2 map[string]int) map[string]int {
	result := make(map[string]int)
	for pair, count1 := range c1 {
		count2, ok := c2[pair]
		if ok {
			result[pair] = count1 + count2
		} else {
			result[pair] = count1
		}
	}

	for pair, count1 := range c2 {
		count2, ok := c1[pair]
		if ok {
			result[pair] = count1 + count2
		} else {
			result[pair] = count1
		}
	}

	return result
}

func count_characters(s string) map[string]int {
	frequencies := make(map[string]int)
	for _, char := range s {
		_, ok := frequencies[string(char)]
		if ok {
			frequencies[string(char)]++
		} else {
			frequencies[string(char)] = 1
		}
	}
	return frequencies
}

func main() {
	data, err := ioutil.ReadFile("day14/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	parts := strings.Split(string(data), "\n\n")
	orig_template := strings.Split(parts[0], "\n")[0]
	rules_strings := strings.Split(parts[1], "\n")
	rules := make(map[string]string)
	for _, rs := range rules_strings {
		rs_split := strings.Split(string(rs), " -> ")
		rules[rs_split[0]] = rs_split[1]
	}

	//initial naive implementation
	template := orig_template
	for i := 0; i < 10; i++ {
		next_template := ""
		for j := 0; j < len(template)-1; j++ {
			val, ok := rules[template[j:j+2]]
			next_template += string(template[j])
			if ok {
				next_template += val
			}
		}
		next_template += string(template[len(template)-1])
		template = next_template
	}

	frequencies := make(map[string]int)
	for _, char := range template {
		_, ok := frequencies[string(char)]
		if ok {
			frequencies[string(char)]++
		} else {
			frequencies[string(char)] = 1
		}
	}

	mf_char := string(template[0])
	lf_char := string(template[0])

	for char, count := range frequencies {
		if count > frequencies[mf_char] {
			mf_char = char
		}
		if count < frequencies[lf_char] {
			lf_char = char
		}
	}

	fmt.Println(frequencies)
	fmt.Println(frequencies[mf_char] - frequencies[lf_char])

	//better implementation
	cached_frequencies := make([]map[string]map[string]int, 40)
	for i := 0; i < 40; i++ {
		cached_frequencies[i] = make(map[string]map[string]int)
		for pair, value := range rules {
			inserted_string := string(pair[0]) + value
			if i == 0 {
				cached_frequencies[i][pair] = count_characters(inserted_string)
			} else {
				cached_frequencies[i][pair] = merge_counters(cached_frequencies[i-1][string(pair[0])+value], cached_frequencies[i-1][value+string(pair[1])])
			}
		}
	}
	final_counts := make(map[string]int)
	for j := 0; j < len(orig_template)-1; j++ {
		final_counts = merge_counters(final_counts, cached_frequencies[39][orig_template[j:j+2]])
	}
	final_counts[string(orig_template[len(orig_template)-1])]++

	mf_char = string(orig_template[0])
	lf_char = string(orig_template[0])

	for char, count := range final_counts {
		if count > final_counts[mf_char] {
			mf_char = char
		}
		if count < final_counts[lf_char] {
			lf_char = char
		}
	}
	fmt.Println(final_counts)
	fmt.Println(final_counts[mf_char] - final_counts[lf_char])
}
