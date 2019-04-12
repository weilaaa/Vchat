package message

type User struct {
	UserId     int    `json:"user_id"`
	UserPW     string `json:"user_pw"`
	UserName   string `json:"user_name"`
	UserStatus int    `json:"user_status"`
}
