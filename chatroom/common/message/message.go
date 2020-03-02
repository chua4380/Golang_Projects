package message

type Message struct {
	Type string `json:"type"`// 消息类型
	Data string `json:"data"`
}

const (
	LoginMesType	        = "LoginMes"
	LoginResMesType	        = "LoginResMes"
	RegisterMesType         = "LoginResMes"
	RegisterResMesType      = "RegisterMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

// 定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type LoginMes struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code int // 返回状态码 500 表示该用户为注册 200表示登陆成功
	UserIds []int    // 增加字段，保存在线用户id
	Error string // 返回错误信息
}

type RegisterMes struct {
	User User `json:"user"`// 类型就是User结构体
}

type RegisterResMes struct {
	Code int `json:"code"`  // 返回状态码：200表示注册成功， 400表示该用户已经被占用
	Error string `json:"error"`
}

// 为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

// 增加了一个smsMes
type SmsMes struct {
	Content string `json:"content"`
	User   // 匿名结构体
}