package main

import (
	"fmt"
	"io"
	"net"

	"alsoon.com/go_code/chatsystem/common/message"
	"alsoon.com/go_code/chatsystem/server/process"
	"alsoon.com/go_code/chatsystem/server/utils"
)

type Processor struct {
	Conn net.Conn
}

// ServerProcessMes to process the mes according the mes type
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		sp := &process.SmsProcess{}
		err = sp.ServerProcessGroupSend(mes)
	default:
		fmt.Println("No Message Type matched")
	}
	return
}

func (this *Processor) processing() (err error) {
	for {

		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		mes, err := tf.ReadPkg()

		if err != nil {
			if err == io.EOF {
				fmt.Printf("client %v quit\n", this.Conn.RemoteAddr())
				return err
			} else {
				// fmt.Println("readPkg err: ", err)
				return err
			}
		}

		// Process message
		err = this.serverProcessMes(&mes)
	}
}
