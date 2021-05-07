package main

import (
	"fmt"
	accounts "github.com/riley817/learngo/account"
)

func main() {
	account := accounts.NewAccount("riley")
	account.Deposit(3000)
	err := account.Withdraw(1000)

	if err != nil {
		fmt.Println()
	}
	fmt.Println(account)

	fmt.Println(account.Balance(), account.Owner())

}
