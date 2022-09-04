package main

import (
	"fmt"

	"alsoon.com/go_code/client-management-software/model"
	"alsoon.com/go_code/client-management-software/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService

	custId     int
	custName   string
	custGender string
	custAge    int
	custPhone  string
	custEmail  string
}

func (this *customerView) mainMenu() {
	for {
		fmt.Println("\n------------------客户信息管理软件------------------")
		fmt.Println("                  1 添加客户")
		fmt.Println("                  2 修改客户")
		fmt.Println("                  3 删除客户")
		fmt.Println("                  4 客户列表")
		fmt.Println("                  5 退    出")
		fmt.Print("请选择(1 - 5): ")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.addClient()
		case "2":
			this.updateClient()
		case "3":
			this.deleteClient()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("Invalid input (1-5)")
		}

		if !this.loop {
			break
		}
	}

	fmt.Println("Exiting...")
}

func (this *customerView) addClient() {
	fmt.Println("------------------添加客户------------------")
	fmt.Print("姓名：")
	custName := ""
	fmt.Scanln(&custName)
	fmt.Print("性别：")
	custGender := ""
	fmt.Scanln(&custGender)
	fmt.Print("年龄：")
	custAge := 0
	fmt.Scanln(&custAge)
	fmt.Print("电话：")
	custPhone := ""
	fmt.Scanln(&custPhone)
	fmt.Print("电邮：")
	custEmail := ""
	fmt.Scanln(&custEmail)

	newCustomer := model.NewCustomer(custName, custGender, custAge, custPhone, custEmail)
	if this.customerService.Add(newCustomer) {
		fmt.Println("Customer added")
	} else {
		fmt.Println("Customer failed")
	}
}

func (this *customerView) updateClient() {
	this.list()
	fmt.Println()
	fmt.Println("------------------修改客户------------------")
	fmt.Print("请选择等待修改客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Print("确认是否修改(Y/N)：")
	//这里可以加入一个循环判断，直到用户输入y或者n，才退出...
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		index := this.customerService.FindById(id)
		//调用customerService的Delete方法
		if index != -1 {
			customer := this.customerService.GetInfoById(id)
			fmt.Printf("姓名（%v：）", customer.Name)
			name := ""
			fmt.Scanln(&name)
			fmt.Printf("性别（%v）：", customer.Gender)
			gender := ""
			fmt.Scanln(&gender)
			fmt.Printf("年龄（%v）：", customer.Age)
			age := 0
			fmt.Scanln(&age)
			fmt.Printf("电话（%v）：", customer.Phone)
			phone := ""
			fmt.Scanln(&phone)
			fmt.Printf("邮箱（%v）：", customer.Email)
			email := ""
			fmt.Scanln(&email)
			customer2 := model.NewCustomer(name, gender, age, phone, email)
			this.customerService.Update(index, customer2)
			fmt.Println("------------------修改完成------------------")
		} else {
			fmt.Println("------------------修改失败，输入的id号不存在")
		}
	}
}

func (this *customerView) deleteClient() {
	fmt.Println("------------------删除客户------------------")
	fmt.Print("请选择等待删除客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}

	fmt.Print("Confirmed? (Y/N)：")
	//这里可以加入一个循环判断，直到用户输入y或者n，才退出...
	choice := ""
	fmt.Scanln(&choice)

	if choice == "y" || choice == "Y" {

		if this.customerService.DeleteById(id) {
			fmt.Printf("Customer #%v deleted\n", id)
		} else {
			fmt.Printf("No customer #%v found\n", id)
		}
	}
}

func (this *customerView) list() {
	fmt.Println("----------------------客户列表----------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

	cv := this.customerService.List()

	for _, v := range cv {
		fmt.Println(v.GetInfo())
	}
}

func (this *customerView) exit() {
	fmt.Print("确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Print("你的输入有误，确认是否退出(Y/N): ")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}

func main() {

	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
