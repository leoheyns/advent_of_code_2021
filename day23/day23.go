package main

import (
	"container/heap"
	"fmt"
	"reflect"
)

type pqIndex struct {
	i        int
	priority int
	index    int
}

type priorityQueue []*pqIndex

func (uq priorityQueue) Len() int           { return len(uq) }
func (uq priorityQueue) Less(i, j int) bool { return uq[i].priority < uq[j].priority }
func (uq priorityQueue) Swap(i, j int) {
	uq[i], uq[j] = uq[j], uq[i]
	uq[i].index = i
	uq[j].index = j
}

func (uq *priorityQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	index := len(*uq)
	uc := x.(*pqIndex)
	uc.index = index
	*uq = append(*uq, uc)
}

func (uq *priorityQueue) Pop() interface{} {
	old := *uq
	n := len(old)
	x := old[n-1]
	*uq = old[0 : n-1]
	return x
}

type aStarState interface {
	neighbours() ([]aStarState, map[aStarState]int)
	equals(state aStarState) bool
}

func reconstructPath(cameFrom map[aStarState]aStarState, current aStarState) []aStarState {
	totalPath := []aStarState{current}
	for true {
		current = cameFrom[current]
		totalPath = append([]aStarState{current}, totalPath...)
		_, ok := cameFrom[current]
		if !ok {
			break
		}
	}
	return totalPath
}

func aStar(start aStarState, goal aStarState, heuristic func(state aStarState) int) (int, map[aStarState]aStarState) {
	var stateArray []aStarState
	pqMap := map[aStarState]*pqIndex{}
	openSet := priorityQueue{}
	stateArray = append(stateArray, start)

	cameFrom := map[aStarState]aStarState{}

	startPqIndex := pqIndex{
		i:        0,
		priority: 0,
		index:    0,
	}
	pqMap[start] = &startPqIndex
	openSet.Push(&startPqIndex)
	//cameFrom := map[aStarState]aStarState{}

	gScore := map[aStarState]int{}
	fScore := map[aStarState]int{}

	gScore[start] = 0
	fScore[start] = heuristic(start)

	for len(openSet) > 0 {
		currentIndex := heap.Pop(&openSet).(*pqIndex)
		currentState := stateArray[(*currentIndex).i]
		delete(pqMap, currentState)
		if currentState.equals(goal) {
			return gScore[currentState], cameFrom
		}
		neighbours, distances := currentState.neighbours()
		for _, neighbour := range neighbours {
			tentativeGScore := gScore[currentState] + distances[neighbour]
			nscore, ok := gScore[neighbour]

			if !ok || tentativeGScore < nscore {
				cameFrom[neighbour] = currentState
				gScore[neighbour] = tentativeGScore
				fScore[neighbour] = tentativeGScore + heuristic(neighbour)
				index, ok := pqMap[neighbour]
				if !ok {
					newPqIndex := pqIndex{
						i:        len(stateArray),
						priority: fScore[neighbour],
						index:    len(openSet),
					}
					pqMap[neighbour] = &newPqIndex
					stateArray = append(stateArray, neighbour)

					heap.Push(&openSet, pqMap[neighbour])
				} else {
					index.priority = fScore[neighbour]
					heap.Fix(&openSet, index.index)
				}
			}
		}
	}
	return -1, cameFrom
}

type amphipodState struct {
	siderooms [4][4]string
	hallway   [11]string
}

func (s1 amphipodState) equals(s2interface aStarState) bool {
	s2, ok := s2interface.(amphipodState)
	if !ok {
		return false
	}
	return reflect.DeepEqual(s1, s2)
}

func (s amphipodState) copy() amphipodState {
	result := amphipodState{
		siderooms: [4][4]string{},
		hallway:   [11]string{},
	}

	for i, sideroom := range s.siderooms {
		for j, cell := range sideroom {
			result.siderooms[i][j] = cell
		}
	}
	for i, cell := range s.hallway {
		result.hallway[i] = cell
	}
	return result
}

