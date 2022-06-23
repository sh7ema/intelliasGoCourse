package main

import (
	// "reflect"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	inputArray := strings.Fields(input)
	min := 0
	max := 0
	// arr := []int{}
	for _, i := range inputArray {
		s, err := strconv.Atoi(i)
		if err == nil {
			if s > max {
				max = s
			}
			if s < min {
				min = s
			}
		}
	}
	r := fmt.Sprint(max, min)
	fmt.Println(input)
	fmt.Println(r)
}
