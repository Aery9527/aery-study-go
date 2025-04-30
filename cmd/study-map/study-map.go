package main

import (
	"aery-study-go/pkg/utils"
	"fmt"
)

func main() {
	// map 宣告方式, key 可以任何定義了 "==" 及 "!=" 操作的型別
	// map 是無序的, 長度不定, 同 slice 一樣是參考型別
	m := map[int]string{5566: "a", 9527: "b"} // 宣告同時賦值
	var m1 map[string]int                     // 這樣宣告是 nil, 就像 java 宣告物件後沒給東西是 null 一樣的概念
	m1 = make(map[string]int)                 // 這樣才有初始化一個 map

	// 雖然都是map, 型態不相同
	utils.WrapPrint("map 型態", func() {
		fmt.Printf("m  type: %T\n", m)
		fmt.Printf("m1 type: %T\n", m1)
	})

	utils.WrapPrint("map 賦值", func() {
		printMap("m1", m1)
		m1["a"] = 1 // 賦值
		m1["b"] = 2
		printMap("m1", m1)
	})

	utils.WrapPrint("map 刪除", func() {
		delete(m1, "b") // 刪除 XXX 居然沒回傳值??
		printMap("m1", m1)
	})

	utils.WrapPrint("map 2 個回傳值", func() {
		a, aExist := m1["a"] // map 有兩個回傳值
		b, bExist := m1["b"]
		fmt.Printf("m1[%s] = %d %t\n", "a", a, aExist)
		fmt.Printf("m1[%s] = %d %t\n", "b", b, bExist)
	})

	utils.WrapPrint("map 指標", func() {
		m2 := m1
		m2["c"] = 3 // 與 m1 指向同一個 map, 所以 m1 也會改變
		printMap("m1", m1)
		printMap("m2", m2)
	})
}

func printMap[K comparable, V any](name string, m map[K]V) {
	fmt.Printf("%s(%d) %v\n", name, len(m), m)
}
