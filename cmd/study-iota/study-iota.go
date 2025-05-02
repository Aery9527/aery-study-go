package main

import (
	"fmt"
	"strconv"
)

// iota 是一個常數產生器, 來自希臘字母表的第九個字母"ι"
// 在英文裡 iota 也常被用來表示"極小的量", 但 Go 裡是借用了"小東西開始累積"概念
// 當成"從 0 開始自動遞增的計數器", 因此在 const 區塊中從 0 開始每出現一次就自動遞增一次,
// 最常被用來定義一組有意義的整數常數, 例如列舉(enum)

const (
	a       = iota
	b       = iota
	c       = "123"            // iota 在同一行值不同
	d                          // 未寫預設值逾值與上一個相同 (go 的淺規則)
	e       = iota             // iota 是以一個 const 為 group, 以宣告的行數來計算, 所以這裡 e 會是 4, 跳過前面兩個 c(2) 和 d(3)
	f, g, h = iota, iota, iota // iota 在同一行值相同
	i       = iota
)

const j = iota // 每遇到一個 const iota 就會重置

func main() {
	fmt.Println("a=" + strconv.Itoa(a))
	fmt.Println("b=" + strconv.Itoa(b))
	fmt.Println("c=" + c)
	fmt.Println("d=" + d)
	fmt.Println("e=" + strconv.Itoa(e))
	fmt.Println("f=" + strconv.Itoa(f))
	fmt.Println("g=" + strconv.Itoa(g))
	fmt.Println("h=" + strconv.Itoa(h))
	fmt.Println("i=" + strconv.Itoa(i))
	fmt.Println("j=" + strconv.Itoa(j))
}
