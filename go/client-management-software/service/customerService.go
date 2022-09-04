package service

import (
	"alsoon.com/go_code/client-management-software/model"
)

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 0
	aCustomer := model.NewCustomer("HELLO", "Male", 27, "112", "xyz@xyz")
	customerService.Add(aCustomer)
	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

// return -1 if not found
func (this *CustomerService) FindById(id int) int {
	for i, v := range this.customers {
		if v.Id == id {
			return i
		}
	}
	return -1
}

func (this *CustomerService) DeleteById(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}

	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

func (this *CustomerService) GetInfoById(id int) *model.Customer {
	for _, v := range this.customers {
		if v.Id == id {
			return &v
		}
	}
	return nil
}

func (this *CustomerService) Update(index int, to model.Customer) bool {
	if to.Name != "" {
		this.customers[index].Name = to.Name
	}
	if to.Gender != "" {
		this.customers[index].Gender = to.Gender
	}
	if to.Age != 0 {
		this.customers[index].Age = to.Age
	}
	if to.Phone != "" {
		this.customers[index].Phone = to.Phone
	}
	if to.Email != "" {
		this.customers[index].Email = to.Email
	}

	return true
}
