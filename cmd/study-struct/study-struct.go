package main

import (
	"aery-study-go/pkg/utils"
	"fmt"
)

// struct 就是一組資料的概念
// struct 若要直接使用 == 或 != 比對語法, 那麼其 field 都必須是可比對的型別

// 無法直接給預設值, 若真要給預設值則寫一個 func 當作 constructor 給定預設值並返回
type person struct {
	name string
	age  int
}

type student struct {
	person // 匿名欄位, 就像繼承概念
	school string
	age    int // 匿名欄位重複並不衝突, 因為處在不同 scope
}

// 定義 struct 身上的 func, 前面的那個 () 就稱為 receiver
func (p person) string() string { // pass by value
	return fmt.Sprintf("name: %s, age: %d", p.name, p.age)
}

func (s *student) string() string { // pass by reference (嚴格來說還是 value, 只是傳入的是指標)
	return fmt.Sprintf("name: %s, age: %d", s.name, s.age)
}

func main() {
	var person1 person
	person1.age, person1.name = 20, "Aery"

	person2 := person{"Rion", 18} // 按順序賦值

	person3 := person{age: 19, name: "Yuma"} // kv 賦值

	person4 := new(person) // 指標
	person4.age, person4.name = 21, "Yuka"

	utils.WrapPrint("showOrder", func() {
		showOrder := func(p1, p2 person) {
			older, diff := older(p1, p2)
			fmt.Printf("%v vs %v, %student1 is older by %d years\n", p1, p2, older.name, diff)
		}
		showOrder(person1, person2)
		showOrder(person1, person3)
		showOrder(person1, *person4) // 再加上 * 代表取值, 因為 showOrder 那邊接收的 person 型別, 所以要取值
	})

	utils.WrapPrint("showModifyStruct", func() {
		showModifyStruct := func(p person) {
			fmt.Printf("%v defore modifySlice\n", p)
			modifyStruct(p)
			fmt.Printf("%v after  modifySlice\n", p)
		}
		showModifyStruct(person1)
		showModifyStruct(person2)
		showModifyStruct(person3)
		showModifyStruct(*person4)
	})

	student1 := student{person{"May", 9}, "haha-School", 3}
	student1.name += "!"        // 直接存取 person.name
	student1.person.name += "?" // 明確存取 person.name
	student1.age = 10           // 存取 student 內的 age
	student1.person.age = 11    // 存取 person 內的 age
	utils.WrapPrint("struct 的繼承", func() {
		fmt.Printf("student: %v\n", student1)
	})

	// func 也有同上 struct 繼承一樣 scope 概念
	utils.WrapPrint("func 的繼承", func() {
		fmt.Printf("person1.string()         : %s\n", person1.string())
		fmt.Printf("student1.string()        : %s\n", student1.string())
		fmt.Printf("student1.person.string() : %s\n", student1.person.string())
	})
}

func older(p1, p2 person) (person, int) {
	if p1.age > p2.age { // 比較 p1 和 p2 這兩個人的年齡
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func modifyStruct(p person) { // pass by value, 不會改到原本的值
	p.age += 1
	p.name += "!"
}
