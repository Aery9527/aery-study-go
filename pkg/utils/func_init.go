package utils

import "fmt"

// init() 會在 main() 被呼叫之前執行
// 可以有多個 init(), 甚至在同一個檔案中也可以定義多次
// 每個檔案的 init() 執行順序依照 Go compiler 決定的 import 順序來定
// 不能帶參數或回傳值
// 主要用途就是為了模組初始化, 不需要特別去呼叫它, Go compiler 會搞定
// Side Effect Import(import _ "fmt") : import 一個 package 但沒打算用裡面的任何 symbol(function, struct, constant 等), 只想讓它的 init() 執行
func init() {
	fmt.Println("func_init.go init()...")
}
