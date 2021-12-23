package main

type priorityQueueObject interface {
	getObject() interface{}
	getPriority() int
	setPriority(i int)
	setIndex(i int)
}

type priorityQueue []*priorityQueueObject

func (pQ priorityQueue) Len() int           { return len(pQ) }
func (pQ priorityQueue) Less(i, j int) bool { return (*pQ[i]).getPriority() < (*pQ[j]).getPriority() }
func (pQ priorityQueue) Swap(i, j int) {
	pQ[i], pQ[j] = pQ[j], pQ[i]
	(*pQ[i]).setIndex(i)
	(*pQ[j]).setIndex(j)
}

func (pQ *priorityQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	index := len(*pQ)
	uc := x.(priorityQueueObject)
	uc.setIndex(index)
	*pQ = append(*pQ, &uc)
}

func (pQ *priorityQueue) Pop() interface{} {
	old := *pQ
	n := len(old)
	x := old[n-1]
	*pQ = old[0 : n-1]
	return x
}

type aStarState interface {
	neighbours() []aStarState
}
