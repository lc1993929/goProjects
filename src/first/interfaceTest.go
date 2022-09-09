package main

import "fmt"

func main() {
	var test1 i1
	var test2 i2
	fmt.Println(test1 == test2)
}

type i1 interface {
}
type i2 interface {
}
