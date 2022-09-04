package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"alsoon.com/go_code/chatsystem/client/utils"
	"alsoon.com/go_code/chatsystem/common/message"
)

type Server struct {
}

func ShowMenu() {
	fmt.Println("-------恭喜登录成功---------")
	fmt.Println("-------1 显示在线用户列表---------")
	fmt.Println("-------2 发送消息---------")
	fmt.Println("-------3 信息列表---------")
	fmt.Println("-------4 退出系统---------")
	fmt.Println("请选择(1-4):")

	var key int
	var content string
	sms := &SmsProcess{}

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		// 显示在线用户列表
		showOnlineUsers()
	case 2:
		fmt.Println("想对大家说点什么:)")
		fmt.Scanln(&content)
		sms.SendGroup(content)
		fmt.Println("You said to all: ", content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统...")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确..")
	}
}

// start a server to listen
func serverProcess(conn net.Conn) {
	tf := utils.Transfer{
		Conn: conn,
	}

	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("Read failed :", err)
			return
		}

		messageProcess(mes)
	}
}

// process message by type
func messageProcess(mes message.Message) {
	switch mes.Type {
	case message.NotifyUserStatusMesType:
		var notifyUserStatusMes message.NotifyUserStatusMes
		err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
		if err != nil {
			fmt.Println("Unmarshal failed: ", err)
		}
		updateUserStatus(&notifyUserStatusMes)
		fmt.Println("User ID: ", notifyUserStatusMes.UserId, " logged in")
	case message.SmsMesType:
		var smsMes message.SmsMes
		err := json.Unmarshal([]byte(mes.Data), &smsMes)
		if err != nil {
			fmt.Println("Unmarshal failed: ", err)
			return
		}
		fmt.Println(smsMes.UserId, " said to all: ", smsMes.Content)
	default:
		fmt.Println("Message Type Unknown")
	}
}
