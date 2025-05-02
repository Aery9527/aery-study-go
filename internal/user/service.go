package user

import "aery-study-go/internal/order"

func FindAllUser() {
	// 調用另外一個 package 的函數
	order.FindAllOrder()
	//order.fixOrder() // 不能調用 private func, 雖然 package 相同但在不同 folder, 實際上為不同 scope
}

func fixUser() {
}
