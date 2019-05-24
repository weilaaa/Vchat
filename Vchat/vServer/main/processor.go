package main

import (
	"Vchat/common/message"
	"Vchat/vServer/process"
	"Vchat/vServer/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) Process() {
	defer this.Conn.Close()
	// read conn on loop
	// keep curUserId to record current user
	var curUserId *int
	a := -1
	curUserId = &a
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		mes, err := tf.ReadPKG()

		if err != nil {
			// if connection closed by client
			if err == io.EOF {
				userID := *curUserId
				if userID == -1{
					break
				}
				fmt.Printf("%d has offline\n", userID)
				userName := process.UserMGR.OnlineUser[userID].UserName
				userProcess := process.UserProcess{
					Conn:     this.Conn,
					UserName: userName,
					UserID:   userID,
				}
				delete(process.UserMGR.OnlineUser, userID)
				userProcess.NotifyMeOffline(userID, userName)
				break
			} else {
				fmt.Println("got message struct failed", err)
				break
			}
		}
		fmt.Println("got message data", mes.Data)

		err = this.ServerProcessMes(&mes, curUserId)
		if err != nil {
			fmt.Println("choose branch failed")
			return
		}
	}
	fmt.Println("waiting for next connection")
	fmt.Println()
}

// process the message came from client
func (this *Processor) ServerProcessMes(mes *message.Message, curUserId *int) (err error) {
	switch mes.Type {
	// process login business
	case message.LoginMesType:
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err := up.ServerProcessLogin(mes, curUserId)
		if err != nil {
			fmt.Println("server process login failed", err)
			return err
		}
	// process register business
	case message.RegisterMesType:
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
		if err != nil {
			fmt.Println("server process register failed", err)
		}
		return err
	// process message sending business
	case message.SmsMesType:
		smsMes := process.SmsProcess{}
		smsMes.SendGroupMes(mes)
	case message.SmsMesP2PType:
		smsMes := process.SmsProcess{}
		smsMes.SendP2PMes(mes)
	case message.BinTransferType:
		smsMes := process.SmsProcess{}
		smsMes.BinFileMes(mes)
	default:
		fmt.Println("unrecognized type")
		return err
	}

	return
}
