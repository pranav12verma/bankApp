package main

import (
	bank "bankApp/bank"
	"fmt"
)

func main() {
loop:
	for {
		fmt.Println("Welcome to the Bank .")
		fmt.Println()
		fmt.Println("1. Create Your Account")
		fmt.Println("2. Withdraw Money")
		fmt.Println("3. Deposit Money")
		fmt.Println("4. History of Transaction")
		fmt.Println("5. Delete Account")
		fmt.Println("6. Exit")

		var choice int
		fmt.Print("Enter Your Choice   : ")

		fmt.Scanf("%d", &choice)
		fmt.Println()
		switch choice {
		case 1:
			var name string
			fmt.Print("Enter Your Name :  ")
			fmt.Scan(&name)
			fmt.Println()

			bank.CreateAccount(name)
		case 2:
			var id string
			var amount int
			fmt.Print("Enter Your Id  :")
			fmt.Scan(&id)
			fmt.Println()
			fmt.Print("Enter 	Amount  :")
			fmt.Scanf("%v", &amount)
			fmt.Println()
			bank.WithdrawMoney(id, amount)
		case 3:
			var id string
			var amount int
			fmt.Print("Enter Your Id  :")
			fmt.Scan(&id)
			fmt.Println()
			fmt.Print("Enter 	Amount  :")
			fmt.Scanf("%v", &amount)
			fmt.Println()
			bank.DepositMoney(id, amount)
		case 4:
			var name string
			fmt.Print("Enter Your Id  :s")
			fmt.Scan(&name)
			fmt.Println()
			bank.TransactionHistory(name)

		case 5:
			var name string
			fmt.Print("Enter Your Id : ")
			fmt.Scan(&name)
			fmt.Println()
			bank.DeleteAccount(name)
		case 6:
			fmt.Println("Thank You!!!")
			break loop
		default:
			fmt.Println("Invalid Choice Exiting")
			break loop
		}
	}
}
