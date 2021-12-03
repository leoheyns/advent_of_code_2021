package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day2/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	//lines = lines[:len(lines)-1]
	fmt.Println(len(lines))
	type Command struct {
		com string
		n   int
	}
	//fmt.Println(0, "  ", lines[0])
	commands := make([]Command, len(lines))
	for i, line := range lines {
		line_split := strings.Split(line, " ")
		//fmt.Println(line_split)
		num, _ := strconv.Atoi(line_split[1])
		commands[i] = Command{line_split[0], num}
	}

	depth := 0
	x := 0

	for _, command := range commands {
		if command.com == "forward" {
			x += command.n
		} else if command.com == "down" {
			depth += command.n
		} else {
			depth -= command.n
		}
	}
	//part 1
	fmt.Println(depth * x)

	depth = 0
	aim := 0

	for _, command := range commands {
		if command.com == "forward" {
			depth += aim * command.n
		} else if command.com == "down" {
			aim += command.n
		} else {
			aim -= command.n
		}
	}

	fmt.Println(depth * x)
}
