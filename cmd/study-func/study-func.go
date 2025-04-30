package main

import (
	"aery-study-go/pkg/utils"
	"fmt"
)

// func 沒有多載, 要從 func name 區分
// 這個設計我覺得比 java 多載性質好, 因為可以強制說明 func 功能區別

type funny func(name string) int // 定義一個型別為匿名函數

func main() {
	//_, remainder := divide(10, 3) // 使用 _ 可以忽略回傳值
	quotient, remainder := divide(10, 3)
	utils.WrapPrint("func init()", func() {
		fmt.Printf("商: %d, 餘數: %d\n", quotient, remainder)
	})

	a, b := split(10)
	utils.WrapPrint("func init()", func() {
		fmt.Printf("x: %d, y: %d\n", a, b)
	})

	d1, d2 := deferTest()
	utils.WrapPrint("func init()", func() {
		fmt.Printf("d1: %d, d2: %s\n", d1, d2)
	})

	result, err := errorTest()
	utils.WrapPrint("func init()", func() {
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			fmt.Printf("result: %s\n", result)
		}
	})

	x := 1
	pointerTest(&x) // 傳址, 參考型別
	utils.WrapPrint("func init()", func() {
		fmt.Printf("x: %d\n", x)
	})

	showAeryFunny(singAndDance) // 定義一個型別為匿名函數, 只要 func 的參數跟回傳值相同, 就可以直接當作 var 在 func 之間傳遞
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
	// 使用情境:想統一加 log/metrics(但不影響回傳值)/安全清理資源/優雅地處理error/recover
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

func pointerTest(a *int) {
	*a += 1
}

// 錯誤用回傳值處理
func errorTest() (string, error) {
	// 假設某種錯誤發生
	return "", fmt.Errorf("出錯囉")
}

func singAndDance(name string) int {
	score := 100
	utils.WrapPrint("func init()", func() {
		fmt.Printf("%s唱歌加跳舞獲得了%d分\n", name, score)
	})
	return score
}

func showAeryFunny(f funny) {
	f("Aery")
}
