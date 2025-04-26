package main

import (
	"aery-study-go/internal/study"
	"aery-study-go/internal/study/variable"
	"fmt"
	// . "fmt" 就像 java 的 import static 一樣, 這樣就可以直接用 Println() 了
	// f "fmt" 給定一個 alias, 這樣就可以用 f.Println() 了
	// _ "fmt" (side effect import), import 一個 package 但沒打算用裡面的任何 symbol(function, struct, constant 等), 只想讓它的 init() 執行就可以這樣寫
)

// GO 有兩個保留 func, init() 跟 main()

func init() {
	// init() 不能帶參數或回傳值, 會在 main() 被呼叫之前執行
	// 可以有多個 init(), 甚至在同一個檔案中也可以定義多次
	// 每個檔案的 init() 執行順序依照 Go compiler 決定的 import 順序來定
	// 主要用途就是為了模組初始化, 不需要特別去呼叫它, Go compiler 會搞定

	fmt.Println("yo~~~")
}

func main() {
	wrapPrint("ShowArgs", study.ShowArgs) // 把 method 當作參數傳入
	wrapPrint("ShowVar", func() { variable.ShowVar(1, "2", 3.14, true, byte(1)) })
	wrapPrint("ShowIota", variable.ShowIota)
	wrapPrint("ShowArray", variable.ShowArray)
	wrapPrint("ShowSlice", variable.ShowSlice)
	wrapPrint("ShowMap", variable.ShowMap)
	wrapPrint("ShowStruct", variable.ShowStruct)
	wrapPrint("ShowMake", study.ShowMake)
	wrapPrint("ShowNew", study.ShowNew)
	wrapPrint("ShowNil", study.ShowNil)
	wrapPrint("ShowIf", func() { study.ShowIf(10, func() int { return 1 }) })
	wrapPrint("ShowSwitch", func() { study.ShowSwitch("Blue Monday") })
	wrapPrint("ShowGoto", func() { study.ShowGoto() })
	wrapPrint("ShowFor", study.ShowFor)
	wrapPrint("ShowFunc", study.ShowFunc)
	wrapPrint("ShowPanic", study.ShowPanic)
	wrapPrint("ShowInterface", study.ShowInterface)
}

type show func() // 定義一個 func 型別的變數, 這樣就可以當作參數傳入, 跟 java lambda 一樣可以把 method 當參數丟來丟去

// type 像是別名的意思
type age int      // 像這樣把基本型別賦予一個領域意義
var aery age = 18 // 搭配變數名稱就可以產生 2 個維度含意, 也就是利用程式寫法把意義大致說明清楚了, 可以大幅減少註解說明

func wrapPrint(scope string, action show) {
	fmt.Println()
	fmt.Printf("===== %s =====\n", scope)
	action()
}
