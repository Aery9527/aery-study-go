package main

import (
	"aery-study-go/pkg/where"
	"fmt"
)

// Java 的 Map 有多種實作應付不同情境, HashMap/TreeMap/LinkedHashMap 等
// Go 看似只給了一種 map[K]V, 但底層實作是經過精心優化的 "哈希表(hash table)", 足以應對大多數情境
// 所以若有特殊需求情境(例如有序 map), 則需要自己封裝或使用第三方套件達成

func main() {
	// map 宣告方式, key 可以是任何定義了 "==" 及 "!=" 操作的型別
	// map 是無序的, 長度不定, 同 slice 一樣是參考型別
	m := map[int]string{5566: "a", 9527: "b"} // 宣告同時賦值
	var m1 map[string]int                     // 這樣宣告是 nil, 就像 java 宣告物件後沒給東西是 null 一樣的概念
	m1 = make(map[string]int)                 // 這樣才是初始化一個可操作的 map

	// 雖然都是map, 型態不相同
	where.WrapPrint("map 型態", func() {
		fmt.Printf("m  type: %T = %v\n", m, m)
		fmt.Printf("m1 type: %T = %v\n", m1, m1) // %v 會將 nil 以該型別的空內容呈現, 例如 map[string]int 就會顯示 map[]
	})

	where.WrapPrint("map 賦值", func() {
		printMap("m1", m1)
		m1["a"] = 1 // 賦值
		m1["b"] = 2
		printMap("m1", m1)
	})

	where.WrapPrint("map 刪除", func() {
		delete(m1, "b") // XXX 居然沒回傳值??
		printMap("m1", m1)
	})

	where.WrapPrint("map 2 個回傳值", func() {
		a, aExist := m1["a"] // map 有兩個回傳值
		b, bExist := m1["b"]
		fmt.Printf("m1[%s] = %d %t\n", "a", a, aExist)
		fmt.Printf("m1[%s] = %d %t\n", "b", b, bExist)
	})

	where.WrapPrint("map 指標", func() {
		m2 := m1
		m2["c"] = 3 // 與 m1 指向同一個 map, 所以 m1 也會改變
		printMap("m1", m1)
		printMap("m2", m2)
	})
}

func printMap[K comparable, V any](name string, m map[K]V) {
	fmt.Printf("%s(%d) %v\n", name, len(m), m)
}
