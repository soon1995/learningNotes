package process

import (
	"encoding/json"
	"fmt"
	"net"

	"alsoon.com/go_code/chatsystem/common/message"
	"alsoon.com/go_code/chatsystem/server/utils"
)

type SmsProcess struct {
}

func (this *SmsProcess) ServerProcessGroupSend(mes *message.Message) (err error) {

	// get data
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("unmarshal failed: ", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("Marshal failed: ", err)
		return
	}

	for id, up := range userMgr.OnlineUsers {
		if id != smsMes.UserId {
			this.SendToOne(data, up.Conn)
		}
	}
	return
}

func (this *SmsProcess) SendToOne(data []byte, conn net.Conn) {

	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("fail to redirect msg")
	}
}
