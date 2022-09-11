package main

import "fmt"

func main() {
	m := make(map[int][]int)

	ints := []int{1, 2, 3}
	m[1] = ints

	fmt.Println(m)

	ints = append(ints, 4, 5, 6)

	fmt.Println(m)
}