func pathFree(hallway [11]string, start int, end int) bool {
	free := true
	if start-end < 0 {
		for i := start + 1; i <= end; i++ {
			if hallway[i] != "." {
				free = false
			}
		}
	} else {
		for i := start - 1; i >= end; i-- {
			if hallway[i] != "." {
				free = false
			}
		}
	}
	return free
}

func (s amphipodState) neighbours() ([]aStarState, map[aStarState]int) {
	possibleStates := []aStarState{}
	distances := map[aStarState]int{}

	roomOpen := make([]bool, 4)
	sideroomChar := map[int]string{0: "A", 1: "B", 2: "C", 3: "D"}
	targets := map[string]int{"A": 0, "B": 1, "C": 2, "D": 3}
	costs := map[string]int{"A": 1, "B": 10, "C": 100, "D": 1000}
	hallwayConnections := []int{2, 4, 6, 8}

	for i, room := range s.siderooms {
		roomOpen[i] = true
		for _, cell := range room {
			if !(cell == "." || cell == sideroomChar[i]) {
				roomOpen[i] = false
			}
		}
	}

	for i, room := range s.siderooms {
		topAmp := "."
		topAmpJ := 0
		for j, cell := range room {
			if cell != "." {
				topAmp = cell
				topAmpJ = j
				break
			}
		}
		if i == targets[topAmp] && roomOpen[i] {
			continue
		}
		if roomOpen[targets[topAmp]] && pathFree(s.hallway, hallwayConnections[i], hallwayConnections[targets[topAmp]]) {
			//move straight from sideroom to correct sideroom
			distance := topAmpJ + 1
			if hallwayConnections[i] < hallwayConnections[targets[topAmp]] {
				distance += hallwayConnections[targets[topAmp]] - hallwayConnections[i]
			} else {
				distance += hallwayConnections[i] - hallwayConnections[targets[topAmp]]
			}

			newState := s.copy()
			newState.siderooms[i][topAmpJ] = "."
			for j := 0; j < len(s.siderooms[0]); j++ {
				distance++
				if j == len(s.siderooms[0])-1 {
					newState.siderooms[targets[topAmp]][j] = topAmp
					break
				}
				if s.siderooms[targets[topAmp]][j+1] != "." {
					newState.siderooms[targets[topAmp]][j] = topAmp
					break
				}
			}
			possibleStates = append(possibleStates, newState)
			distances[newState] = distance * costs[topAmp]
		} else {
			//move from sideroom to hallway
			for _, j := range []int{0, 1, 3, 5, 7, 9, 10} {
				if pathFree(s.hallway, hallwayConnections[i], j) {
					distance := topAmpJ + 1
					if hallwayConnections[i] < j {
						distance += j - hallwayConnections[i]
					} else {
						distance += hallwayConnections[i] - j
					}

					newState := s.copy()
					newState.siderooms[i][topAmpJ] = "."

					newState.hallway[j] = topAmp
					possibleStates = append(possibleStates, newState)
					distances[newState] = distance * costs[topAmp]
				}
			}
		}
	}
	//move from hallway to correct sideroom
	for _, i := range []int{0, 1, 3, 5, 7, 9, 10} {
		cell := s.hallway[i]
		if cell != "." {
			if roomOpen[targets[cell]] && pathFree(s.hallway, i, hallwayConnections[targets[cell]]) {
				distance := 0
				if i < hallwayConnections[targets[cell]] {
					distance += hallwayConnections[targets[cell]] - i
				} else {
					distance += i - hallwayConnections[targets[cell]]
				}

				newState := s.copy()
				newState.hallway[i] = "."
				for j := 0; j < len(s.siderooms[0]); j++ {
					distance++
					if j == len(s.siderooms[0])-1 {
						newState.siderooms[targets[cell]][j] = cell
						break
					}
					if s.siderooms[targets[cell]][j+1] != "." {
						newState.siderooms[targets[cell]][j] = cell
						break
					}
				}
				possibleStates = append(possibleStates, newState)
				distances[newState] = distance * costs[cell]
			}
		}
	}
	return possibleStates, distances
}

