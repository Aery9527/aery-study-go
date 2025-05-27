package main

import "fmt"

func main() {
	fmt.Printf("%d\n", add(1, 1)) // XXX 輸出為何?

	defer func() {
		fmt.Printf("A 何時執行?") // XXX
	}()
	fmt.Printf("%d\n", sub(3, 1)) // XXX 輸出為何?
	defer func() {
		fmt.Printf("B 何時執行?") // XXX
	}()
}

func add(x, y int) (result int) {
	defer func() {
		result += 2
	}()
	defer func() {
		result *= 2
	}()

	result = x + y
	return
}

func sub(x, y int) (result int) {
	result = x - y
	if result%2 == 0 {
		panic("result is even")
	} else {
		return
	}
}
