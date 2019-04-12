package message

// stipulate message type const
const (
	LoginMesType            = "loginMes"
	LoginResMesType         = "loginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMesType"
)

const (
	UserOnline = iota
	UserOffline
	UserBusyin
)

// pack the info when transfer data
// as a ship on the ocean
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// message of sending message
type SmsMes struct {
	Content string `json:"content"`
	User
}

// message of user status
type NotifyUserStatusMes struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Status   int    `json:"status"`
}

// message of login
type LoginMes struct {
	UserID int    `json:"user_id"`
	UserPW string `json:"user_pw"`
}

type LoginResMes struct {
	Code      int    `json:"code"`
	UsersName []string `json:"users_name"`
	UsersID   []int  `json:"users_id"`
	Error     error  `json:"error"`
}

// message of register
type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int   `json:"code"`
	Error error `json:"error"`
}
