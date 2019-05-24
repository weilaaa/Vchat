package process

import (
	"Vchat/common/message"
	"Vchat/vServer/model"
	"Vchat/vServer/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"
)

type UserProcess struct {
	Conn     net.Conn
	UserID   int
	UserName string
}

func (this *UserProcess) NotifyOthers(userId int, userName string) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserID = userId
	notifyUserStatusMes.UserName = userName
	notifyUserStatusMes.Status = message.UserOffline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("notifyUserStatusMes marshal failed")
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("message.message form notification failed")
		return
	}

	tf := utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePKG(data)
	if err != nil {
		fmt.Println("write data from notification failed")
		return
	}
}

func (this *UserProcess) NotifyMeOffline(userID int, userName string) {
	for _, v := range UserMGR.OnlineUser {
		v.NotifyOthers(userID, userName)
	}
}

func (this *UserProcess) NotifyOtherUserOnline(userID int, userName string) {
	for i, v := range UserMGR.OnlineUser {
		if i == userID {
			continue
		}
		// push message
		v.NotifyMeOnline(userID, userName)
	}

}

func (this *UserProcess) NotifyMeOnline(userID int, userName string) {

	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserID = userID
	notifyUserStatusMes.UserName = userName
	notifyUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("notifyUserStatusMes marshal failed")
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("message.message form notification failed")
		return
	}

	tf := utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePKG(data)
	if err != nil {
		fmt.Println("write data from notification failed")
		return
	}

}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		err = errors.New("mes.Data unmarshal failed")
		return
	}

	// create a ship
	var resMes message.Message
	// set tag
	resMes.Type = message.RegisterResMesType

	// create a container
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXIST {
			registerResMes.Code = 505
			err.Error()
		}
	} else {
		registerResMes.Code = 200
		fmt.Println("register successfully")
	}

	// package container
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("registerResMes marshal failed", err)
		return
	}

	// loading container to the ship
	resMes.Data = string(data)

	// ready to ship off
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("resMes marshal failed", err)
		return
	}

	// crate transfer instant
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePKG(data)
	if err != nil {
		fmt.Println("send pkg failed", err)
		return
	}

	return
}

// here we need transfer the pointer of mes, cause we want to modify it
func (this *UserProcess) ServerProcessLogin(mes *message.Message, curUserId *int) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		err = errors.New("mes.Data unmarshal failed")
		return
	}

	// create a ship
	var resMes message.Message
	// set tag
	resMes.Type = message.LoginResMesType

	// create a container
	var loginResMes message.LoginResMes
	loginResMes.Users = make(map[int]string)

	// fill container
	user, err := model.MyUserDao.Login(loginMes.UserID, loginMes.UserPW)
	flag := false
	// process error
	if err != nil {
		if err == model.ERROR_USER_NOEXIST {
			loginResMes.Code = 500
			//err.Error to get the content of err
			err.Error()
		} else if err == model.ERROR_USER_PW {
			loginResMes.Code = 403
			err.Error()
		} else {
			fmt.Println("userDao inside error")
		}

	} else {
		loginResMes.Code = 200
		this.UserID = loginMes.UserID
		this.UserName = user.UserName
		*curUserId = loginMes.UserID
		// add user to userMgr
		UserMGR.AddOnlineUser(this)
		fmt.Println(this.UserName, "has been added to usersOnline")

		// notify other users new user online
		this.NotifyOtherUserOnline(loginMes.UserID, user.UserName)

		// append usersId online to container
		for id, up := range UserMGR.OnlineUser {
			loginResMes.Users[id] = up.UserName
		}

		loginResMes.LoginUserName = user.UserName

		fmt.Println(user, "login successfully")
		flag = true
	}

	// package container
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("loginResMes marshal failed", err)
		return
	}

	// loading container to the ship
	resMes.Data = string(data)

	// ready to ship off
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("resMes marshal failed", err)
		return
	}

	// crate transfer instant
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePKG(data)
	if err != nil {
		fmt.Println("send pkg failed", err)
		return
	}

	if flag == true {
		//check if there is offline message
		b, data := model.MyUserDao.GetMesById(loginMes.UserID)
		if b == false {
			fmt.Println(loginMes.UserID, "has no offline messages")
		} else {
			//send offline message
			fmt.Println(loginMes.UserID, "has offline messages")
			for _, v := range data {
				smsMes := SmsProcess{}
				smsMes.SendMes([]byte(v), UserMGR.OnlineUser[loginMes.UserID].Conn)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}

	return

}
