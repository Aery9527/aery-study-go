package main

import (
	"aery-study-go/pkg/utils"
	"fmt"
	"reflect"
)

func main() {
	utils.WrapPrint("int", func() { showType(9527) })
	utils.WrapPrint("string", func() { showType("9527") })
}

func showType(a any) {
	var t reflect.Type = reflect.TypeOf(a)   // 得到型別的 Meta 資料, 透過 t 我們能取得型別定義裡面的所有元素
	var v reflect.Value = reflect.ValueOf(a) // 得到實際的值, 透過 v 我們取得儲存在裡面的值, 還可以去改變值

	fmt.Printf("t %v (%T)\n", t, t)
	fmt.Printf("v %v (%T)\n", v, v)
}
