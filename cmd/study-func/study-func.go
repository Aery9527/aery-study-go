package main

import (
	"aery-study-go/pkg/where"
	"fmt"
)

// func 可以有多個回傳值, 沒有多載,
// 這個設計我覺得比 java 多載性質好,
// 因為可以強制開發者使用 func name 將功能說清楚

// 沒有像 java 這樣 `() -> {}` 匿名函數語法糖, 只能使用 `func() {}` 來表示匿名函數

// 定義一個 func 型別
type funny func(name string) int // name string 是參數, int 是回傳值

// 不定引數使用方式, 基本上跟 java 一樣,
// 不定引數要放在最後一個, 前面可以放其他引數,
// 雖然進到 func 內是 slice, 但呼叫時不能丟 slice 進去, 這點跟 java 可以接受 array 不一樣
type funnyMany func(showTimes int, names ...string) int

func main() {
	//_, remainder := divide(10, 3) // 使用 _ 可以忽略回傳值
	quotient, remainder := divide(10, 3)
	where.WrapPrint("func init()", func() {
		fmt.Printf("商: %d, 餘數: %d\n", quotient, remainder)
	})

	a, b := split(10)
	where.WrapPrint("func init()", func() {
		fmt.Printf("x: %d, y: %d\n", a, b)
	})

	d1, d2 := deferTest()
	where.WrapPrint("func init()", func() {
		fmt.Printf("d1: %d, d2: %s\n", d1, d2)
	})

	result, err := errorTest()
	where.WrapPrint("func init()", func() {
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			fmt.Printf("result: %s\n", result)
		}
	})

	x := 1
	add1(&x) // 傳址, 參考型別
	where.WrapPrint("func init()", func() {
		fmt.Printf("x: %d\n", x)
	})

	where.WrapPrint("func init()", func() {
		showAeryFunny(singAndDance) // 定義一個型別為匿名函數, 只要 func 的參數跟回傳值相同, 就可以直接當作 var 在 func 之間傳遞
	})

	where.WrapPrint("func init()", func() {
		showAeryFunnyMany(func(showTimes int, names ...string) int {
			fmt.Printf("%v 愛情動作片演出 %d 次\n", names, showTimes)
			return 0
		})
	})
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

// 命名回傳值, 要嘛全命名, 要嘛全都不要命名
func split(sum int) (x, y int) { // 採用命名回傳值, 若只有最後一個寫型態, 則是一個語法糖, 表示所有回傳值都是該型態
	x = sum * 4 / 9 // 整數除法運算跟 java 一樣, 結果無條件捨去, 一樣返回 int
	y = sum
	return // 會自動 return x, y
}

func deferTest() (x int, y string) {
	// defer 類似 java 的 finally, 在當前 func 結束後, 但真正 return 前執行
	// 只可以影響 "命名回傳值" 的內容, 多個 defer 採 LIFO(後進先出) 執行
	// defer 使用情境: 統一加 log/metrics(但不影響回傳值)/安全清理資源/優雅地處理error/recover
	defer func() { // 後執行
		x += 10
		y += "?"
	}()
	defer func() { // 先執行
		x += 2
		y += "!"
	}()
	return
}

func add1(a *int) {
	*a += 1
}

// 錯誤用回傳值處理
func errorTest() (string, error) {
	// 假設某種錯誤發生
	return "", fmt.Errorf("出錯囉")
}

func singAndDance(name string) int {
	fmt.Printf("%s 唱歌加跳舞\n", name)
	return 100
}

func showAeryFunny(f funny) {
	name := "Aery"
	cost := f(name)
	fmt.Printf("%s 酬勞報價 %d\n", name, cost)
}

func showAeryFunnyMany(f funnyMany) {
	artists := []string{"Aery", "Rion", "Yuma"}
	//f(20, artists) // 雖然引數是 slice, 但不能這樣傳入, 必須向下面這樣展開
	cost := f(20, artists...)
	fmt.Printf("%v 酬勞報價 %d\n", artists, cost)
}
