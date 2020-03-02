package process

import (
	"go_code/chatroom/common/message"
	"net"
	"encoding/json"
	"fmt"
	"go_code/chatroom/server/util"
)

type SmsProcess struct {}

// 转发群聊消息
func (this *SmsProcess) SengGroupMes(mes *message.Message) {
	// 取出mes的内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 取到服务器端的onlineUsers
	onlineUsers := userMgr.GetAllOnlineUser()
	for id, up := range onlineUsers {
		// 需要过滤自己，
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUsers(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUsers(data[]byte, conn net.Conn) {
	transfer := util.Transfer{
		Conn: conn,
	}
	err := transfer.WritePkg(data)
	if err != nil {
		fmt.Println("转发群聊消息失败, err=", err)
	}
}