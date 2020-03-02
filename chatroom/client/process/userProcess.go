package process

import (
	"net"
	"fmt"
	"go_code/chatroom/common/message"
	"encoding/json"
	"encoding/binary"
	"go_code/chatroom/client/util"
	"os"
)

type UserProcess struct {

}

func (this *UserProcess) Register(userId int, userPwd string,
	userName string) (err error){
	// 1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	// 2.准备通过conn发送信息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 3.创建一个LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	// 4.registerMes
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 5.把data赋给mes.Data字段
	mes.Data = string(data)

	// 6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 创建一个Transfer实例
	tf := &util.Transfer{
		Conn: conn,
	}

	// 发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误")
	}

	mes, err = tf.ReadPkg()  // mes就是RegisterResMes
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}
	// 将mes的data部分反序列化成 LoginResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功！")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

func (this *UserProcess)Login(userId int, userPwd string) (err error){
	// 定协议
	//fmt.Printf("userId = %v, passWd = %v\n", userId, userPwd)
	//return nil

	// 1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	// 2.准备通过conn发送信息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 5.把data赋给mes.Data字段
	mes.Data = string(data)

	// 6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 延时关闭
	defer conn.Close()

	// 7.这个时候，data就是我们要发送的信息
	// 7.1先把data的长度发动给服务器
	// 先获取到data的长度 -> 表示长度的字节切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	// 发送长度
	_, err = conn.Write(buf[0:4])
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	fmt.Printf("发送信息的长度=%d 内容=%s\n", len(data), string(data))

	// 发送信息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}

	// 这里还需要处理服务端返回的消息
	//time.Sleep(time.Second*10)
	//fmt.Println("休眠了10秒")
	// 创建一个Transfer实例
	tf := &util.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}
	// 将mes的data部分反序列化成 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
		// 初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		// 显示当前用户的列表
		fmt.Println("当前在线用户列表如下：")
		for _, v :=  range loginResMes.UserIds {
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			// 完成客户端的onlineUsers初始化
			user := &message.User{
				UserId: v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		// 在这里还需要启动一个协程，
		// 改协程保持和服务器端的通讯
		// 如果服务器有数据推送给客户端
		// 则接收并显示在客户端的终端
		go serverProcessMes(conn)

		// 1. 显示我们的登录成功的菜单(循环显示)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println("登录失败")
		fmt.Println(loginResMes.Error)
	}
	return
}
