package process

import "errors"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	OnlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		OnlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.OnlineUsers[up.UserId] = up
}

func (this *UserMgr) RemoveOnlineUser(userId int) {
	delete(this.OnlineUsers, userId)
}

func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.OnlineUsers
}

func (this *UserMgr) GetOneUserById(userId int) (up *UserProcess, err error) {
	up = this.OnlineUsers[userId]
	if up == nil {
		err = errors.New("User not found")
	}
	return
}
