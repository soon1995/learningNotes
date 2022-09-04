package model

import (
	"encoding/json"
	"fmt"

	"alsoon.com/go_code/chatsystem/common/message"
	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *userDao
)

type userDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) *userDao {
	return &userDao{
		pool: pool,
	}
}

func (this *userDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		// if data not found
		if err == redis.ErrNil {
			err = ERROR_USER_NOT_FOUND
		}
		return
	}

	user = &User{}

	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("unmarshal failed: ", err)
	}

	return
}

func (this *userDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//fmt.Println(userPwd, user.UserPwd)

	// check password
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}

	return
}

func (this *userDao) Register(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()

	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	data, err := json.Marshal(user)
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("redis hset fault", err)
		return
	}

	return
}
