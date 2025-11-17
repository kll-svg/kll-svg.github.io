package main

import "fmt"

type User struct {
	Id      int
	Name    string
	Balance float64
}

func updateBalanceValue(user User, amount float64) {
	user.Balance += amount
	// 打印更新后的余额
	fmt.Printf("用户 %s 的余额更新为: %.2f\n", user.Name, user.Balance)
}
func updateBalancePointer(user *User, amount float64) {
	user.Balance += amount
	// 打印更新后的余额
	fmt.Printf("用户 %s 的余额更新为: %.2f\n", user.Name, user.Balance)
}

func main() {
	user := User{
		Id:      1,
		Name:    "张三",
		Balance: 1000.0,
	}
	fmt.Println(user)
	updateBalanceValue(user, 1000.0)
	fmt.Println(user)
	updateBalancePointer(&user, 1000.0)
	fmt.Println(user)

}
