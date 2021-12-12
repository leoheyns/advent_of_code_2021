package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func count_paths_part2(edges map[string]map[string]bool, visited map[string]bool, current string, has_visited_twice bool) int {
	if current == "end" {
		return 1
	}

	_, already_visited := visited[current]
	if already_visited && unicode.IsLower(rune(current[0])) {
		if has_visited_twice || current == "start" {
			return 0
		} else {
			has_visited_twice = true
		}
	}
	new_visited := make(map[string]bool)
	for k, v := range visited {
		new_visited[k] = v
	}
	new_visited[current] = true

	path_sum := 0
	for node := range edges[current] {
		path_sum += count_paths_part2(edges, new_visited, node, has_visited_twice)
	}
	return path_sum
}

func main() {
	data, err := ioutil.ReadFile("day12/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	edges := make(map[string]map[string]bool)
	for _, l := range lines {
		edge := strings.Split(l, "-")
		_, has_node1 := edges[edge[0]]
		if !has_node1 {
			edges[edge[0]] = make(map[string]bool)
		}
		_, has_node2 := edges[edge[1]]
		if !has_node2 {
			edges[edge[1]] = make(map[string]bool)
		}

		edges[edge[0]][edge[1]] = true
		edges[edge[1]][edge[0]] = true
	}
	visited := map[string]bool{}
	fmt.Println(count_paths_part2(edges, visited, "start", true))
	fmt.Println(count_paths_part2(edges, visited, "start", false))

}
