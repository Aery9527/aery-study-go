package main

import (
	"aery-study-go/pkg/utils"
	"fmt"
)

func main() {
	// array 宣告方式[]內一定要有數字, 大小固定不能變動
	var a1 [1]int
	a2 := [2]int{1, 2}
	a3 := [3]int{1, 2}         // 沒賦值的就是預設值, int 就是0
	a4 := [...]int{1, 2, 3, 4} // 透過賦值決定陣列大小
	aa := [3][4]int{           // 可以寫[...][4], 但不能寫[...][...], 因為第二層開始無法透過定義推斷陣列大小
		{1, 2, 3, 4},
		{5, 6, 7},
	} // 二維陣列
	aaa := [...][2][2]int{
		{{1, 1}, {2}},
		{{}, {}},
	} // 三維陣列, 第一層大小透過定義推斷

	// 長度也是型別的一部份, 所以這邊顯示的 type 是帶長度的, 也就是說他們是不同型別
	utils.WrapPrint("array 型別", func() {
		fmt.Printf("a1:%v type(%T) len(%d) cap(%d)\n", a1, a1, len(a1), cap(a1))
		fmt.Printf("a2:%v type(%T) len(%d) cap(%d)\n", a2, a2, len(a2), cap(a2))
		fmt.Printf("a3:%v type(%T) len(%d) cap(%d)\n", a3, a3, len(a3), cap(a3))
		fmt.Printf("a4:%v type(%T) len(%d) cap(%d)\n", a4, a4, len(a4), cap(a4))
		fmt.Printf("aa:%v type(%T) len(%d) cap(%d)\n", aa, aa, len(aa), cap(aa))
		fmt.Printf("aaa:%v type(%T) len(%d) cap(%d)\n", aaa, aaa, len(aaa), cap(aaa))
	})

	A1 := a1 // 指定給另外一個變數是複製整個 array(傳遞給 func 也是一樣), 其比對也是完整比對內容相等
	utils.WrapPrint("array 複製比對", func() {
		fmt.Printf("A1 == a2:%v\n", A1 == a1)
	})

	a1[0] = 9 // a1 改變了, A1 不會改變
	utils.WrapPrint("array 複製給值", func() {
		fmt.Printf("a1:%v\n", a1)
		fmt.Printf("A1:%v\n", A1)
		fmt.Printf("A1 == a2:%v\n", A1 == a1)
	})
}
