package message

// stipulate message type const
const (
	LoginMesType            = "loginMes"
	LoginResMesType         = "loginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMesType"
	SmsMesP2PType           = "SmsMesP2P"
	BinTransferType         = "BinTransfer"
	FeedBackMesType			= "FeedBackMes"
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

type BinTransfer struct {
	Content    []byte `json:"content"`
	FileName   string `json:"file_name"`
	SenderId   int    `json:"sender_id"`
	ReceiverId int    `json:"receiver_id"`
	SenderName string `json:"sender_name"`
}

type SmsMesP2P struct {
	Content    string `json:"content"`
	SenderId   int    `json:"sender_id"`
	ReceiverId int    `json:"receiver_id"`
	SenderName string `json:"sender_name"`
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
	Code          int            `json:"code"`
	LoginUserName string         `json:"login_user_name"`
	Users         map[int]string `json:"users"`
	//UsersName     []string       `json:"users_name"`
	//UsersID       []int          `json:"users_id"`
	Error error `json:"error"`
}

// message of register
type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int   `json:"code"`
	Error error `json:"error"`
}

type FeedBackMes struct {
	Content string `json:"content"`
}
