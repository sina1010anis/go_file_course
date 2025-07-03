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

// func Spl(str string) (dir, path string) {

// 	return "ABC", "DFG"

// }

// type ver struct {
// 	X, Y int
// }

// func createPassword(len_password int) string {

// 	alf := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNMM1234567890!@#$%^&*()_-=+"

// 	var res string

// 	for i := 0; i < len_password; i++ {

// 		rand_number := rand.Intn(73)

// 		res += string(alf[rand_number])

// 	}

// 	return res

// }

// func maskSubstring(s string, start, end int) string {
// 	if start < 0 || end > len(s) || start >= end {
// 		return s // invalid range, return original
// 	}

// 	masked := strings.Repeat("*", end-start)
// 	return s[:start] + masked + s[end:]
// }

// type User struct {
// 	name   string
// 	family string
// 	age    int
// }

func main() {

	// ---------------------- bubble sort (start)

	// vars := os.Args[1:] // smapel input = 5 4 3 2 1 => outeput(typesort"true") 1 2 3 4 5 => output(type"false") 5 4 3 2 1

	// var (
	// 	m int

	// 	tvars []int

	// 	tnumber int
	// )

	// ssort := false

	// typesort := false

	// if len(vars) <= 0 {

	// 	fmt.Println("⚠️ Type a numbers...!")

	// 	return

	// }

	// for _, v := range vars {

	// 	val, err := strconv.Atoi(v)

	// 	if err != nil {

	// 		fmt.Printf("⚠️ Opss Error a get number (%s)\n", err)

	// 		return

	// 	}

	// 	if val < 0 {

	// 		fmt.Printf("⚠️ a Number Error (%d) tyoe a number positive\n", val)

	// 		return

	// 	}

	// 	tvars = append(tvars, val)

	// }

	// for {

	// 	ssort = false

	// 	for i := 0; i < len(tvars); i++ {

	// 		if i+1 >= len(tvars) {

	// 			m = i

	// 		} else {

	// 			m = i + 1

	// 		}

	// 		if typesort {

	// 			if tvars[i] > tvars[m] {

	// 				tnumber = tvars[i]

	// 				tvars[i] = tvars[m]

	// 				tvars[m] = tnumber

	// 				ssort = true

	// 			}

	// 		} else {

	// 			if tvars[i] < tvars[m] {

	// 				tnumber = tvars[i]

	// 				tvars[i] = tvars[m]

	// 				tvars[m] = tnumber

	// 				ssort = true

	// 			}
	// 		}

	// 	}

	// 	if !ssort {

	// 		fmt.Println(tvars)

	// 		return

	// 	}
	// }
	// ---------------------- bubble sort (end)
	// ---------------------- struct local var in main (start)
	// type Admin struct {
	// 	name, lastname string
	// 	age            int
	// }

	// var user_admin Admin

	// user_admin.name = "root"

	// user_admin.lastname = "local"

	// user_admin.age = 45

	// fmt.Println(user_admin)

	// user_1 := User{"sina", "nz", 26}

	// fmt.Println(user_1)
	// ---------------------- struct local var in main (end)
	// ----------------------- get inpute as trminal for bufio.NewScanner => in.Scan (start)
	// in := bufio.NewScanner(os.Stdin)

	// fmt.Println("Type your name: ")

	// in.Split(bufio.ScanWords)

	// for in.Scan() {

	// 	fmt.Println(in.Text())

	// }

	// if in.Err() != nil {

	// 	fmt.Println("-----------------------")

	// 	fmt.Println("Has a error")

	// 	return

	// }

	// fmt.Println("-----------------------")

	// fmt.Println("Hello: ", in.Text())
	// ----------------------- get inpute as trminal for bufio.NewScanner => in.Scan (end)
	// ------------ map section 2 (start)
	// var kv map[string]string

	// kv = make(map[string]string, 0)

	// fmt.Println(kv)
	// ------------ map section 2 (end)

	// -------------------- proejct maske (start)+

	// str := "//:www.test.irhi ////:www.test.ir sina //:www.test.ir"

	// match := "//"

	// keys := []int{}

	// var s int

	// str = str + " "

	// if len(str) < len(match) {

	// 	fmt.Println("Errro: 800 (len error)")

	// 	return

	// }

	// for i := 0; i <= len(str)-len(match); i++ {

	// 	search := str[i : len(match)+i]

	// 	if search == match {

	// 		m := i

	// 		for j := m; j < len(str); j++ {

	// 			if str[j] == 32 {

	// 				keys = append(keys, m, j)

	// 				break

	// 			}

	// 		}

	// 	}

	// }

	// for s < len(keys) {

	// 	str = maskSubstring(str, keys[s], keys[s+1])

	// 	s += 2

	// }

	// fmt.Println(str)

	// -------------------- proejct maske (end)

	// str := "Hello Wolrd...!"

	// std := []byte(str)

	// fmt.Println(std)

	// start, stop := 'A', 'Z'

	// fmt.Println(start, stop)

	// ---------------- use file for ioutil (start)
	// byte_data := make([]byte, 50)

	// err := ioutil.WriteFile("new.txt", byte_data, 0654)

	// if err != nil {

	// 	fmt.Println(err)

	// }

	// file, err_2 := ioutil.ReadFile("new.txt")

	// if err_2 != nil {

	// 	fmt.Printf("%s\n", file)

	// }

	// fmt.Println(file)
	// ---------------- use file for ioutil (end)
	// -------------------- ioutil package (start)
	// file, err := ioutil.TempFile("test", "mms")

	// if err != nil {

	// 	fmt.Println(err)

	// 	return

	// }

	// fmt.Println(file)

	// for _, v := range file {

	// 	fmt.Println(v)

	// }
	// -------------------- ioutil package (end)
	//-------------------- append tip (start)
	// x := make([]int, 2, 8)

	// y := []int{4, 5, 6, 7, 8, 9}

	// z := append([]int{}, y...)

	// fmt.Println(z)
	//-------------------- append tip (end)
	// ----------------- cap in slice (start)
	// x := []int{1, 2, 3, 4, 5}

	// y := x[0:3:5]

	// fmt.Println(cap(y))
	// ----------------- cap in slice (end)
	// ----------------- cap slice for append(start)
	// var num []int

	// num = []int{1, 2}
	// fmt.Println(num)

	// num = append(num, 3, 4)
	// fmt.Println(num)

	// num = append(num, 3, 4)
	// fmt.Println(num)

	// num = append(num[2:4], 5, 7)
	// fmt.Println(num)

	// fmt.Println(num[:cap(num)])
	// ----------------- cap slice for append(end)
	// -------------- cap function (start)
	// sli := []int{1, 2, 3}
	// fmt.Println(cap(sli))
	// -------------- cap function (end)
	// --------------------- project (start)
	// var passwords []string

	// for i := 0; i < 25; i++ {

	// 	passwords = append(passwords, createPassword(8))
	// }

	// fmt.Println(passwords)
	// --------------------- project (end)
	// ----------------------- project time (start)
	// number := [...][5]string{
	// 	{
	// 		"000",
	// 		"0 0",
	// 		"0 0",
	// 		"0 0",
	// 		"000",
	// 	},
	// 	{
	// 		"00 ",
	// 		" 0 ",
	// 		" 0 ",
	// 		" 0 ",
	// 		"000",
	// 	},
	// 	{
	// 		"000",
	// 		"  0",
	// 		"000",
	// 		"0  ",
	// 		"000",
	// 	},
	// 	{
	// 		"000",
	// 		"  0",
	// 		"000",
	// 		"  0",
	// 		"000",
	// 	},
	// 	{
	// 		"0 0",
	// 		"0 0",
	// 		"000",
	// 		"  0",
	// 		"  0",
	// 	},
	// 	{
	// 		"000",
	// 		"0  ",
	// 		"000",
	// 		"  0",
	// 		"000",
	// 	},
	// 	{
	// 		"000",
	// 		"0  ",
	// 		"000",
	// 		"0 0",
	// 		"000",
	// 	},
	// 	{
	// 		"000",
	// 		"  0",
	// 		"  0",
	// 		"  0",
	// 		"  0",
	// 	},
	// 	{
	// 		"000",
	// 		"0 0",
	// 		"000",
	// 		"0 0",
	// 		"000",
	// 	},
	// 	{
	// 		"000",
	// 		"0 0",
	// 		"000",
	// 		"  0",
	// 		"000",
	// 	},
	// 	{
	// 		"   ",
	// 		" 0 ",
	// 		"   ",
	// 		" 0 ",
	// 		"   ",
	// 	},
	// }

	// hour := time.Now().Hour() / 10

	// hour_s := time.Now().Hour() % 10

	// second := time.Now().Second() / 10

	// second_s := time.Now().Second() % 10

	// minute := time.Now().Minute() / 10

	// minute_s := time.Now().Minute() % 10

	// for i := 0; i <= 0; i++ {

	// 	for j := 0; j <= 4; j++ {

	// 		fmt.Println(number[hour+1][j], " ", number[hour_s+1][j], " ", number[11][j], " ", number[minute+1][j], " ", number[minute_s+1][j], " ", number[11][j], " ", number[second+1][j], " ", number[second_s+1][j])

	// 	}

	// }

	// for _, v := range number {

	// 	for _, v2 := range v {

	// 		fmt.Printf("%s\n", v2)

	// 	}

	// }
	// ----------------------- project time (end)
	// ----------------- my type or new type (start)
	// type book [3]int
	// arr := book{1, 2, 3}

	// arr2 := [3]int{1, 2, 3}

	// fmt.Printf("%T\n", arr2)
	// ----------------- my type or new type (end)
	// ------------- keyed Element array(start)
	// arr := [...]int{
	// 	2: 45,
	// 	78,
	// 	5: 44,
	// }

	// fmt.Println(arr)
	// ------------- keyed Element array(end)
	// ---------------- arraymap (start)
	// arr1 := map[string]int{"item_1": 25, "item_2": 46}

	// fmt.Println(arr1)
	// ---------------- arraymap (end)
	// ---------------- mini project (start)
	// x := [5]int{1, 2, 3, 4, 5}

	// input := os.Args[1]

	// fmt.Printf("%s, => %d\n", input, x[rand.Intn(5)])
	// ---------------- mini project (end)
	// ------------------ path (start)
	// dir, file := path.Split("src/css/style.css")

	// fmt.Printf("Dir=> %v, File=> %v\n", dir, file)
	// ------------------ path (end)
	// p1 := ver{X: 1, Y: 80}

	// fmt.Println(p1)

	// var arr map[int]string

	// fmt.Println(arr)

	// --------------------- project (start)
	// text := os.Args[1]

	// var new_text string

	// for i := len(text) - 1; i >= 0; i-- {

	// 	new_text += string(text[i])

	// }

	// if text == new_text {

	// 	fmt.Println("Match... you string=> ", text, " reverce string=> ", new_text)

	// 	return

	// }

	// fmt.Println("Not match... you string=> ", text, " reverce string=> ", new_text)
	// --------------------- project (end)
	// ---------------- range for(start)
	// str := "" + "Lorem ipsum dolor sit amet, consectetur amet adipiscing elit, sed eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquid ex ea commodi consequat. Quis aute iure reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint obcaecat cupiditat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	// chunk_str := strings.Fields(str)

	// query := os.Args[1:]

	// for _, i := range query {

	// 	for _, j := range chunk_str {

	// 		if i == j {

	// 			fmt.Println("Search ok", i)

	// 		}

	// 	}

	// }
	// ---------------- range for(end)
	// --------------- Adad Aval and morakab (start)
	// var d float64

	// if len(os.Args) != 2 {

	// 	fmt.Println("Error Type only one number")

	// 	return

	// }

	// n, err := strconv.Atoi(os.Args[1])

	// if err != nil {

	// 	fmt.Println("Error a Number")

	// 	return

	// }

	// if n < 0 {

	// 	fmt.Println("Error number +")

	// 	return

	// }

	// // fmt.Println(math.Mod(5.1, 1.0))

	// for i := 2; i < n; i++ {

	// 	d = float64(n) / float64(i)

	// 	if math.Mod(d, 1.0) == 0 {

	// 		fmt.Println("Morakab")

	// 		return

	// 	}

	// }

	// fmt.Println("Aval")
	// --------------- Adad Aval and morakab (end)
	// ------------ Test rand (start)
	// for i := 0; i < 10; i++ {

	// 	n := rand.Intn(10)

	// 	fmt.Printf("%d ", n)
	// }
	// ------------ Test rand (end)
	// ------------- Proejct section (start)
	// you_num := rand.IntN(11)

	// var my_num int

	// var count int

	// for {

	// 	if count == 5 {

	// 		fmt.Printf("Game Over (%d)\n", you_num)

	// 		break

	// 	}

	// 	fmt.Printf("Type new Number = ")

	// 	(fmt.Scan(&my_num))

	// 	if my_num == you_num {

	// 		if count == 0 {

	// 			fmt.Printf("WoW run a first number winer =>(%d)\n", my_num)

	// 		} else {

	// 			fmt.Printf("Done number is =>(%d)\n", my_num)

	// 		}

	// 		break

	// 	}

	// 	fmt.Printf("not a number (%d)\n", my_num)

	// 	count++

	// }
	// ------------- Proejct section (end)
	// ------------------ Fibonacci number (start)
	// f := []int{0, 1, 1}

	// n, _ := strconv.Atoi(os.Args[1])

	// if n <= 3 {

	// 	fmt.Println(f[:n])

	// 	return

	// }

	// for i := 3; i < n; i++ {

	// 	o_p_n := int(f[i-1])

	// 	t_p_n := int(f[i-2])

	// 	f = append(f, o_p_n+t_p_n)

	// }

	// fmt.Println(f)
	// ------------------ Fibonacci number (end)
	// for _, v := range str_val {

	// 	fmt.Printf("%s", v)

	// }

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