//#############
//#...........#
//###A#D#B#C###
//  #D#C#B#A#
//  #D#B#A#C#
//  #B#C#D#A#
//  #########
func heur(state aStarState) int {
	s := state.(amphipodState)
	//sideroomChar := map[int]string{0:"A", 1:"B", 2:"C", 3:"D"}
	targets := map[string]int{"A": 0, "B": 1, "C": 2, "D": 3}
	costs := map[string]int{"A": 1, "B": 10, "C": 100, "D": 1000}
	hallwayConnections := []int{2, 4, 6, 8}

	totalCost := 0
	for i, room := range s.siderooms {
		topAmp := "."
		topAmpJ := 0
		for j, cell := range room {
			if cell != "." {
				topAmp = cell
				topAmpJ = j
				break
			}
		}
		//move straight from sideroom to correct sideroom
		distance := topAmpJ + 1
		if hallwayConnections[i] < hallwayConnections[targets[topAmp]] {
			distance += hallwayConnections[targets[topAmp]] - hallwayConnections[i]
		} else {
			distance += hallwayConnections[i] - hallwayConnections[targets[topAmp]]
		}

		for j := 0; j < len(s.siderooms[0]); j++ {
			distance++
			if j == len(s.siderooms[0])-1 {
				break
			}
			if s.siderooms[targets[topAmp]][j+1] != "." {
				break
			}
		}
		totalCost += distance * costs[topAmp]
	}

	//move from hallway to correct sideroom
	for _, i := range []int{0, 1, 3, 5, 7, 9, 10} {
		cell := s.hallway[i]
		if cell != "." {
			distance := 0
			if i < hallwayConnections[targets[cell]] {
				distance += hallwayConnections[targets[cell]] - i
			} else {
				distance += i - hallwayConnections[targets[cell]]
			}

			newState := s.copy()
			newState.hallway[i] = "."
			for j := 0; j < len(s.siderooms[0]); j++ {
				distance++
				if j == len(s.siderooms[0])-1 {
					newState.siderooms[targets[cell]][j] = cell
					break
				}
				if s.siderooms[targets[cell]][j+1] != "." {
					newState.siderooms[targets[cell]][j] = cell
					break
				}
			}
			totalCost = distance * costs[cell]
		}
	}

	return totalCost
}

func printState(s amphipodState) {
	fmt.Println()
	fmt.Println("#############")
	print("#")
	for i := 0; i < 11; i++ {
		print(s.hallway[i])
	}
	print("#\n")
	for i := range []int{0, 1, 2, 3} {
		if i == 0 {
			print("###")
		} else {
			print("  #")
		}
		for j := range []int{0, 1, 2, 3} {
			print(s.siderooms[j][i])
			print("#")
		}
		if i == 0 {
			print("##\n")
		} else {
			print("\n")
		}
	}
	fmt.Println("  #########")
	fmt.Println()
}

func main() {
	start := amphipodState{
		siderooms: [4][4]string{{"A", "D", "D", "B"}, {"D", "C", "B", "C"}, {"B", "B", "A", "D"}, {"C", "A", "C", "A"}},
		hallway:   [11]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}

	goal := amphipodState{
		siderooms: [4][4]string{{"A", "A", "A", "A"}, {"B", "B", "B", "B"}, {"C", "C", "C", "C"}, {"D", "D", "D", "D"}},
		hallway:   [11]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}
	//printState(start)
	dist, camefrom := aStar(start, goal, heur)
	path := reconstructPath(camefrom, goal)
	fmt.Println(dist)
	//fmt.Println(len(path))
	for _, state := range path {
		fmt.Println("------------------------")
		printState(state.(amphipodState))
		fmt.Println()
	}
}
