package main

import "fmt"

type iets struct {
	value string
}

func main() {
	x := iets{"hallo"}
	fmt.Println(x.value)
}
