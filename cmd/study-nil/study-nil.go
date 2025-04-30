package main

import "fmt"

// nil 是某些型別的"零值"或"空值"的代表
// nil 這個詞自拉丁文"nihil"(什麼都沒有), 也源自於 Lisp, Pascal, Modula, Smalltalk 等語言

func main() {
	var ptr *int
	fmt.Println("ptr == nil:", ptr == nil)

	var s []int
	fmt.Println("slice == nil:", s == nil)

	var m map[string]int
	fmt.Println("map == nil:", m == nil)

	var ch chan int
	fmt.Println("channel == nil:", ch == nil)

	var fn func()
	fmt.Println("function == nil:", fn == nil)

	var i interface{}
	fmt.Println("interface == nil:", i == nil)

	var err error
	fmt.Println("error == nil:", err == nil)

	var ptr2 *int = nil
	var i2 interface{} = ptr2 // i2 因為是裝了 "prt2" 這個指標, 所以它不是 nil
	fmt.Println("interface with nil pointer == nil:", i2 == nil)

	// 永遠都有值, 不允許 nil 判斷
	// var a int    // 基本型別
	// var b Person // struct
	// var c [3]int // array
}

type Person struct {
	Name string
	Age  int
}
