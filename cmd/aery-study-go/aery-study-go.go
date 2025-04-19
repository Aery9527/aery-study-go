package main

import (
	"aery-study-go/internal/study"
	"fmt"
)

func main() {
	wrap_print("ShowArgs", func() { study.ShowArgs() })
	wrap_print("ShowVar", func() { study.ShowVar() })
}

func wrap_print(scope string, action func()) {
	fmt.Printf("===== %s =====\n", scope)
	action()
	fmt.Println()
}
