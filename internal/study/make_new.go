package study

import "fmt"

func ShowMake() {
	// 用於內建型別 map/slice/channel 的記憶體分配, 同時會回傳相對應的初始值(結構)
	// 這三個有所不同是因為建立後要能直接使用必須要先有一些初始值, 否則就會像 java 一樣 "Object a;" 其實是 null 一樣

	// 建立一個 map
	m := make(map[string]int)
	m["apple"] = 5
	m["banana"] = 10
	fmt.Println("map m:", m)

	// 建立一個 slice，長度 3，容量 5
	s := make([]int, 3, 5)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println("slice s:", s)

	// 建立一個 channel，容量 2
	ch := make(chan int, 2)
	ch <- 42
	ch <- 100
	fmt.Println("channel ch:", <-ch, <-ch)
}

func ShowNew() {
	// new 用於分配所有型別的記憶體分配

	// new 分配空間給一個 int 並回傳指標
	i := new(int)
	*i = 42
	fmt.Println("int i:", *i)

	// new 分配空間給一個自訂 struct
	p := new(Point)
	p.X = 10
	p.Y = 20
	fmt.Printf("Point p: %+v\n", *p)
}

type Point struct {
	X int
	Y int
}
