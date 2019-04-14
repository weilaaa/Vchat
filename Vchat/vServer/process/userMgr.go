package process

import "fmt"

// define only global variable as userMgr
var UserMGR *UserMgr

// keep online user
type UserMgr struct {
	OnlineUser map[int]*UserProcess
}

func init() {
	UserMGR = &UserMgr{
		make(map[int]*UserProcess, 1024),
	}
}

// manage onlineUser list
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.OnlineUser[up.UserID] = up
}

func (this *UserMgr) DelOnlineUser(userID int) {
	delete(this.OnlineUser, userID)
}

func (this *UserMgr) GetOnlineUser(userID int) (up *UserProcess, err error) {
	up, ok := this.OnlineUser[userID]
	if !ok {
		err = fmt.Errorf("userID %d doesn't exist", userID)
		return
	}
	return
}

func (this *UserMgr) GetAllOnlineUser() (onlineUser map[int]*UserProcess) {
	return this.OnlineUser
}
