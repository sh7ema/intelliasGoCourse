package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	// var result []int
	keys := make(map[int]bool)
	newArr := []int{}
	for _, entry := range arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			newArr = append(newArr, entry)
		}
	}
	fmt.Println(arr)
	fmt.Println(newArr)

}
