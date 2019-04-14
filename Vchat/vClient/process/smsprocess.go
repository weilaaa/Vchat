package process

import (
	"Vchat/common/message"
	"Vchat/vClient/utils"
	"encoding/json"
	"fmt"
)

type Smsprocess struct {
}

// send group message
func (this *Smsprocess) SendGroupSms(content string) (err error) {
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
	if err != nil {
		fmt.Println("write message from sendGroupSms failed", err)
		return
	}

	return
}

// send message peer to peer
func (this *Smsprocess) SendP2PSMs(content string, receiverID int) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesP2PType

	var smsMesP2P message.SmsMesP2P
	smsMesP2P.Content = content
	smsMesP2P.ReceiverId = receiverID
	smsMesP2P.SenderId = curUser.UserId

	data, err := json.Marshal(smsMesP2P)
	if err != nil {
		fmt.Println("smsMes from SendP2PSms failed", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes marshal failed from SendP2PSms", err)
		return
	}

	tf := utils.Transfer{
		Conn: curUser.Conn,
	}

	err = tf.WritePKG(data)
	if err != nil {
		fmt.Println("write data failed from SendP2PSms", err)
		return
	}

	return
}
