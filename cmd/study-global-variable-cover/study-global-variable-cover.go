package main

import (
	"aery-study-go/pkg/where"
	"fmt"
)

var g = 9527

func getG() int {
	return g
}

func main() {
	where.WrapPrint("全域變數", func() { fmt.Println(g) }) // 使用到全域變數

	g := 5566                                          // 這個 scope 底下重新定義了 g
	where.WrapPrint("區域變數", func() { fmt.Println(g) }) // 使用到區域變數

	// 要重新取得全域 g, 就得透過 func 了
	where.WrapPrint("getG()", func() { fmt.Println(getG()) })
}
