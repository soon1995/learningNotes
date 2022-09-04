package process

import (
	"fmt"

	"alsoon.com/go_code/chatsystem/client/model"
	"alsoon.com/go_code/chatsystem/common/message"
)

var (
	onlineUser map[int]*message.User = make(map[int]*message.User, 10)
	CurUser    model.CurrUser
)

func showOnlineUsers() {
	fmt.Println("Current Online Users:")
	for id, _ := range onlineUser {
		fmt.Println("UserID:\t", id)
	}

}

func updateUserStatus(notifyStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineUser[notifyStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyStatusMes.UserId,
		}
	}

	user.UserStatus = notifyStatusMes.UserStatus
	onlineUser[notifyStatusMes.UserId] = user
}
