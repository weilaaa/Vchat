package process

import (
	"Vchat/common/message"
	"Vchat/vClient/utils"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userID int, userPW, userNAME string) (err error) {
	// get user entering
	fmt.Println("ENTER YOUR ID")
	fmt.Scanf("%d\n", &userID)
	fmt.Println("ENTER YOUR PASSWORD")
	fmt.Scanf("%s\n", &userPW)
	fmt.Println("ENTER YOUR USERNAME")
	fmt.Scanf("%s\n", &userNAME)

	// generate a connection
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		log.Fatalln("client connected failed", err)
	}
	defer conn.Close()

	// crate a message instant
	var Mes message.Message
	Mes.Type = message.RegisterMesType

	// crate a loginMes instant
	var RegisterMes message.RegisterMes
	RegisterMes.User.UserId = userID
	RegisterMes.User.UserPW = userPW
	RegisterMes.User.UserName = userNAME

	// marshal the LoginMes
	data, err := json.Marshal(RegisterMes)
	if err != nil {
		// log.fatal will cause os.exit
		log.Fatalln("registerMes marshaled failed", err)
	}
	Mes.Data = string(data)

	// marshal the message
	data, err = json.Marshal(Mes)
	if err != nil {
		log.Fatalln("Message marshaled failed", err)
	}

	fmt.Println(string(data))

	// ship arrived port
	tf := utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePKG(data)
	if err != nil {
		log.Fatalln("data send failed")
	}

	resMes, err := tf.ReadPKG()
	if err != nil {
		fmt.Println("client read package failed")
		return
	}

	// unload container
	var RegisterResMes message.RegisterResMes

	err = json.Unmarshal([]byte(resMes.Data), &RegisterResMes)
	if err != nil {
		fmt.Println("resMes.Data unmarshal failed")
		return
	}

	// unpack container
	if RegisterResMes.Code == 200 {
		fmt.Println("register successfully")
		os.Exit(0)
	} else {
		fmt.Println("register failed from register")
		os.Exit(0)
	}
	return

}

func (this *UserProcess) Login(userID int, userPW string) (err error) {
	// get user entering
	fmt.Println("ENTER YOUR ID")
	fmt.Scanf("%d\n", &userID)
	fmt.Println("ENTER YOUR PASSWORD")
	fmt.Scanf("%s\n", &userPW)

	// generate a connection
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		log.Fatalln("client connected failed", err)
	}
	defer conn.Close()

	// crate a message instant
	var Mes message.Message
	Mes.Type = message.LoginMesType

	// crate a loginMes instant
	var LoginMes message.LoginMes
	LoginMes.UserID = userID
	LoginMes.UserPW = userPW

	// marshal the LoginMes
	data, err := json.Marshal(LoginMes)
	if err != nil {
		// log.fatal will cause os.exit
		log.Fatalln("loginMes marshaled failed", err)
	}
	Mes.Data = string(data)

	// marshal the message
	data, err = json.Marshal(Mes)
	if err != nil {
		log.Fatalln("Message marshaled failed", err)
	}

	fmt.Println(string(data))

	// ship arrived port
	tf := utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePKG(data)
	if err != nil {
		return
	}

	fmt.Println("client transfer data finished")
	fmt.Println("client starts to receive data")

	resMes, err := tf.ReadPKG()
	if err != nil {
		fmt.Println("client read package failed")
		return
	}

	// unload container
	var loginResMes message.LoginResMes

	err = json.Unmarshal([]byte(resMes.Data), &loginResMes)
	if err != nil {
		fmt.Println("resMes.Data unmarshal failed")
		return
	}

	// unpack container
	if loginResMes.Code == 200 {
		fmt.Println("login successfully")

		// initialize curUsr
		curUser.Conn = conn
		curUser.UserId = userID
		curUser.UserStatus = message.UserOnline

		// acquire usersID online
		fmt.Println("users online")
		for id, name := range loginResMes.Users {
			fmt.Printf("userID:%d\tuserName:%s\n", id, name)
			user := &message.User{
				UserName:   name,
				UserId:     id,
				UserStatus: message.UserOnline,
			}
			onlineUser[name] = user
		}

		// begin to watch connection between S&C
		go watcher(conn)
		showMenu(loginResMes.LoginUserName)
	} else if loginResMes.Code == 500 {
		fmt.Println("account doesn't exist, register now")
	} else {
		fmt.Println("invalid operation")
	}

	return nil
}
