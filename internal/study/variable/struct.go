package variable

import "fmt"

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

func ShowStruct() {
	var p1 person
	p1.age, p1.name = 20, "Aery"

	p2 := person{"Rion", 18} // 按順序賦值

	p3 := person{age: 19, name: "Yuma"} // kv 賦值

	p4 := new(person) // 指標
	p4.age, p4.name = 21, "Yuka"

	showOrder := func(p1, p2 person) {
		older, diff := older(p1, p2)
		fmt.Printf("%v vs %v, %s is older by %d years\n", p1, p2, older.name, diff)
	}
	showOrder(p1, p2)
	showOrder(p1, p3)
	showOrder(p1, *p4)

	fmt.Println()
	showModifyStruct := func(p person) {
		fmt.Printf("%v defore modifySlice\n", p)
		modifyStruct(p)
		fmt.Printf("%v after  modifySlice\n", p)
	}
	showModifyStruct(p1)
	showModifyStruct(p2)
	showModifyStruct(p3)
	showModifyStruct(*p4) // 因為 func 那邊是用 person 型別, 所以是 pass by value, 因此不會改到原本的值

	fmt.Println()
	s := student{person{"May", 9}, "haha-School", 3}
	s.name += "!"        // 直接存取 person.name
	s.person.name += "?" // 明確存取 person.name
	s.age = 10           // 存取 student 內的 age
	s.person.age = 11    // 存取 person 內的 age
	fmt.Printf("student: %v\n", s)
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
