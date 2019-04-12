package process

import (
	"Vchat/common/message"
	"encoding/json"
	"fmt"
)

func displayGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	fmt.Println("2.", smsMes.UserName)
	if err != nil {
		fmt.Println("mes.Data unmarshal failed from displayGroupMes", err)
		return
	}

	info := fmt.Sprintf("%s:%s", smsMes.UserName, smsMes.Content)
	fmt.Println(info)
}
