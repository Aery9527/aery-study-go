package main

import (
	"aery-study-go/pkg/utils"
	"fmt"
)

// 在型別前加上"*"就表示該變數為指標, 被宣告為指標的變數其內容是一個地址, 指向真實資料的記憶體位址
// 傳指標時會做 escape analysis (逃逸分析), 如果其內容離開 scope 會被放到 heap 上, 後續自動 GC
// 使用時機就是當該 array 或 struct 等資料結構過於龐大時,
// pass by value 會複製全部的內容到 func 裡去, 就會造成大量記憶體的浪費,
// 這時候就可以使用指標來傳遞, 這樣就只會傳遞一個 "記憶體位址" 過去 func,
// 但相對的 func 內就有可能會影響到本來的資料

func main() {
	var x int = 10
	var p *int = &x // p 是指向 x 的指標

	utils.WrapPrint("原始值", func() { fmt.Println(*p) })
	modify(p)
	utils.WrapPrint("修改後", func() { fmt.Println(*p) })

	ps := &Person{Name: "Aery"}
	ps.Name = "Goddess" // 這邊 go 自動幫解指標, 不用(*ps).Name
	utils.WrapPrint("struct", func() { fmt.Println(ps.Name) })

	var i interface{} = x // interface 不是指標
	utils.WrapPrint("interface", func() {
		fmt.Println(i == *p) // true
		fmt.Println(i == p)  // false
	})
}

func modify(n *int) {
	*n = 100
}

type Person struct {
	Name string
}
