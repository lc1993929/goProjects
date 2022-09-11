package main

import "fmt"

// 注释
func main() {
	//fmt.Print("test")
	//var arr [10]int
	//arr[1] = 1
	//fmt.Print(arr)

	//slice := []int{}
	//fmt.Println(len(slice), cap(slice))
	//slice = append(slice, 1)
	//fmt.Println(len(slice), cap(slice))
	//slice = append(slice, 1)
	//fmt.Println(len(slice), cap(slice))
	//slice = append(slice, 1)
	//fmt.Println(len(slice), cap(slice))
	//slice = append(slice, 1)
	//fmt.Println(len(slice), cap(slice))
	//fmt.Println(slice)

	test := Test{name: "123"}
	fmt.Println(test == Test{})

}

type Test struct {
	name string
}
