package process

import (
	"go_code/chatroom/common/message"
	"fmt"
	"go_code/chatroom/client/model"
)

// 客户端维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser  // 在完成用户登陆后，完成读CurUser初始化
// 编写方法处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes * message.NotifyUserStatusMes) {
	// 适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {  // 原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
			UserStatus: notifyUserStatusMes.Status,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
}

// 在客户端显示当前在线的用户
func outputOnlineUser() {
	// 遍历onlieUsers
	fmt.Println("当前在线用户列表:")
	for id, _:= range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

