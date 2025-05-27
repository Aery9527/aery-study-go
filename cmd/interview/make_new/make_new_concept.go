package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	s[1] = 3
	s[2] = 6
	s[3] = 9
	fmt.Printf("%v\n", s) // XXX 輸出為何?

	m1 := make(map[string]int)
	m2 := map[string]int(nil)
	put(m1)
	put(m2)
	fmt.Printf("%v\n", m1) // XXX 輸出為何?
	fmt.Printf("%v\n", m2) // XXX 輸出為何?

	// XXX 這兩個 channel 有何不同?
	ch1 := make(chan int)
	ch2 := make(chan int, 2)
	fmt.Printf("%v\n", <-ch1)
	fmt.Printf("%v\n", <-ch2)

	// XXX 為何只有這三種型別使用 make 創建?

	i := new([3]int)
	fmt.Printf("%v\n", *i) // XXX 輸出為何?
}

func put(m map[string]int) {
	m["key"] = 1
}
