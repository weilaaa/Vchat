package process

import (
	"Vchat/common/message"
	"Vchat/vServer/model"
	"Vchat/vServer/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn     net.Conn
	UserID   int
	UserName string
}

func (this *UserProcess) NotifyOtherUserOnline(userID int, userName string) {
	for i, v := range userMgr.OnlineUser {
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

	fmt.Println("4.", notifyUserStatusMes.UserName)

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
			return
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
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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

	// fill container
	user, err := model.MyUserDao.Login(loginMes.UserID, loginMes.UserPW)
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
		// add user to userMgr
		userMgr.AddOnlineUser(this)
		fmt.Println(this.UserName, "has been added to usersOnline")

		// notify other users new user online
		this.NotifyOtherUserOnline(loginMes.UserID, user.UserName)

		// append usersId online to container
		for i := range userMgr.OnlineUser {
			loginResMes.UsersID = append(loginResMes.UsersID, i)
		}

		for _, up := range userMgr.OnlineUser {
			loginResMes.UsersName = append(loginResMes.UsersName, up.UserName)
		}

		fmt.Println(user, "login successfully")
	}

	fmt.Println("3.", loginResMes.UsersName)
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

	return

}
