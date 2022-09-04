package model

import (
	"net"

	"alsoon.com/go_code/chatsystem/common/message"
)

type CurrUser struct {
	message.User
	Conn net.Conn
}
