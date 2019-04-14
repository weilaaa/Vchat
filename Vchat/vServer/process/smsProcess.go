package process

import (
	"Vchat/common/message"
	"Vchat/vServer/model"
	"Vchat/vServer/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

// if you want to store offline message, you should push message
// into database
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("smsMes from server failed", err)
		return
	}

	conn := model.MyUserDao.Pool.Get()
	defer conn.Close()
	user, err := model.MyUserDao.GetUserById(conn, smsMes.UserId)
	if err != nil {
		fmt.Println("get user by id failed", err)
		return
	}
	smsMes.UserName = user.UserName

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("smsMes marshal failed from SendGroupMes", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes marshal failed from SendGroupMes", err)
	}

	for id, up := range UserMGR.OnlineUser {
		// filter yourself
		if id == smsMes.UserId {
			continue
		}
		this.SendMes(data, up.Conn)
	}
}

func (this *SmsProcess) SendMes(data []byte, conn net.Conn) {

	tf := utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePKG(data)
	if err != nil {
		fmt.Println("write data failed from SendMes2EchoOthers", err)
		return
	}

}

func (this *SmsProcess) SendP2PMes(mes *message.Message) {
	var smsMesP2P message.SmsMesP2P
	err := json.Unmarshal([]byte(mes.Data), &smsMesP2P)
	if err != nil {
		fmt.Println("mes.Data unmarshal failed from SendP2PMes", err)
	}

	conn := model.MyUserDao.Pool.Get()
	defer conn.Close()

	user, err := model.MyUserDao.GetUserById(conn, smsMesP2P.SenderId)
	if err != nil {
		fmt.Println("get user by id form SendP2PMes failed", err)
		return
	}
	smsMesP2P.SenderName = user.UserName

	data, err := json.Marshal(smsMesP2P)
	if err != nil {
		fmt.Println("smsMes marshaled failed from SendP2PMes", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes marshal failed from SendP2PMes", err)
		return
	}

	for id, up := range UserMGR.OnlineUser {
		if id == smsMesP2P.ReceiverId {
			this.SendMes(data, up.Conn)
		}
	}

}
