package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID      string
	Balance int
}

type BankSystem struct {
	Accounts map[string]*Account
}

func NewBankSystem() *BankSystem {
	return &BankSystem{Accounts: make(map[string]*Account)}
}

func (bs *BankSystem) CreateAccount(id string) bool {
	// if the account does not already exist
	if _, isPresent := bs.Accounts[id]; isPresent {
		return false
	}

	bs.Accounts[id] = &Account{ID: id, Balance: 0}

	return true
}

func (bs *BankSystem) Deposit(id string, num int) (int, error) {
	// if the account acct exists, update the balance
	if _, isPresent := bs.Accounts[id]; !isPresent {
		return -1, errors.New("This account does not exist")
	}

	bs.Accounts[id].Balance += num

	return bs.Accounts[id].Balance, nil
}

func (bs *BankSystem) Transfer(transferFrom string, transferTo string, num int) (int, error) {
	if _, isPresent := bs.Accounts[transferFrom]; !isPresent {
		return -1, errors.New("The account you are attempting to transfer from DNE")
	}
	if _, isPresent := bs.Accounts[transferTo]; !isPresent {
		return -1, errors.New("The account you are attempting to transfer to DNE")
	}
	if bs.Accounts[transferFrom].Balance < num {
		return -1, errors.New("Insufficent Funds")
	}

	bs.Accounts[transferFrom].Balance -= num
	bs.Accounts[transferTo].Balance += num

	return bs.Accounts[transferFrom].Balance, nil
}

func main() {
	// create a banking system that has an account creation and deposit functions
	// next implment a transfer from one account to another

	bank := NewBankSystem()

	newAccount1 := bank.CreateAccount("123456789")
	newAccount2 := bank.CreateAccount("123456789")
	newAccount3 := bank.CreateAccount("12345678")

	fmt.Println(newAccount1)
	fmt.Println(newAccount2)
	fmt.Println(newAccount3)

	deposit1, error1 := bank.Deposit("123456789", 100)
	deposit2, error2 := bank.Deposit("1234567", 100)
	deposit3, error3 := bank.Deposit("123456789", 100)

	fmt.Println(deposit1, error1)
	fmt.Println(deposit2, error2)
	fmt.Println(deposit3, error3)

	transfer1, transferError1 := bank.Transfer("123456789", "12345678", 100)
	transfer2, transferError2 := bank.Transfer("123456", "12345678", 100)
	transfer3, transferError3 := bank.Transfer("123456789", "12345678", 200)

	fmt.Println(transfer1, transferError1)
	fmt.Println(transfer2, transferError2)
	fmt.Println(transfer3, transferError3)
}
