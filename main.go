package main

import (
	"fmt"
	"product/format"
)

type my_type = float64

const Num my_type = 45.5

const (
	a int = iota

	b

	c
)

func main() {

	i := new(42)

	fmt.Println("i => ", i)

	fmt.Println()

	var user_id string = "50"

	var temp *string = &user_id

	fmt.Println("*temp => ", *temp)

	user_id = "88"

	fmt.Println("*temp (update) => ", *temp)

	format.TypeValue(a)

	fmt.Println(a, b, c)

	var num_1 my_type = Num

	var num_2 float64 = Num

	fmt.Printf("%v, %v\n", num_1, num_2)

}
