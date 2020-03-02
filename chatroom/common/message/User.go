package message

// 定义用户的结构体
// 用户信息的json字段
type User struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
	UserStatus int `json:"userStatus"`   // 用户状态，在线/离线
}
