package variable

import (
	"fmt"
	"reflect"
)

// array 宣告方式, 大小固定不能變動
var a1 [1]int
var a2 = [2]int{1, 2}
var a3 = [3]int{1, 2}         // 沒賦值的就是預設值, int 就是0
var a4 = [...]int{1, 2, 3, 4} // 透過賦值決定陣列大小
var aa = [3][4]int{           // 可以寫[...][4], 但不能寫[...][...], 因為第二層開始無法透過定義推斷陣列大小
	{1, 2, 3, 4},
	{5, 6, 7},
} // 二維陣列
var aaa = [...][2][2]int{
	{{1, 1}, {2}},
	{{}, {}},
} // 三維陣列, 第一層大小透過定義推斷

// slice 宣告方式, 大小可以變動, 類似 java 的 List<>
// 它不是真正意義上的動態陣列, 而是一個 reference type, 指向一個底層 array
// slice 是個結構, 包含三個欄位: 指向底層 array 的指標、slice 的長度、slice 的容量
var s0 []int // 宣告方式跟 array 只是[]內沒有數字, 則型態就會是 slice

func ShowArray() {
	// 長度也是型別的一部份, 所以這邊顯示的 type 是帶長度的, 也就是說他們是不同型別
	fmt.Printf("a1:%v type(%T) len(%d) cap(%d)\n", a1, a1, len(a1), cap(a1))
	fmt.Printf("a2:%v type(%T) len(%d) cap(%d)\n", a2, a2, len(a2), cap(a2))
	fmt.Printf("a3:%v type(%T) len(%d) cap(%d)\n", a3, a3, len(a3), cap(a3))
	fmt.Printf("a4:%v type(%T) len(%d) cap(%d)\n", a4, a4, len(a4), cap(a4))
	fmt.Printf("aa:%v type(%T) len(%d) cap(%d)\n", aa, aa, len(aa), cap(aa))
	fmt.Printf("aaa:%v type(%T) len(%d) cap(%d)\n", aaa, aaa, len(aaa), cap(aaa))

	A1 := a1 // array 重新賦值會複製整個 array, 其比對也是完整比對內容相等
	fmt.Printf("A1 == a2:%v\n", A1 == a1)
	a1[0] = 9 // a1 改變了, A1 不會改變
	fmt.Printf("a1:%v\n", a1)
	fmt.Printf("A1:%v\n", A1)
	fmt.Printf("A1 == a2:%v\n", A1 == a1)
}

func ShowSlice() {
	fmt.Println()
	s12 := a4[1:3] // 生成 slice 為指向 a4 的 1~2
	s02 := a4[:3]  // 生成 slice 為指向 a4 的 0~2, 從第一個元素[0]開始
	s03 := a4[:]   // 生成 slice 為指向 a4 的 0~3, 從第一個元素[0]開始, 到最後一個元素[3]結束
	a4[1] = 0      // 所以這邊 a4[1] 改變了, 上面幾個 slice 也會改變
	fmt.Printf("a4:%v type(%T)\n", a4, a4)
	fmt.Printf("s12:%v type(%T) len(%d) cap(%d)\n", s12, s12, len(s12), cap(s12))
	fmt.Printf("s02:%v type(%T) len(%d) cap(%d)\n", s02, s02, len(s02), cap(s02))
	fmt.Printf("s03:%v type(%T) len(%d) cap(%d)\n", s03, s03, len(s03), cap(s03))
	fmt.Printf("&s12 == &s02:%v\n", &s12 == &s02) // slice 只能比較指標, 不能比較內容
	fmt.Printf("&s12 == &s03:%v\n", &s12 == &s03)
	fmt.Printf("&s02 == &s03:%v\n", &s02 == &s03) // 雖然都是指向 a4 這個 array 的起點位置, 但其實是兩個 slice header 實體
	// reflect.DeepEqual() 深度比較內容是否相同但效能差(走 interface 與 reflect), 所以效能場景還是自己寫比對較快
	fmt.Printf("a4    == s03:%v\n", reflect.DeepEqual(a4, s03)) // 型態不同直接 false
	fmt.Printf("a4[:] == s03:%v\n", reflect.DeepEqual(a4[:], s03))

	fmt.Println()
	s5 := make([]int, 2, 3)
	fmt.Printf("s5:%v type(%T) len(%d) cap(%d)\n", s5, s5, len(s5), cap(s5))

	fmt.Println()
	s6 := append(s5, 1) // 回傳的 slice 添加了元素 1, 但還是指向同個 array
	s7 := append(s5, 2) // 這邊又再添加一次元素, 其實還是對同個 array 添加, 但 s5 添加是覆蓋第 3 個位置, 因此把 s6[2] 覆蓋掉
	s6[0] = 9
	s7[1] = 8
	fmt.Printf("s5:%v type(%T) len(%d) cap(%d)\n", s5, s5, len(s5), cap(s5))
	fmt.Printf("s6:%v type(%T) len(%d) cap(%d)\n", s6, s6, len(s6), cap(s6))
	fmt.Printf("s7:%v type(%T) len(%d) cap(%d)\n", s7, s7, len(s7), cap(s7))

	fmt.Println()
	s8 := append(s6, 3) // 超過本來 s5 的容量了, 所以會重新分配一個新的 array, 然後把原本的 array 複製過去
	s9 := append(s7, 4) // 同上
	s6[2] = 7
	s7[2] = 6 // 由於 s6 跟 s7 還是指向同個 array, 因此會覆蓋掉上面的設定
	s8[2] = 5 // 新的 array
	s9[2] = 4 // 新的 array
	fmt.Printf("s6:%v type(%T) len(%d) cap(%d)\n", s6, s6, len(s6), cap(s6))
	fmt.Printf("s7:%v type(%T) len(%d) cap(%d)\n", s7, s7, len(s7), cap(s7))
	fmt.Printf("s8:%v type(%T) len(%d) cap(%d)\n", s8, s8, len(s8), cap(s8))
	fmt.Printf("s9:%v type(%T) len(%d) cap(%d)\n", s9, s9, len(s9), cap(s9))

	// 1.2 開始支援了三個參數的 slice[i:j:k], 從 i 開始, 取到第 j 個(不包含 j) 限制 capacity 為 k - i
	// 這樣就可以控制產生出來的 slice 的 capacity, 不會讓原 slice 往後的資料被意外影響或改到
	// 也因此只要讓 j=k, 在 append 的時候就不會影響到原本的 array
	fmt.Println()
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sss := arr[2:4:7] // [2, 3], 長度2, 容量為5(7-2)
	fmt.Printf("arr:%v type(%T) len(%d) cap(%d)\n", arr, arr, len(arr), cap(arr))
	fmt.Printf("sss:%v type(%T) len(%d) cap(%d)\n", sss, sss, len(sss), cap(sss))

	fmt.Println()
	sss = append(sss, 0) // 影響到原 array 的資料
	fmt.Printf("arr:%v type(%T) len(%d) cap(%d)\n", arr, arr, len(arr), cap(arr))
	fmt.Printf("sss:%v type(%T) len(%d) cap(%d)\n", sss, sss, len(sss), cap(sss))

	fmt.Println()
	sss = append(sss, -1, -2, -3) // 超過 capacity 了, 所以會重新分配一個新的 array, 就不再影響原本的 array 了
	fmt.Printf("arr:%v type(%T) len(%d) cap(%d)\n", arr, arr, len(arr), cap(arr))
	fmt.Printf("sss:%v type(%T) len(%d) cap(%d)\n", sss, sss, len(sss), cap(sss))
}
