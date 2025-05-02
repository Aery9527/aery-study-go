package main

import (
	"aery-study-go/pkg/utils"
	"fmt"
	"reflect"
)

// slice 宣告方式, 大小可以變動, 類似 java 的 List<>
// 它不是真正意義上的動態陣列, 而是一個 reference type, 指向一個底層 array

var s0 []int // 宣告方式跟 array 一樣, 只是[]內沒有數字就會是 slice

func main() {
	a := [...]int{1, 2, 3, 4}

	s12 := a[1:3] // 生成 slice 為指向 a 的 1~2
	s02 := a[:3]  // 生成 slice 為指向 a 的 0~2, 從第一個元素[0]開始
	s03 := a[:]   // 生成 slice 為指向 a 的 0~3, 從第一個元素[0]開始, 到最後一個元素[3]結束
	a[1] = 0      // 所以這邊 a[1] 改變了, 上面幾個 slice 也會改變

	utils.WrapPrint("slice 型態", func() {
		fmt.Printf("a  :%v type(%T)\n", a, a)
		fmt.Printf("s12:%v type(%T) len(%d) cap(%d)\n", s12, s12, len(s12), cap(s12))
		fmt.Printf("s02:%v type(%T) len(%d) cap(%d)\n", s02, s02, len(s02), cap(s02))
		fmt.Printf("s03:%v type(%T) len(%d) cap(%d)\n", s03, s03, len(s03), cap(s03))
	})

	utils.WrapPrint("slice 比對", func() {
		fmt.Printf("&s12 == &s02:%v\n", &s12 == &s02) // slice 只能比較指標, 不能比較內容
		fmt.Printf("&s12 == &s03:%v\n", &s12 == &s03)
		fmt.Printf("&s02 == &s03:%v\n", &s02 == &s03) // 雖然都是指向 a 這個 array 的起點位置, 但其實是兩個 slice header 實體
		// reflect.DeepEqual() 深度比較內容是否相同但效能差(走 interface 與 reflect), 所以效能場景還是自己寫比對較快
		fmt.Printf("a    == s03:%v\n", reflect.DeepEqual(a, s03)) // 型態不同直接 false
		fmt.Printf("a[:] == s03:%v\n", reflect.DeepEqual(a[:], s03))
	})

	s5 := make([]int, 2, 3) // 長度 2, 容量 3
	utils.WrapPrint("make(slice)", func() {
		fmt.Printf("s5:%v type(%T) len(%d) cap(%d)\n", s5, s5, len(s5), cap(s5))
	})

	s6 := append(s5, 1) // 回傳的 slice 添加了元素 1, 但還是指向同個 array
	s7 := append(s5, 2) // 這邊又再添加一次元素, 其實還是對同個 array 添加, 但 s5 再添加是覆蓋第 3 個位置, 因此 s6[2] 的設定就會被蓋掉
	s6[0] = 9
	s7[1] = 8
	utils.WrapPrint("append(slice) <= cap", func() {
		fmt.Printf("s5:%v type(%T) len(%d) cap(%d)\n", s5, s5, len(s5), cap(s5))
		fmt.Printf("s6:%v type(%T) len(%d) cap(%d)\n", s6, s6, len(s6), cap(s6))
		fmt.Printf("s7:%v type(%T) len(%d) cap(%d)\n", s7, s7, len(s7), cap(s7))
	})

	s8 := append(s6, 3) // 超過本來 s5 的容量了, 所以會重新分配一個新的 array, 然後把原本的 array 複製過去
	s9 := append(s7, 4) // 同上
	s6[2] = 7
	s7[2] = 6 // 由於 s6 跟 s7 還是指向同個 array, 因此會覆蓋掉上面的設定
	s8[2] = 5 // 新的 array
	s9[2] = 4 // 新的 array
	utils.WrapPrint("append(slice) > cap", func() {
		fmt.Printf("s6:%v type(%T) len(%d) cap(%d)\n", s6, s6, len(s6), cap(s6))
		fmt.Printf("s7:%v type(%T) len(%d) cap(%d)\n", s7, s7, len(s7), cap(s7))
		fmt.Printf("s8:%v type(%T) len(%d) cap(%d)\n", s8, s8, len(s8), cap(s8))
		fmt.Printf("s9:%v type(%T) len(%d) cap(%d)\n", s9, s9, len(s9), cap(s9))
	})

	// 1.2 開始支援了三個參數的 slice[i:j:k], 從 i 開始, 取到第 j 個(不包含 j) 限制 capacity 為 k - i
	// 這樣就可以控制產生出來的 slice 的 capacity, 不會讓原 slice 往後的資料被意外影響或改到
	// 也因此只要讓 j=k, 在 append 的時候就不會影響到原本的 array
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sss := arr[2:4:7] // [2, 3], 長度2, 容量為5(7-2)
	utils.WrapPrint("slice[i:j:k]", func() {
		fmt.Printf("arr:%v type(%T) len(%d) cap(%d)\n", arr, arr, len(arr), cap(arr))
		fmt.Printf("sss:%v type(%T) len(%d) cap(%d)\n", sss, sss, len(sss), cap(sss))
	})

	sss = append(sss, 0) // 影響到原 array 的資料
	utils.WrapPrint("append(slice[i:j:k]) <= cap", func() {
		fmt.Printf("arr:%v type(%T) len(%d) cap(%d)\n", arr, arr, len(arr), cap(arr))
		fmt.Printf("sss:%v type(%T) len(%d) cap(%d)\n", sss, sss, len(sss), cap(sss))
	})

	fmt.Println()
	sss = append(sss, -1, -2, -3) // 超過 capacity 了, 所以會重新分配一個新的 array, 就不再影響原本的 array 了
	utils.WrapPrint("append(slice[i:j:k]) > cap", func() {
		fmt.Printf("arr:%v type(%T) len(%d) cap(%d)\n", arr, arr, len(arr), cap(arr))
		fmt.Printf("sss:%v type(%T) len(%d) cap(%d)\n", sss, sss, len(sss), cap(sss))
	})

	s0 := s12 // reference
	utils.WrapPrint("slice 傳遞是指標, 因此會指向同個底層 array", func() {
		fmt.Printf("s0  defore modifySlice : %v\n", s0)
		fmt.Printf("s12 defore modifySlice : %v\n", s12)
		modifySlice(s0)
		fmt.Printf("s0  after  modifySlice : %v\n", s0)
		fmt.Printf("s12 after  modifySlice : %v\n", s12)
	})
}

func modifySlice(a []int) {
	for i := range a {
		a[i] += 1
	}
}
