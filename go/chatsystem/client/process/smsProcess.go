package process

import (
	"encoding/json"
	"fmt"

	"alsoon.com/go_code/chatsystem/client/utils"
	"alsoon.com/go_code/chatsystem/common/message"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroup(content string) (err error) {

	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.User = CurUser.User
	smsMes.Content = content

	b, err := json.Marshal(smsMes)
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
		Conn: CurUser.Conn,
	}

	// 2. send message
	err = tf.WritePkg(b)
	if err != nil {
		fmt.Println("write pkg failed: ", err)
		return
	}

	return
}
