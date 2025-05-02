package main

import (
	"aery-study-go/pkg/utils"
	"fmt"
	"reflect"
)

// 1.18 開始支援泛型, 有比 java 更彈性的泛型設計
// "~" 是用來表示型別約束, 就像是 java 的 Object 繼承概念

type customString string
type customInt int
type customFloat float64

type keyConstraint interface {
	~string | int // 允許 "string同系(子類)" 或 "只能是int" 的別名
}
type valueConstraint interface {
	float64 | ~int // 允許 "只能是float64" 或 "int同系(子類)" 的別名
}

// Stack struct 泛型寫法
type Stack[T any] struct {
	items []T
}

func main() {
	utils.WrapPrint("receiveAnySlice[string]", func() { receiveAnySlice([]string{"a", "b", "c"}) })
	utils.WrapPrint("receiveAnySlice[int]", func() { receiveAnySlice([]int{1, 2, 3, 4}) })

	utils.WrapPrint("receiveMap", func() {
		receiveMap(map[string]any{
			"Aery": 1,
			"Rion": "A",
		})
	})

	utils.WrapPrint("receiveConstraintMap map[customString]customInt", func() {
		receiveConstraintMap(map[customString]customInt{
			"Aery": 1,
			"Rion": 2,
		})
	})
	utils.WrapPrint("receiveConstraintMap map[int]customInt", func() {
		receiveConstraintMap(map[int]customInt{ // key 不能用 customInt
			10: 1,
			20: 2,
		})
	})
	utils.WrapPrint("receiveConstraintMap map[int]float64", func() {
		receiveConstraintMap(map[int]float64{ // value 不能用 customFloat
			10: 1.1,
			20: 2.2,
		})
	})
	utils.WrapPrint("receiveConstraintMap map[string]int", func() {
		receiveConstraintMap(map[string]int{
			"Aery": 1,
			"Rion": 2,
		})
	})

	utils.WrapPrint("IsEqual[int]", func() { IsEqual(1, 2) })
}

// receiveAnySlice 接收泛型, T 可以是任意型別
func receiveAnySlice[T any](slice []T) {
	for index, element := range slice {
		fmt.Println(index, element)
	}
}

func receiveMap[K string, V any](m map[K]V) {
	for key, value := range m {
		fmt.Println(key, reflect.TypeOf(key), value, reflect.TypeOf(value))
	}
}

func receiveConstraintMap[K keyConstraint, V valueConstraint](m map[K]V) {
	for key, value := range m {
		fmt.Println(key, reflect.TypeOf(key), value, reflect.TypeOf(value))
	}
}

// IsEqual comparable 表示任何能進行 == / != 比較的型別(int, string, struct, ...)
func IsEqual[T comparable](a, b T) {
	fmt.Printf("%v(%T) == %v(%T)", a, a, b, b)
}
