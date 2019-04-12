package model

import (
	"Vchat/common/message"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var MyUserDao *UserDao

type UserDao struct {
	Pool *redis.Pool
}

// factory model of UserDao
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		Pool: pool,
	}
	return
}

func (this *UserDao) GetUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOEXIST
		}
		return
	}

	user = &User{}

	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("get user unmarshal failed")
		return
	}

	return
}

// judge if account exist in redis
func (this *UserDao) Login(userID int, userPW string) (user *User, err error) {

	//get a connection from pool
	conn := this.Pool.Get()
	defer conn.Close()

	user, err = this.GetUserById(conn, userID)
	if err != nil {
		return
	}

	//judge the statement of user
	if user.UserPW != userPW {
		err = ERROR_USER_NOEXIST
		return
	}

	return
}

// register if account doesn't exist
func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.Pool.Get()
	defer conn.Close()

	_, err = this.GetUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXIST
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("user marshal fialed")
		return
	}

	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("register failed")
		return
	}
	return
}
