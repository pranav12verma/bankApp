package bank

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	name      string
	id        string
	t_history []Transaction
	amount    int
}
type Transaction struct {
	id        string
	wd        string
	timeStamp string
	amount    int
}

var accounts = make(map[string]*Account)

func CreateAccount(s string) {
	var id string = uuid.NewString()
	var acc Account
	acc.id = id
	acc.name = s
	acc.amount = 0
	acc.t_history = []Transaction{}
	accounts[id] = &acc
	fmt.Println("Account created sucessfully")
	fmt.Println("ID : ", id)
}

func DepositMoney(id string, amt int) {
	var acc *Account
	acc, ok := accounts[id]
	if ok {
		acc.amount = acc.amount + amt
		var t Transaction
		t.id = uuid.NewString()
		t.amount = amt
		t.timeStamp = time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
		t.wd = "deposit"
		acc.t_history = append(acc.t_history, t)
		fmt.Println("Amount deposited ", amt, " account ID : ", acc.id)
	} else {
		fmt.Println("Wrong ID : ", id)
	}
}

func WithdrawMoney(id string, amt int) {
	var acc *Account
	acc, ok := accounts[id]
	if ok {
		if acc.amount > amt {
			acc.amount = acc.amount - amt
			var t Transaction
			t.id = uuid.NewString()
			t.amount = amt
			t.timeStamp = time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
			t.wd = "withdraw"
			acc.t_history = append(acc.t_history, t)
			fmt.Println("Amount deposited ", amt, " account ID : ", acc.id)
		} else {
			fmt.Println("Insufficient balance")
		}
	} else {
		fmt.Println("Wrong ID : ", id)
	}
}

func TransactionHistory(id string) {
	var acc *Account
	acc, ok := accounts[id]
	if ok {
		fmt.Printf("|%4v|%37v|%7v|%5v|%10v|\n", "sno", "transactionId", "type", "amount", "Time")
		for i, t := range acc.t_history {
			fmt.Printf("|%4v|%37v|%7v|%5v|%10v|\n", i+1, t.id, t.wd, t.amount, t.timeStamp)
		}
	} else {
		fmt.Println("Wrong ID : ", id)
	}
}

func DeleteAccount(id string) {
	var acc *Account
	acc, ok := accounts[id]
	if ok {
		acc = nil
		delete(accounts, id)
		fmt.Println("account deleted succesfully")
		fmt.Println(acc)
	} else {
		fmt.Println("Wrong ID : ", id)
	}
}
