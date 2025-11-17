package main

import (
	"fmt"
)

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {
	//var cart []CartItem
	cart := make([]CartItem, 0)
	cart = append(cart, CartItem{
		Name:     "手机",
		Price:    1000.0,
		Quantity: 1,
	})
	cart = append(cart, CartItem{
		Name:     "电脑",
		Price:    2000.0,
		Quantity: 1,
	})
	//if len(cart) > 0 {
	//	cart = cart[1:]
	//}
	var totalPrice float64
	for _, item := range cart {

		totalPrice += item.Price * float64(item.Quantity)
	}
	fmt.Println(totalPrice)

}
