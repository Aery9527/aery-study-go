package utils

import "fmt"

type Show func()

func WrapPrint(scope string, action Show) {
	fmt.Println()
	fmt.Printf("[%s] ", scope)
	PrintWhereAt(1)
	action()
}
