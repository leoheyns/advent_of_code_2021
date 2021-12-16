package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coord struct {
	i int
	j int
}

type unvisitedCoord struct {
	c     coord
	dist  int
	index int
}

type unvisitedQueue []*unvisitedCoord

func (uq unvisitedQueue) Len() int           { return len(uq) }
func (uq unvisitedQueue) Less(i, j int) bool { return uq[i].dist < uq[j].dist }
func (uq unvisitedQueue) Swap(i, j int) {
	uq[i], uq[j] = uq[j], uq[i]
	uq[i].index = i
	uq[j].index = j
}

func (uq *unvisitedQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	index := len(*uq)
	uc := x.(unvisitedCoord)
	uc.index = index
	*uq = append(*uq, &uc)
}

func (uq *unvisitedQueue) Pop() interface{} {
	old := *uq
	n := len(old)
	x := old[n-1]
	*uq = old[0 : n-1]
	return x
}

func neighbours(cave [][]int, c coord) []coord {
	result := make([]coord, 0)
	i := c.i
	j := c.j
	if c.j > 0 {
		result = append(result, coord{i: i, j: j - 1})
	}
	if j+1 < len(cave[0]) {
		result = append(result, coord{i: i, j: j + 1})
	}
	if i+1 < len(cave) {
		result = append(result, coord{i: i + 1, j: j})
	}
	if i > 0 {
		result = append(result, coord{i: i - 1, j: j})
	}
	return result
}
func dijkstra(cave [][]int) int {

	//create vertex set Q

	prev := map[coord]coord{}
	unvisited_contains := map[coord]bool{}
	unvisited := unvisitedQueue{}
	queueMap := map[coord]*unvisitedCoord{}

	//for each vertex v in Graph:
	for i, row := range cave {
		for j := range row {
			//dist[v] ← INFINITY
			//add v to Q
			uc := &unvisitedCoord{coord{i, j}, 1000000, i*len(row) + j}
			unvisited = append(unvisited, uc)
			unvisited_contains[coord{i, j}] = true
			queueMap[coord{i, j}] = uc
		}
	}
	//dist[source] ← 0
	queueMap[coord{0, 0}].dist = 0
	heap.Init(&unvisited)
	//while Q is not empty:
	for len(unvisited) > 0 {
		//u ← vertex in Q with min dist[u]
		u := heap.Pop(&unvisited).(*unvisitedCoord)
		//remove u from Q
		unvisited_contains[u.c] = false
		//for each neighbor v of u still in Q:
		for _, v := range neighbours(cave, u.c) {
			//if still in unvisited
			if unvisited_contains[v] {
				alt := u.dist + cave[v.i][v.j]
				//if alt < dist[v]:
				if alt < queueMap[v].dist {
					//dist[v] ← alt
					queueMap[v].dist = alt
					heap.Fix(&unvisited, queueMap[v].index)
					//prev[v] ← u
					prev[v] = u.c
				}
			}
		}
	}
	//22      return dist[], prev[]
	//	fmt.Println(dist)
	return queueMap[coord{len(cave) - 1, len(cave[0]) - 1}].dist
}

func main() {
	data, err := ioutil.ReadFile("day15/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	cave := make([][]int, len(lines))
	for i, l := range lines {
		cave[i] = make([]int, len(l))
		for j, chr := range l {
			risk, _ := strconv.Atoi(string(chr))
			cave[i][j] = risk
		}
	}
	fmt.Println(dijkstra(cave))
	//cave = [][]int{{1,2,3},{4,5,6},{7,8,9}}
	large_cave := make([][]int, len(cave)*5)
	for i := range large_cave {
		large_cave[i] = make([]int, len(cave[0])*5)
	}
	for k := 0; k < 5; k++ {
		for l := 0; l < 5; l++ {
			for i, row := range cave {
				for j, small_risk := range row {
					large_cave[k*len(cave)+i][l*len(row)+j] = (small_risk+k+l-1)%9 + 1
				}
			}
		}
	}
	//fmt.Println(large_cave)
	fmt.Println(dijkstra(large_cave))
}
