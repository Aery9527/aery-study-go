package main

import (
	"aery-study-go/internal/order"
	"aery-study-go/internal/user"
	// 預設使用最後一層 folder 當作 namespace, 有衝突則需要使用別名分開 namespace
	or "aery-study-go/internal/order/repository"
	ur "aery-study-go/internal/user/repository"
)

func main() {
	order.FindAllOrder()
	//order.fixOrder() // package private func

	user.FindAllUser()
	//user.fixUser() // package private func

	or.CreateOrder()
	or.FindOrder()
	or.UpdateOrder()
	or.DelectOrder()

	ur.CreateUser()
	ur.FindUser()
	ur.UpdateUser()
	ur.DelectUser()
}
