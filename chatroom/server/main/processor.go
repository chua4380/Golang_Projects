package main

import (
	"net"
	"go_code/chatroom/common/message"
	"fmt"
	"go_code/chatroom/server/process"
	"io"
	"go_code/chatroom/server/util"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor)serverProcessMes(mes *message.Message) (err error) {
	// 看看是否能够接收到客户端发送的群发消息
	//fmt.Println("mes=", mes)
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		// 创建UserProcess实例
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册
		// 创建UserProcess实例
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		// 转发群聊消息
		smsProcess := &process.SmsProcess{}
		smsProcess.SengGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

//主控
func (this *Processor) main_process() (err error){
	// 循环读取客户端发送的信息
	// 创建一个Transfer 实例完成读报任务
	for {
		tf := &util.Transfer{
			Conn:this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出")
				return err
			} else {
				fmt.Println("readpkg err=", err)
				return err
			}
		}
		//fmt.Println("mes=", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
