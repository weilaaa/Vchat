package process

import (
	"Vchat/common/message"
	"Vchat/vClient/utils"
	"encoding/json"
	"fmt"
)

type Smsprocess struct {
}

// send message
func (this *Smsprocess) SendGroupSms(content string) (err error){
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = curUser.UserId
	smsMes.UserStatus = curUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("smsMes marshal failed", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes from sendGroupSms failed", err)
		return
	}

	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}

	err = tf.WritePKG(data)
	if err!=nil{
		fmt.Println("write message from sendGroupSms failed",err)
		return
	}

	return
}
