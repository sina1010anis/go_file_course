package main

import (
	"fmt"
)

// func setAddress(x *int) {

// 	*x = 25

// }

// type DB struct {
// 	w, h float64
// }

// type SET interface {
// 	sum() float64
// }

// func add(d DB) float64 {

// 	return d.w + d.h*d.sum()

// }

// func (d) sum() float64 {

// 	return 3.17 - 0.5/14

// }

// type SW struct {
// 	w, h float64
// }

// type GET interface {
// 	sumGET() float64
// }

// func (d DB) sumGET() float64 {

// 	return 8000.0001

// }

// type User struct {
// 	name   string
// 	family string
// 	age    int
// }

// func (u User) showUser() {

// 	fmt.Println("User : (name)->", u.name)
// 	fmt.Println("User : (name)->", u.family)
// 	fmt.Println("User : (name)->", u.age)

// }

// func hello(ch chan int) {

// 	time.Sleep(time.Second * 2)

// 	ch <- 50

// }

func main() {

	ch := make(chan bool, 1)

	ch <- true

	select {
	case <-ch:
		fmt.Println("OK")
		// default:
		// 	fmt.Println("END")
	}

	// go func() {
	// 	val := 45

	// 	ch <- val
	// }()

	// fmt.Println(<-ch)

	// u := User{name: "sina", family: "nz", age: 45}

	// u.showUser()

	// return fmt.Errorf("Opps")

	// r := DB{w: 25, h: 45}

	// w := SW{w: 25, h: 45}

	// fmt.Println(add(r))

	// fmt.Println(add(W))

	// m := make(map[int]string)

	// m[1] = "Item_A"

	// m[2] = "Item_B"

	// n := map[int]string{1: "Item_A", 2: "Item_B"}

	// val, has := m[1]

	// fmt.Println(val)

	// fmt.Println(has)

	// fmt.Println(maps.Equal(m, n))

	// users := make(map[int]string, 5)

	// users[0] = "Hi...!"

	// fmt.Println(users)

	// arr := [5]int{1, 2, 3, 4, 5}

	// arr2 := [...]int{1, 2, 3, 4, 5}

	// arr3 := []int{1, 2, 3, 4, 5}

	// fmt.Printf("%T\n", arr)

	// fmt.Printf("%T\n", arr2)

	// fmt.Printf("%T\n", arr3)

	// a := 10

	// b := &a

	// fmt.Println(*b)

	// setAddress(b)

	// fmt.Println(*b)

	// numbers := []int{1, 2, 3}

	// numbers_2 := []int{4, 5, 6}

	// new_numbers := append(numbers, numbers_2...)

	// fmt.Println(new_numbers)

	// number := [5]int{1, 2, 3, 4, 5}

	// new_number := number[1:3]

	// fmt.Println(new_number)

	// userIDs := [3]int{1, 3, 6}

	// for _, i := range userIDs {

	// 	fmt.Println(i)

	// }

	// x := 5

	// p := &x

	// fmt.Println(*p)

	// *p = 45

	// fmt.Println(*p)

	// var x int32 = 100

	// fmt.Println(x)

	// fmt.Println(add(45, 50))

	// panic("Oppe")

	// fmt.Println(recover())

	// f, err := os.Open("./file.txt")

	// if err != nil {

	// 	fmt.Println("Error: ", err)

	// 	return

	// }

	// fmt.Println("File: ", f.Close())

	// fmt.Println("Val1 public")

	// defer fmt.Println("Val2 public")

	// fmt.Println("Val3 public")

	// 	fmt.Println("Hello ...")

	// 	goto mms

	// 	fmt.Println("Hello 2 ...")

	// mms:

	// 	fmt.Println("Hello 3 ...")

	// var name string = "sina"

	// var userIDs = []int{} // Slices

	// var userIDs = [5]int{1, 2, 3, 4, 5} // Array

	// var userIDs = map[string]int{"x": 1, "y": 2} // Map

	// fmt.Println(userIDs["sina"])

	// version := 13

	// switch version {

	// case 10:
	// 	fmt.Println("Oh No you version 10")
	// case 11:
	// 	fmt.Println("Oh No you version 11")
	// case 12:
	// 	fmt.Println("Oh No you version 12")
	// default:
	// 	fmt.Println("Opss not a veriosn")
	// }

	// for key, val := range userIDs {

	// 	fmt.Println(key, val)

	// }

}

// func val1() {

// 	fmt.Println("Val1")

// }
// func val2() {

// 	fmt.Println("Val2")

// }
// func val3() {

// 	fmt.Println("Val3")

// }
