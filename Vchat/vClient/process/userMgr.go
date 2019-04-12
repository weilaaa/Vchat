package process

import (
	"Vchat/common/message"
	"Vchat/vClient/model"
	"fmt"
)

// client need keep a map to update user status
var onlineUser = make(map[string]*message.User, 10)
// declare curUser as global
var curUser model.CurUser

// update onlineUser map
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUser[notifyUserStatusMes.UserName]
	if !ok {
		user = &message.User{
			UserId:     notifyUserStatusMes.UserID,
			UserName:   notifyUserStatusMes.UserName,
			UserStatus: notifyUserStatusMes.UserID,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUser[notifyUserStatusMes.UserName] = user

	displayOnlineUser()
}

// display onlineUser
func displayOnlineUser() {
	fmt.Println("display users online")
	for _, v := range onlineUser {
		fmt.Println(v.UserName)
	}
}
