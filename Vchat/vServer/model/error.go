package model

import "errors"

// custom error
var (
	ERROR_USER_NOEXIST = errors.New("user id doesn't exist")
	ERROR_USER_EXIST   = errors.New("user already exist")
	ERROR_USER_PW      = errors.New("wrong password ")
	CONTENT_USER_NOEXIST = "user doesn't exist, please enter again"
)
