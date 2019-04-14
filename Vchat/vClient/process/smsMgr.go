package process

import (
	"Vchat/common/message"
	"encoding/json"
	"fmt"
)

func displayGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("mes.Data unmarshal failed from displayGroupMes", err)
		return
	}

	info := fmt.Sprintf("%s:%s", smsMes.UserName, smsMes.Content)
	fmt.Println(info)
}

func displayP2PMes(mes *message.Message) {
	var smsMesP2P message.SmsMesP2P
	err := json.Unmarshal([]byte(mes.Data), &smsMesP2P)
	if err != nil {
		fmt.Println("mes.Data unmarshal failed from displayP2PMes", err)
	}

	info := fmt.Sprintf("%s:%s", smsMesP2P.SenderName, smsMesP2P.Content)
	fmt.Println(info)
}
