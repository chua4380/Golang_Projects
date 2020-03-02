package process

import (
	"fmt"
	"os"
	"net"
	"go_code/chatroom/client/util"
	"go_code/chatroom/common/message"
	"encoding/json"
)

// 显示登陆成功后的界面...
func ShowMenu()  {
	fmt.Println("------恭喜xxx登录成功------")
	fmt.Println("-------1. 显示在线用户列表-------")
	fmt.Println("-------2. 发送消息--------")
	fmt.Println("-------3. 消息列表--------")
	fmt.Println("-------4. 退出系统--------")
	fmt.Println("请选择(1-4):")
	var key int
	fmt.Scanf("%d\n", &key)
	var content string

	smsProcess := &SmsProcess{}
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		//fmt.Println("发送消息")
		fmt.Println("你相对大家说什么")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确,请重新输入")
	}
}

func serverProcessMes(conn net.Conn)  {
	// 创建一个Transger实例， 不停地读取服务器发送的消息
	tf := &util.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg fail, err=", err)
			return
		}
		// 读取到消息，进行下一步处理逻辑
		switch mes.Type {
		case message.NotifyUserStatusMesType:  // 通知有人上线
			// 1. 取出NotifyUserStatusMesType
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			// 2. 把这个用户状态的信息 保存到客户map[int]User中
			updateUserStatus(&notifyUserStatusMes)
			// 3. 处理
		case message.SmsMesType:  // 有人群发消息
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器返回了未知的消息类型")

		}
		//fmt.Println("mes=", mes)
	}
}