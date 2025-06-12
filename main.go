package main

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

// func hellpHttp(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprintln(w, "Hello")

// }

func Spl(str string) (dir, path string) {

	return "ABC", "DFG"

}

func main() {

	// ----------------- for range (start)
	// var (
	// 	i int
	// 	v int
	// )

	// l_x := []int{0, 1, 2, 3, 4, 5}

	// for i, v = range l_x {

	// 	fmt.Printf("#%d: %d\n", i+1, v)

	// }

	// fmt.Println(i, v)
	// ----------------- for range (end)
	// ----------------- min project (start)
	// l_x := []int{0, 1, 2, 3, 4, 5}

	// l_y := []int{0, 1, 2, 3, 4, 5}

	// for i := 0; i < len(l_x); i++ {

	// 	for j := 0; j < len(l_y); j++ {

	// 		fmt.Printf("%d * %d = %d\n", l_x[i], l_y[j], l_x[i]*l_y[j])

	// 	}

	// }
	// ----------------- min project (end)
	// ----------------- basic for (statrt)
	// for i := 1; i <= 50; i++ {

	// 	time.Sleep(time.Second / 2)
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second / 2)

	// }
	// ----------------- basic for (end)
	// -------------------- project time day (start)
	// switch hour := time.Now().Hour(); true {

	// case hour >= 6 && hour < 12:
	// 	fmt.Println("Sobhe")

	// case hour >= 12 && hour < 18:
	// 	fmt.Println("Asre")

	// case hour >= 18 && hour <= 24:
	// 	fmt.Println("Shab")

	// case hour < 6 && hour > 0:
	// 	fmt.Println("Bamdad")
	// default:
	// 	fmt.Println("Error")
	// }
	// -------------------- project time day (end)
	// -------------------- swith case (start)
	// name := "ahmad"
	// switch name {
	// case "sina", "sin":
	// 	fmt.Println("Hi sina")
	// case "ahmad":
	// 	fmt.Println("Hi ahmad")
	// }
	// -------------------- swith case (end)
	// -------------------- Value in if (start)
	// if _, err := strconv.Atoi("85"); err != nil {

	// 	fmt.Println("Error")

	// 	return

	// }
	// -------------------- Value in if (end)
	// --------------- project (statrt)
	// if len(os.Args) != 2 {

	// 	fmt.Println("Type a Vlaue")

	// 	return

	// }

	// num, err := strconv.Atoi(os.Args[1])

	// if err != nil {

	// 	fmt.Println("Error: ", err)

	// 	return

	// }

	// fote := float64(num) * 0.3048

	// fmt.Println("Number(Meter)=> ", num, "m | Number(Fote)", fote, "f")
	// --------------- project (end)
	// --------------- strconv (statrt)

	// n, err := strconv.Atoi(os.Args[1])

	// fmt.Println("Error: ", err, " Data: ", n)

	// --------------- strconv (end)

	// --------------- project (statrt)
	// if len(os.Args) != 3 {

	// 	panic("type a username and password => format 'go run . [username] [password]'")

	// } else {

	// 	username, password := strings.ToLower(os.Args[1]), strings.ToLower(os.Args[2])

	// 	def_username, def_password := "jack", "1888"

	// 	if username != def_username || password != def_password {

	// 		panic("Invalide username or password")

	// 	} else {

	// 		fmt.Println("Done Login username: ", username)

	// 	}
	// }
	// --------------- project (end)

	// ------------------ IOTA && get args in trminal for first (start)
	// const (
	// 	a = iota + 1
	// 	b
	// 	c
	// )

	// fmt.Println(a, b, c)

	// val := strings.ToUpper(os.Args[1])

	// cap := strings.Repeat("!", utf8.RuneCountInString(val))

	// fmt.Printf("%s%s%s\n", cap, val, cap)

	// ------------------ IOTA && get args in trminal for first (end)

	// ------------ Get APIs (start)
	// url := "https://digiaparteman.ir/api/test/up"

	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Fatalf("خطا در ارسال درخواست: %v", err)
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("خطا در خواندن پاسخ: %v", err)
	// }

	// fmt.Println("پاسخ دریافت‌شده:")
	// fmt.Println(string(body))
	// ------------ Get APIs (end)

	// ------------ Make webserver (strat)
	// http.HandleFunc("/root", hellpHttp)

	// fmt.Println("Run Server 8010...")

	// http.ListenAndServe(":8010", nil)
	// ------------ Make webserver (end)

	// ------------ Get input in user (start)
	// var x string

	// y := []string{"S", "i", "n", "a"}

	// fmt.Print("Type a number: ")
	// fmt.Scanln(&x)

	// fmt.Printf("Number select: %s\n", x)

	// m := strings.Join(y, "")

	// fmt.Println(m)
	// ------------ Get input in user (end)

	// ------------ channel whit select (start)
	// ch := make(chan bool, 1)

	// ch <- true

	// select {
	// case <-ch:
	// 	fmt.Println("OK")
	// 	default:
	// 	fmt.Println("END")
	// }

	// go func() {
	// 	val := 45

	// 	ch <- val
	// }()

	// fmt.Println(<-ch)
	// ------------ channel whit select (end)

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
