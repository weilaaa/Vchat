package model

import (
	"Vchat/common/message"
	"net"
)

// to tell server who you are
type CurUser struct {
	Conn net.Conn
	message.User
}
