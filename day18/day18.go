package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type pair struct {
	left_pair  *pair
	right_pair *pair
	int_value  int
	depth      int
}

func to_pair(lant_inter interface{}, depth int) pair {
	flt, ok := lant_inter.(float64)
	if ok {
		return pair{
			left_pair:  nil,
			right_pair: nil,
			int_value:  int(flt),
			depth:      depth,
		}
	} else {
		arr := lant_inter.([]interface{})
		left := to_pair(arr[0], depth+1)
		right := to_pair(arr[1], depth+1)
		return pair{
			left_pair:  &left,
			right_pair: &right,
			int_value:  -1,
			depth:      depth,
		}
	}
}

func explode(p *pair) bool {
	queue := []*pair{p}
	var last_int *pair = nil
	right_increment := -1
	exploded := false
	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if current.int_value == -1 && current.left_pair.int_value != -1 && current.right_pair.int_value != -1 && current.depth >= 4 && right_increment == -1 {
			//this should explode
			//fmt.Println(current.left_pair.int_value, current.right_pair.int_value)
			if last_int != nil {
				last_int.int_value += current.left_pair.int_value
			}
			right_increment = current.right_pair.int_value
			exploded = true
			current.int_value = 0
			current.left_pair = nil
			current.right_pair = nil
			continue
		}
		if current.int_value != -1 {
			last_int = current
			if exploded {
				current.int_value += right_increment
				return true
			}
		} else {
			queue = append(queue, current.right_pair, current.left_pair)
		}
	}
	return exploded
}

func split(p *pair) bool {
	queue := []*pair{p}
	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if current.int_value >= 10 {
			new_left := pair{
				left_pair:  nil,
				right_pair: nil,
				int_value:  current.int_value / 2,
				depth:      current.depth + 1,
			}

			new_right := pair{
				left_pair:  nil,
				right_pair: nil,
				int_value:  current.int_value - (current.int_value / 2),
				depth:      current.depth + 1,
			}

			current.left_pair = &new_left
			current.right_pair = &new_right
			current.int_value = -1
			return true
		}
		if current.int_value == -1 {
			queue = append(queue, current.right_pair, current.left_pair)
		}
	}
	return false
}

func reduce(p *pair) {
	exploded := true
	splitted := true
	for exploded || splitted {
		exploded = false
		splitted = false
		exploded = explode(p)
		if !exploded {
			splitted = split(p)
		}
	}
}

func pair_string(p *pair) string {
	if p.int_value != -1 {
		return strconv.Itoa(p.int_value)
	} else {
		return "[" + pair_string(p.left_pair) + "," + pair_string(p.right_pair) + "]"
	}
}

func add_pairs(p1 *pair, p2 *pair) *pair {
	new_pair := pair{
		left_pair:  p1,
		right_pair: p2,
		int_value:  -1,
		depth:      0,
	}

	queue := []*pair{p1, p2}
	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		current.depth++
		if current.int_value == -1 {
			queue = append(queue, current.right_pair, current.left_pair)
		}
	}

	return &new_pair
}

func magnitude(p *pair) int {
	if p.int_value != -1 {
		return p.int_value
	} else {
		return magnitude(p.left_pair)*3 + magnitude(p.right_pair)*2
	}
}

func main() {
	data, err := ioutil.ReadFile("day18/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	//fmt.Println(len(lines))
	var current_number *pair
	for i, l := range lines {
		b := []byte(l)
		var f interface{}
		err = json.Unmarshal(b, &f)
		p := to_pair(f, 0)
		if i == 0 {
			current_number = &p
		} else {
			current_number = add_pairs(current_number, &p)
		}
		reduce(current_number)
	}

	fmt.Println(pair_string(current_number))
	fmt.Println(magnitude(current_number))
	//expldd := explode(&p)
	//fmt.Println(expldd)
	max_mag := 0
	for i, l1 := range lines {
		for j, l2 := range lines {
			if i != j {
				b := []byte(l1)
				var f interface{}
				err = json.Unmarshal(b, &f)
				p1 := to_pair(f, 0)
				b = []byte(l2)
				err = json.Unmarshal(b, &f)
				p2 := to_pair(f, 0)

				added := add_pairs(&p1, &p2)
				reduce(added)
				mag := magnitude(added)
				if mag > max_mag {
					max_mag = mag
				}
			}
		}
	}
	fmt.Println(max_mag)
}
