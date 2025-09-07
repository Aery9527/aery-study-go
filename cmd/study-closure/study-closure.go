package main

import "fmt"

func main() {
	a := 0

	func() {
		a++ // pointer to a
	}()

	fmt.Println(a)
}
