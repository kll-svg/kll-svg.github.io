package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	AccountHolder string
	Balance       float64
	IsActive      bool
}

func (acc *BankAccount) GetAccountInfo() string {
	acc.IsActive = true
	status := "活跃"
	if !acc.IsActive {
		status = "未激活"
	}
	return fmt.Sprintf("账号:%s,账户持有人:%s,余额:%.2f,状态:%s", acc.AccountNumber, acc.AccountHolder, acc.Balance, status, acc.IsActive)
}

func (acc *BankAccount) Deposit(amount float64) error {
	if !acc.IsActive {
		fmt.Println("账号未激活，无法存款")
		return errors.New("账号未激活，无法存款")
	}
	if amount <= 0 {
		fmt.Println("存款金额必须大于0")
		return errors.New("存款金额必须大于0")
	}
	acc.Balance += amount
	fmt.Printf("存款%.2f成功，当前余额:%.2f\n", amount, acc.Balance)
	return nil
}

func (acc *BankAccount) Withdraw(amount float64) error {
	if !acc.IsActive {
		fmt.Println("<UNK>")
		return errors.New("<UNK>")
	}
	if amount <= 0 {
		fmt.Println("取款金额必须大于0")
		return errors.New("<UNK>0")
	}
	if amount > acc.Balance {
		fmt.Println("余额不足")
		return errors.New("余额不足")
	}
	acc.Balance -= amount
	fmt.Printf("取款%.2f成功，当前余额:%.2f\n", amount, acc.Balance)
	return nil
}
func (acc *BankAccount) Freeze(amount float64) {
	acc.IsActive = false
	fmt.Println("账号已冻结")
}
func (acc *BankAccount) Unfreeze() {
	acc.IsActive = true
	fmt.Println("账号已解冻")
}
func main() {
	account := BankAccount{
		AccountNumber: "123456",
		AccountHolder: "张三",
		Balance:       1000,
		IsActive:      true,
	}
	fmt.Println(account.GetAccountInfo())
	account.Deposit(500)
	fmt.Println(account.Balance)
	account.Withdraw(200)
	fmt.Println(account.Balance)
}
