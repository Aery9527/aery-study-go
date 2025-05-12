package main

import (
	"aery-study-go/pkg/where"
	"errors"
	"fmt"
	"reflect"
)

// 大寫開頭就是 public, 小寫開頭就是 private, 僅有這種兩種可見性

// 整數型態
var _int8 int8
var _int16 int16
var _int32 int32
var _int64 int64
var _int int // 平台決定 32 or 64
// 加個u是無號型態
var _uint8 uint8
var _uint16 uint16
var _uint32 uint32
var _uint64 uint64
var _uint uint // 平台決定 32 or 64

// 浮點數類型
var _float32 float32
var _float64 float64 // default

// 複數(Complex Numbers)
var _complex64 complex64   // float32 實部 + 虛部
var _complex128 = 95 + 27i // float64 實部 + 虛部

// other
var _byte byte // uint8 的別名 (type byte = uint8)
var _true bool
var _false = false        // 型態自動推斷, 宣告時給定值可以省略型態
var _string string = "字串" // 不可變更

// unicode, Go 是 "Unicode-first 語言", 所以沒有 char 概念, 直接就是操作每個 unicode 符號
// Rune 是北歐古代用來寫字的"符文", 每個符號都代表一個意思, 所以用這個詞來代表一個"Unicode code point(字元的唯一代碼編號)"
var _rune rune = '😎' // int32 的別名 (type rune = int32), 因為是 4 bytes 所以完整支援 unicode

// any 是 interface 的別名, 可以是任何型別
var everything any = "everything"

// 常數, 不可變更
const format = "%-20s: "

// 任何一個地方使用底線(_)當作變數名稱, 稱為"匿名變數"
// 匿名變數在後續無法被使用, 其主要用途是當變數用不到但又必須接起來時, 可避免代碼中的雜訊
// 在 java 21 也導入了這個用法
var _ int

func main() {
	showVar(1, "2", 3.14, true, byte(1))
}

func showVar(args ...any) { // 1.18 新增 any, 是 interface{} 的別名, 這樣寫比較簡潔, 所以也可以寫 interface{}
	where.WrapPrint("傳入參數", func() {
		for index, arg := range args {
			fmt.Printf(format+"[%d] %v (%T)\n", "args", index, arg, arg)
		}
	})

	a := 123        // block scope 才能使用 := 語法糖, package scope 不能使用
	b := float64(a) // 轉型
	where.WrapPrint("轉型", func() {
		fmt.Printf(format+"%d (%T)\n", "a", a, a) // %T 直接顯示型態
		fmt.Printf(format+"%f (%T)\n", "b", b, b)
	})

	// reflect.TypeOf 取得型態
	where.WrapPrint("取得型態", func() {
		fmt.Printf(format+"%d (%s)\n", "_int", _int, reflect.TypeOf(_int))
		fmt.Printf(format+"%d (%s)\n", "_uint", _uint, reflect.TypeOf(_uint))
		fmt.Printf(format+"%d (%s)\n", "_byte", _byte, reflect.TypeOf(_byte))                // type 是 uint8, 因為 byte 是 uint8 的別名
		fmt.Printf(format+"%v (%s)\n", "_complex64", _complex64, reflect.TypeOf(_complex64)) // %v 是萬用型的格式化符號, Go 會自動處理各種型別
		fmt.Printf(format+"%v (%s)\n", "_complex128", _complex128, reflect.TypeOf(_complex128))
	})

	s := `123
456
	789
000`
	where.WrapPrint("多行字串", func() {
		fmt.Printf(format+"%v (%s)\n", "s", s, reflect.TypeOf(s))
	})

	// error 型別
	err := errors.New("oops")
	where.WrapPrint("錯誤型別", func() {
		fmt.Println(err)
	})
}
