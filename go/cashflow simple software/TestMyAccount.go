package main

import (
	"fmt"
	"strings"
)

type FamilyAccount struct {
	key      string
	loop     bool
	balance  float64
	money    float64
	note     string
	detail   string
	isRecord bool
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:      "",
		loop:     true,
		balance:  10000,
		money:    0,
		note:     "",
		detail:   "Type\tbalance\t\tIncome\t\tNote",
		isRecord: false,
	}
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("--------------- Cash  Flow  Detail ---------------")
	if this.isRecord {
		fmt.Println(this.detail)
	} else {
		fmt.Println("No records found")
	}
}

func (this *FamilyAccount) newIncome() {
	fmt.Print("Income: ")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Print("Note: ")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\nIncome\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.isRecord = true
}

func (this *FamilyAccount) newExpense() {
	fmt.Print("Expense: ")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("Not sufficient balance")
		return
	}
	this.balance -= this.money
	fmt.Print("Note: ")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\nExpense\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.isRecord = true
}

func (this *FamilyAccount) exit() {
	fmt.Println("Confirm? (Y/N): ")
	input := ""
	for {
		fmt.Scanln(&input)
		if strings.ToLower(input) == "y" || strings.ToLower(input) == "n" {
			break
		}
		fmt.Print("wrong input..  (Y/N): ")
	}
	if strings.ToLower(input) == "y" {
		this.loop = false
	}
}

func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n--------------- Cash Flow Software ---------------")
		fmt.Println("                1. Detail")
		fmt.Println("                2. Record Income")
		fmt.Println("                3. Record Expense")
		fmt.Println("                4. Exit")
		fmt.Print("Please select (1-4): ")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.newIncome()
		case "3":
			this.newExpense()
		case "4":
			this.exit()
		default:
			fmt.Println("Invalid input, please enter (1-4)")
		}

		if !this.loop {
			break
		}
	}

	fmt.Println("Exiting...")
}

func main() {

	var application = NewFamilyAccount()
	application.MainMenu()

}
