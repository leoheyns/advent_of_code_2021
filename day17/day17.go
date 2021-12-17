package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func calc_stopping_xs(x_min int, x_max int) []int {
	result := []int{}
	if x_min < 0 {
		for v_x := 0; v_x > x_min; v_x-- {
			if -(v_x*(v_x+1))/2 >= x_min && -(v_x*(v_x+1))/2 <= x_max {
				result = append(result, v_x)
			}
		}
	} else {
		for v_x := 0; v_x < x_max; v_x++ {
			if v_x*(v_x+1)/2 >= x_min && v_x*(v_x+1)/2 <= x_max {
				result = append(result, v_x)
			}
		}
	}
	return result
}

func calc_valid_x_steps(x_min int, x_max int) map[int][]int {
	result := map[int][]int{}
	for v_x := 1; v_x <= x_max; v_x++ {
		pos := 0
		vel := v_x
		steps := 0
		for pos <= x_max && vel >= 0 {
			steps++
			pos += vel
			vel--
			if pos >= x_min && pos <= x_max {
				_, ok := result[steps]
				if !ok {
					result[steps] = make([]int, 0)
				}
				result[steps] = append(result[steps], v_x)
			}
		}
	}
	return result
}

type velocity_2d struct {
	x_vel int
	y_vel int
}

func calc_possible_ys(y_min int, y_max int, valid_steps map[int][]int, always_works []int) ([]int, map[velocity_2d]bool) {
	result := []int{}
	resultMap := map[velocity_2d]bool{}
	for v_y := y_min; v_y < -y_min; v_y++ {
		pos := 0
		vel := v_y
		steps := 0
		for pos > y_min {
			pos += vel
			vel--
			steps++
			if pos >= y_min && pos <= y_max {
				result = append(result, -v_y)
				xs, ok := valid_steps[steps]
				if ok {
					for _, x := range xs {
						resultMap[velocity_2d{
							x_vel: x,
							y_vel: -v_y,
						}] = true
					}
				}

				for _, x := range always_works {
					if steps > x {
						resultMap[velocity_2d{
							x_vel: x,
							y_vel: -v_y,
						}] = true
					}
				}
			}
		}
	}
	return result, resultMap
}

//wtf hoe moet dit
func main() {
	data, err := ioutil.ReadFile("day17/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	parts := strings.Split(string(data), " ")
	x_range_string, y_range_string := parts[2][2:len(parts[2])-1], parts[3][2:]
	x_min, _ := strconv.Atoi(strings.Split(x_range_string, "..")[0])
	x_max, _ := strconv.Atoi(strings.Split(x_range_string, "..")[1])
	y_min, _ := strconv.Atoi(strings.Split(y_range_string, "..")[0])
	y_max, _ := strconv.Atoi(strings.Split(y_range_string, "..")[1])
	always_works := calc_stopping_xs(x_min, x_max)
	valid_steps := calc_valid_x_steps(x_min, x_max)

	valid_ys, valid_velocities := calc_possible_ys(y_min, y_max, valid_steps, always_works)
	fmt.Println(valid_ys[len(valid_ys)-1] * (valid_ys[len(valid_ys)-1] + 1) / 2)
	//fmt.Println(valid_velocities)
	//max_y_vel := valid_ys[len(valid_ys) - 1]
	//data2, err2 := ioutil.ReadFile("day17/example_result")
	//if err2 != nil {
	//	fmt.Println("File reading error", err)
	//	return
	//}
	//re := regexp.MustCompile("[ \n]+")
	//coord_strings := re.Split(string(data2), -1)
	//correct_results := map[velocity_2d]bool{}
	//for _,coord_string := range coord_strings{
	//	x,_ := strconv.Atoi(strings.Split(coord_string, ",")[0])
	//	y,_ := strconv.Atoi(strings.Split(coord_string, ",")[1])
	//	_,ok := valid_velocities[velocity_2d{x,y}]
	//	correct_results[velocity_2d{x,y}] = true
	//	if !ok{
	//		fmt.Println(x,y,ok)
	//	}
	//}

	//for vel := range valid_velocities{
	//	_,ok := correct_results[vel]
	//	if !ok{
	//		fmt.Println(vel)
	//	}
	//}
	fmt.Println(len(valid_velocities))
	//fmt.Println(valid_ys)
	//fmt.Println(always_works)
	//fmt.Println(valid_steps)
}
