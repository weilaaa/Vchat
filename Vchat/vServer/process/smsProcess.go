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

	for id, up := range userMgr.OnlineUser {
		// filter yourself
		if id == smsMes.UserId {
			continue
		}
		this.SendMes2EchoOthers(data, up.Conn)
	}
}

func (this *SmsProcess) SendMes2EchoOthers(data []byte, conn net.Conn) {

	tf := utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePKG(data)
	if err != nil {
		fmt.Println("write data failed from SendMes2EchoOthers", err)
		return
	}

}
