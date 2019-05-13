package main

import "fmt"

type Dog struct {
	name string
}

func main() {
	var i int
	var f float32
	var r rune
	var s string
	var ia [2]int
	var inf interface{}
	var dog1 Dog
	var iptr *int
	var sl1 []int
	var m1 map[string]int

	fmt.Println(i, f, r, ia, inf, dog1, iptr, sl1, m1)
	fmt.Printf("%#v\n", s)
	i = 3
	f = 3.4
	r = 10
	s = "hello"
	ia[0] = 1
	ia[1] = 2
	inf = 3
	dog1.name = "wangcai"
	iptr = &i

	sl1 = make([]int, 2)
	sl1[0] = 11
	sl1[1] = 22

	m1 = make(map[string]int, 2)
	m1["aa"] = 3
	m1["bb"] = 4
	fmt.Println(i, f, r, s, ia, inf, dog1, iptr, sl1, m1)
}
