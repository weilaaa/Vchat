package process

import (
	"Vchat/common/message"
	"Vchat/vClient/utils"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

// display second menu after login
func showMenu() {
	smsProcess := Smsprocess{}
	for {
		fmt.Println("Welcome home")
		fmt.Println("1.Users online")
		fmt.Println("2.Send message")
		fmt.Println("3.Message list")
		fmt.Println("4.out of system")

		var content string
		var key int
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			displayOnlineUser()
		case 2:
			fmt.Println("say as free")
			fmt.Scanf("%s\n", &content)
			err := smsProcess.SendGroupSms(content)
			if err != nil {
				fmt.Println("send group sms failed", err)
			}
		case 3:
			fmt.Println("Message list")
		case 4:
			fmt.Println("Out of system")
			os.Exit(1)
		default:
			fmt.Println("Invalid operation from showMenu")
		}
	}
}

// keep connection to read message from server
func watcher(conn net.Conn) {
	tf := utils.Transfer{
		Conn: conn,
	}
	for {
		mes, err := tf.ReadPKG()
		if err != nil {
			if err == io.EOF {
				fmt.Println("watch finished")
				fmt.Println("connection closed by server")
				os.Exit(0)
			}
			fmt.Println("watch connection failed")
		}

		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatus message.NotifyUserStatusMes
			err := json.Unmarshal([]byte(mes.Data), &notifyUserStatus)
			fmt.Println("1.",notifyUserStatus.UserName)
			if err != nil {
				fmt.Println("notify message unmarshal failed")
			}

			updateUserStatus(&notifyUserStatus)
		case message.SmsMesType:
			displayGroupMes(&mes)

		default:
			fmt.Println("unrecognizable type")
		}
	}
}
