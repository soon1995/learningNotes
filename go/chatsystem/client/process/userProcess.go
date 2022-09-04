package process

import (
	"encoding/json"
	"fmt"
	"net"

	"alsoon.com/go_code/chatsystem/client/utils"
	"alsoon.com/go_code/chatsystem/common/message"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	// 1. make connection
	c, err := net.Dial("tcp", "192.168.149.1:8888")
	if err != nil {
		fmt.Println("failed to connect - err:", err)
		return
	}
	defer c.Close()

	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	b, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	mes.Data = string(b)

	b, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	tf := &utils.Transfer{
		Conn: c,
	}

	// 2. send user detail
	err = tf.WritePkg(b)
	if err != nil {
		fmt.Println("write pkg failed: ", err)
		return
	}

	// 3. get response from server
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("read pkg failed: ", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("Unmarshal data failed: ", err)
		return
	}

	// 4. do operation
	if loginResMes.Code == 200 {
		CurUser.Conn = c
		CurUser.UserId = userId
		CurUser.UserStatus = message.ONLINE

		// login sucessfully
		fmt.Println("Online Users:")
		for _, v := range loginResMes.UsersIds {
			if v != userId {
				fmt.Printf("UserID: %v\n", v)
				user := &message.User{
					UserId:     v,
					UserStatus: message.ONLINE,
				}
				onlineUser[v] = user
			}
		}

		// go routine, start a server to receive info
		go serverProcess(c)

		// show menu
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Err)
	}

	return
}

func (this *UserProcess) Register(userId int,
	userPwd string, userName string) (err error) {

	// 1. make connection
	c, err := net.Dial("tcp", "192.168.149.1:8888")
	if err != nil {
		fmt.Println("failed to connect - err:", err)
		return
	}
	defer c.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType

	var registMes message.RegisterMes
	registMes.User.UserId = userId
	registMes.User.UserPwd = userPwd
	registMes.User.UserName = userName

	b, err := json.Marshal(registMes)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	mes.Data = string(b)

	b, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	tf := &utils.Transfer{
		Conn: c,
	}

	// 2. send user detail
	err = tf.WritePkg(b)
	if err != nil {
		fmt.Println("write pkg failed: ", err)
		return
	}

	// 3. get response from server
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("read pkg failed: ", err)
		return
	}

	var registRes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registRes)
	if err != nil {
		fmt.Println("Unmarshal data failed: ", err)
		return
	}

	// 4. do operation
	if registRes.Code == 200 {
		fmt.Println("Register succeesfully")
	} else {
		fmt.Println(registRes.Err)
	}

	return
}
