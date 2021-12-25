package main

import (
	"fmt"
	"strconv"
	"strings"
)

var reg = map[string]int{"w": 0, "x": 1, "y": 2, "z": 3}

func applyInstruction(instruction string, state [4]int) [4]int {
	splt := strings.Split(instruction, " ")
	newState := state
	switch splt[0] {
	case "add":
		val, err := strconv.Atoi(splt[2])

		if err != nil {
			val = state[reg[splt[2]]]
		}
		newState[reg[splt[1]]] = state[reg[splt[1]]] + val
	case "mul":
		val, err := strconv.Atoi(splt[2])
		if err != nil {
			val = state[reg[splt[2]]]
		}
		newState[reg[splt[1]]] = state[reg[splt[1]]] * val
	case "div":
		val, err := strconv.Atoi(splt[2])
		if err != nil {
			val = state[reg[splt[2]]]
		}
		newState[reg[splt[1]]] = state[reg[splt[1]]] / val
	case "mod":
		val, err := strconv.Atoi(splt[2])
		if err != nil {
			val = state[reg[splt[2]]]
		}
		newState[reg[splt[1]]] = state[reg[splt[1]]] % val
	case "eql":
		val, err := strconv.Atoi(splt[2])
		if err != nil {
			val = state[reg[splt[2]]]
		}
		if state[reg[splt[1]]] == val {
			newState[reg[splt[1]]] = 1
		} else {
			newState[reg[splt[1]]] = 0
		}
	}
	return newState
}

func greaterThan(inputs1 []int, inputs2 []int) bool {
	for i, input := range inputs1 {
		if input > inputs2[i] {
			return true
		}
	}
	return false
}

func iter(instruction string, states map[[4]int][]int) map[[4]int][]int {
	newStates := map[[4]int][]int{}
	splt := strings.Split(instruction, " ")
	if splt[0] == "inp" {
		for state, inputs := range states {
			if state[3] != 0 {
				continue
			}
			for i := 1; i < 10; i++ {
				newState := state
				newState[reg[splt[1]]] = i
				newInputs := make([]int, len(inputs))
				copy(newInputs, inputs)
				newInputs = append(newInputs, i)

				other, ok := newStates[newState]
				if ok {
					if greaterThan(inputs, other) {
						newStates[newState] = newInputs
					}
				} else {
					newStates[newState] = newInputs
				}
			}
		}
	} else {
		for state, inputs := range states {
			new := applyInstruction(instruction, state)
			other, ok := newStates[new]
			if ok {
				if greaterThan(inputs, other) {
					newStates[new] = inputs
				}
			} else {
				newStates[new] = inputs
			}
		}
	}
	return newStates
}

func main() {
	//data, err := ioutil.ReadFile("day24/input")
	//if err != nil {
	//	fmt.Println("File reading error", err)
	//	return
	//}
	//lines := strings.Split(string(data), "\n")
	//states := map[[4]int][]int{[4]int{0,0,0,0}:[]int{}}
	//for _,line := range lines{
	//	states = iter(line, states)
	//	fmt.Println(line, len(states))
	//	//if i == 7{
	//	//	fmt.Println(line)
	//	//	break
	//	//}
	//}
	//fmt.Println(states)
	N1s := []int{14, 10, 13, -8, 11, 11, 14, -11, 14, -1, -8, -5, -16, -6}
	N2s := []int{12, 9, 8, 3, 0, 11, 10, 13, 3, 10, 10, 14, 6, 5}
	stack := [][2]int{}
	result := [14]int{}
	fmt.Println(len(N1s), len(N2s))

	for i := 0; i < 14; i++ {
		N1 := N1s[i]
		N2 := N2s[i]

		if N1 > 0 {
			stack = append(stack, [2]int{i, N2})
		} else {
			pop := stack[len(stack)-1]
			resultIndex := pop[0]
			N2push := pop[1]
			stack = stack[:len(stack)-1]
			resultPush := 9 - (N2push + N1)
			if resultPush > 9 {
				resultPush = 9
			}
			result[resultIndex] = resultPush
			result[i] = resultPush + N2push + N1
		}
	}
	fmt.Println(result)

	for i := 0; i < 14; i++ {
		N1 := N1s[i]
		N2 := N2s[i]

		if N1 > 0 {
			stack = append(stack, [2]int{i, N2})
		} else {
			pop := stack[len(stack)-1]
			resultIndex := pop[0]
			N2push := pop[1]
			stack = stack[:len(stack)-1]
			resultPush := 1 - (N2push + N1)
			if resultPush < 1 {
				resultPush = 1
			}
			result[resultIndex] = resultPush
			result[i] = resultPush + N2push + N1
		}
	}
	fmt.Println(result)
}
