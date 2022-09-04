package process

import (
	"encoding/json"
	"fmt"
	"net"

	"alsoon.com/go_code/chatsystem/common/message"
	"alsoon.com/go_code/chatsystem/server/model"
	"alsoon.com/go_code/chatsystem/server/utils"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (this *UserProcess) NotifyOtherActiveUser(userId int) {

	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	notifyMes := message.NotifyUserStatusMes{
		UserId:     userId,
		UserStatus: message.ONLINE,
	}

	data, err := json.Marshal(notifyMes)
	if err != nil {
		fmt.Println("Marshal failed: ", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("Marshal failed: ", err)
		return
	}

	for id, up := range userMgr.GetAllOnlineUser() {
		if id != userId {
			this.NotifySingle(data, up.Conn)
		}
	}
}

func (this *UserProcess) NotifySingle(data []byte, conn net.Conn) (err error) {

	tr := &utils.Transfer{
		Conn: conn,
	}

	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("Failed to deliver to: ", tr.Conn.LocalAddr())
	}

	return
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("unmarshal failed: ", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	// 1. check user detail
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOT_FOUND {
			loginResMes.Code = 500
			loginResMes.Err = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Err = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Err = "Server fault"
		}
	} else {
		loginResMes.Code = 200
		this.UserId = user.UserId
		userMgr.AddOnlineUser(this)
		this.NotifyOtherActiveUser(this.UserId)
		for id, _ := range userMgr.OnlineUsers {
			loginResMes.UsersIds = append(loginResMes.UsersIds, id)
		}
		fmt.Println(user, "Login successfully")
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("Marshal failed: ", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("Marshal failed: ", err)
		return
	}

	// 2. send response to user after identifying user
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	return
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registMes)
	if err != nil {
		fmt.Println("unmarshal failed: ", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	var registResMes message.RegisterResMes

	// 1. check user detail
	err = model.MyUserDao.Register(&registMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registResMes.Code = 500
			registResMes.Err = err.Error()
		} else {
			registResMes.Code = 505
			registResMes.Err = "Server fault"
		}
	} else {
		registResMes.Code = 200
		fmt.Println("Register successfully")
	}

	data, err := json.Marshal(registResMes)
	if err != nil {
		fmt.Println("Marshal failed: ", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("Marshal failed: ", err)
		return
	}

	// 2. send response to user after identifying user
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	return
}
