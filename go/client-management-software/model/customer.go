package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewCustomer(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (this *Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
		this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
}
