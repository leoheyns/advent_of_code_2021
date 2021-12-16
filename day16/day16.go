package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var hex_to_bin_string = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

type packet struct {
	version  int
	tipe     int
	value    int
	children []*packet
	length   int
}

func bits_to_int(bitstring string) int {
	total := 0
	for i, r := range bitstring {
		if string(r) == "1" {
			total += 1 << (len(bitstring) - i - 1)
		}
	}
	return total
}

func read_literal(bitstring string) (int, int) {
	literal_bitstring := ""
	length := 0
	for i := 0; true; i++ {
		literal_bitstring += bitstring[i*5+1 : i*5+5]
		length += 5
		if bitstring[i*5] == '0' {
			break
		}
	}
	return bits_to_int(literal_bitstring), length
}

func read_packet(data string) packet {
	this := packet{
		version:  0,
		tipe:     0,
		value:    0,
		children: []*packet{},
		length:   0,
	}
	this.version = bits_to_int(data[:3])
	this.tipe = bits_to_int(data[3:6])
	if this.tipe == 4 {
		lit_value, lit_length := read_literal(data[6:])
		this.value = lit_value
		this.length = lit_length + 6
	} else {
		length_type := data[6] == '1'
		if length_type {
			child_count := bits_to_int(data[7:18])
			data_index := 18
			for i := 0; i < child_count; i++ {
				child := read_packet(data[data_index:])
				data_index += child.length
				this.children = append(this.children, &child)
			}
			this.length = data_index
		} else {
			this.length = 22 + bits_to_int(data[7:22])
			data_index := 22
			for data_index < this.length {
				child := read_packet(data[data_index:])
				data_index += child.length
				this.children = append(this.children, &child)
			}
		}
	}
	return this
}

func version_sum(pack packet) int {
	result := pack.version
	for _, child := range pack.children {
		result += version_sum(*child)
	}
	return result
}

func value_of(pack packet) int {
	switch pack.tipe {
	case 0:
		sum := 0
		for _, child := range pack.children {
			sum += value_of(*child)
		}
		return sum
	case 1:
		prod := 1
		for _, child := range pack.children {
			prod *= value_of(*child)
		}
		return prod
	case 2:
		//minimum
		min := -1
		for _, child := range pack.children {
			v := value_of(*child)
			if min == -1 || v < min {
				min = v
			}
		}
		return min
	case 3:
		//maximum
		max := 0
		for _, child := range pack.children {
			v := value_of(*child)
			if v > max {
				max = v
			}
		}
		return max
	case 4:
		//literal
		return pack.value
	case 5:
		//greater than
		if value_of(*pack.children[0]) > value_of(*pack.children[1]) {
			return 1
		} else {
			return 0
		}
	case 6:
		//less than
		if value_of(*pack.children[0]) < value_of(*pack.children[1]) {
			return 1
		} else {
			return 0
		}
	case 7:
		//equal to
		if value_of(*pack.children[0]) == value_of(*pack.children[1]) {
			return 1
		} else {
			return 0
		}
	}
	return 0
}

func main() {
	data, err := ioutil.ReadFile("day16/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		packet_hex := l
		packet_bin := ""
		for _, char := range packet_hex {
			packet_bin += hex_to_bin_string[char]
		}
		pack := read_packet(packet_bin)
		fmt.Println(version_sum(pack))
		fmt.Println(value_of(pack))
	}
}
