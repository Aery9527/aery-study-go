package main

import (
	"aery-study-go/internal/study"
	"aery-study-go/internal/study/variable"
	"fmt"
)

func main() {
	wrapPrint("ShowArgs", study.ShowArgs) // 把 method 當作參數傳入
	wrapPrint("ShowVar", func() { variable.ShowVar(1, "2", 3.14, true, byte(1)) })
	wrapPrint("ShowIota", variable.ShowIota)
	wrapPrint("ShowArray", variable.ShowArray)
	wrapPrint("ShowSlice", variable.ShowSlice)
	wrapPrint("ShowMap", variable.ShowMap)
	wrapPrint("ShowIf", func() { study.ShowIf(10, func() int { return 1 }) })
	wrapPrint("ShowSwitch", func() { study.ShowSwitch("Blue Monday") })
	wrapPrint("ShowFor", study.ShowFor)
	wrapPrint("ShowFunc", study.ShowFunc)
	wrapPrint("ShowNil", study.ShowNil)
}

func wrapPrint(scope string, action func()) {
	fmt.Println()
	fmt.Printf("===== %s =====\n", scope)
	action()
}
