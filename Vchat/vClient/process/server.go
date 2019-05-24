package process

import (
	"Vchat/common/message"
	"Vchat/vClient/utils"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

// display second menu after login
func showMenu(loginUserName string) {
	smsProcess := Smsprocess{}
	fmt.Printf("Welcome home,%s\n", loginUserName)
	for {
		fmt.Println("1.Users online")
		fmt.Println("2.Send group message")
		fmt.Println("3.Send P2P message")
		fmt.Println("4.binary file transfer")
		fmt.Println("5.out of system")

		reader := bufio.NewReader(os.Stdin)
		scanner := bufio.NewScanner(reader)

		var content string
		var receiverID int
		var fileName string
		var key int
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			displayOnlineUser()
		case 2:
			fmt.Println("say as free")
			for scanner.Scan() {
				content = scanner.Text()
				if content == "q1" {
					break
				}
				err := smsProcess.SendGroupSms(content)
				if err != nil {
					fmt.Println("send group sms failed", err)
				}
			}
			//fmt.Scanf("%s\n", &content)
		case 3:
			fmt.Println("choose a friend")
			fmt.Scanf("%d\n", &receiverID)
			fmt.Println("say something")
			//fmt.Scanf("%s\n", &content)
			for scanner.Scan() {
				content = scanner.Text()
				if content == "q1" {
					break
				}
				err := smsProcess.SendP2PSMs(content, receiverID)
				if err != nil {
					fmt.Println("send P2P sms failed", err)
				}
			}
		case 4:
			fmt.Println("choose a friend")
			fmt.Scanf("%d\n", &receiverID)
			fmt.Println("choose a file")
			//fmt.Scanf("s%\n", &fileName)
			for scanner.Scan() {
				fileName = scanner.Text()
				if fileName == "q1" {
					break
				}
				err := smsProcess.BinFileSms(fileName, receiverID)
				if err != nil {
					fmt.Println("binary file transfer failed", err)
				}
			}
		case 5:
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
			if err != nil {
				fmt.Println("notify message unmarshal failed")
			}

			updateUserStatus(&notifyUserStatus)
		case message.SmsMesType:
			displayGroupMes(&mes)
		case message.SmsMesP2PType:
			displayP2PMes(&mes)
		case message.BinTransferType:
			displayBinMes(&mes)
		case message.FeedBackMesType:
			displayFeedbackMes(&mes)
		default:
			fmt.Println("unrecognizable type")
		}
	}
}
