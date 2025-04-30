package main

import (
	"aery-study-go/pkg/utils"
	"fmt"
)

func main() {
	// new 用於分配所有型別的記憶體分配

	// new 分配空間給一個 int 並回傳指標
	i := new(int)
	*i = 42
	utils.WrapPrint("new(int)", func() {
		fmt.Println("int i:", *i)
	})

	// new 分配空間給一個自訂 struct
	p := new(point)
	p.X = 10
	p.Y = 20
	utils.WrapPrint("new(struct)", func() {
		fmt.Printf("point p: %+v\n", *p)
	})
}

type point struct {
	X int
	Y int
}
