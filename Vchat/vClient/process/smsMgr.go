package process

import (
	"Vchat/common/message"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func displayGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("mes.Data unmarshal failed from displayGroupMes", err)
		return
	}

	h, m, s := time.Now().Clock()
	info := fmt.Sprintf("%d:%d:%d %s:%s", h, m, s, smsMes.UserName, smsMes.Content)
	fmt.Println(info)
}

func displayBinMes(mes *message.Message) {
	var binMes message.BinTransfer
	err := json.Unmarshal([]byte(mes.Data), &binMes)
	if err != nil {
		fmt.Println("mes.Data unmarshal failed from displayBinMes", err)
	}

	info := fmt.Sprintf("%s send you a file:%s", binMes.SenderName, binMes.FileName)
	fmt.Println(info)

	fileName := binMes.FileName
	err = ioutil.WriteFile(fileName, binMes.Content, os.ModePerm)
	if err != nil {
		fmt.Println("write file failed from displayBinMes", err)
	}
	

	/*cmd := exec.Command("open", "./file.jpg")
	err = cmd.Run()
	if err != nil {
		fmt.Println("open file failed", err)
	}*/
}

func displayP2PMes(mes *message.Message) {
	var smsMesP2P message.SmsMesP2P
	err := json.Unmarshal([]byte(mes.Data), &smsMesP2P)
	if err != nil {
		fmt.Println("mes.Data unmarshal failed from displayP2PMes", err)
	}

	h, m, s := time.Now().Clock()
	info := fmt.Sprintf("%d:%d:%d %s:%s", h, m, s, smsMesP2P.SenderName, smsMesP2P.Content)
	fmt.Println(info)
}

func displayFeedbackMes(mes *message.Message) {
	var smsMesFB message.FeedBackMes
	err := json.Unmarshal([]byte(mes.Data), &smsMesFB)
	if err != nil {
		fmt.Println("mes.Data unmarshal failed from displayFeedbackMes", err)
	}

	info := fmt.Sprintf("%s", smsMesFB.Content)
	fmt.Println(info)
}
