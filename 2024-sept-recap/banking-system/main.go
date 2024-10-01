package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID      string
	Balance int
}

type BankingSystem struct {
	Accounts map[string]*Account
}

func Constructor() *BankingSystem {
	return &BankingSystem{Accounts: make(map[string]*Account)}
}

func (bs *BankingSystem) CreateAccount(id string) bool {
	// check to make sure the account does not already exist
	// if not create the acct
	if _, isPresent := bs.Accounts[id]; isPresent {
		return false
	}
	bs.Accounts[id] = &Account{ID: id, Balance: 0}
	return true
}

func (bs *BankingSystem) Deposit(id string, num int) (int, error) {
	// check to make sure the acct exists
	// if so update the balance

	if _, isPresent := bs.Accounts[id]; !isPresent {
		return -1, errors.New("Account does not exist")
	}

	bs.Accounts[id].Balance += num

	return bs.Accounts[id].Balance, nil
}

func main() {
	// create a banking system that has an account creation and deposit functions
	bank := Constructor()

	newAccount1 := bank.CreateAccount("123456789")
	newAccount2 := bank.CreateAccount("123456789")

	deposit1, error1 := bank.Deposit("123456789", 100)
	deposit2, error2 := bank.Deposit("12345678", 100)

	fmt.Println(newAccount1)
	fmt.Println(newAccount2)

	fmt.Println(deposit1, error1)
	fmt.Println(deposit2, error2)
}
