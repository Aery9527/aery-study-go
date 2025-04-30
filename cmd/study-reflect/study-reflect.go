package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i any = "9527"
	var t reflect.Type = reflect.TypeOf(i)   // 得到型別的 Meta 資料, 透過 t 我們能取得型別定義裡面的所有元素
	var v reflect.Value = reflect.ValueOf(i) // 得到實際的值, 透過 v 我們取得儲存在裡面的值, 還可以去改變值

	fmt.Printf("t %s (%T)\n", t, t)
	fmt.Printf("v %s (%T)\n", v, v)
}
