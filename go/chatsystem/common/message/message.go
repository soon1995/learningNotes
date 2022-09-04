package message

const (
	LoginMesType = iota
	LoginResMesType
	RegisterMesType
	RegisterResMesType
	NotifyUserStatusMesType
	SmsMesType
)

// user status
const (
	ONLINE = iota
	OFFLINE
	BUSY
)

type Message struct {
	Type int    `json:"type"` // Message Type: eg login
	Data string `json:"data"` // Message Detail
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code     int    `json:"code"` // 500: user not found, 200: success
	UsersIds []int  `json:"usersIds"`
	Err      string `json:"err"` // error message
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code int    `json:"code"` // 400: user exist, 200: success
	Err  string `json:"err"`  // error message
}

type NotifyUserStatusMes struct {
	UserId     int `json:"userId"`
	UserStatus int `json:"userStatus"`
}

type SmsMes struct {
	User
	Content string `json:"content"`
}
