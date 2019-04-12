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
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		mes, err := tf.ReadPKG()
		if err != nil {
			if err == io.EOF {
				fmt.Println("got message successfully")
				break
			} else {
				fmt.Println("got message struct failed", err)
				break
			}
		}
		fmt.Println("got message data", mes.Data)

		err = this.ServerProcessMes(&mes)
		if err != nil {
			fmt.Println("choose bunch failed")
			return
		}
	}
	fmt.Println("waiting for next connection")
	fmt.Println()
}

// process the message came from client
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	// process login business
	case message.LoginMesType:
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err := up.ServerProcessLogin(mes)
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
		return
	// process message sending business
	case message.SmsMesType:
		smsMes := process.SmsProcess{}
		smsMes.SendGroupMes(mes)
	default:
		fmt.Println("unrecognized type")
		return err
	}

	return err
}
