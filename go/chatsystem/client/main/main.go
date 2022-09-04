package main

import (
	"fmt"
	"os"

	"alsoon.com/go_code/chatsystem/client/process"
)

var userId int
var userPwd string
var userName string

func main() {
	// 接收用户的选择
	var key int
	// 判断是否还继续显示菜单
	// var loop bool = true
	loop := true

	for loop {
		fmt.Println("----------------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			// 说明用户要登陆
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名称")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}

	}
	//更加用户的输入，显示新的提示信息
	if key == 1 {
	} else if key == 2 {
		fmt.Println("进行用户注册的逻辑....")
	}
}
