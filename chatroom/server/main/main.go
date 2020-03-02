package main

import (
	"fmt"
	"net"
	"go_code/chatroom/server/model"
	"time"
)

// 处理和客户端的通讯
func run(conn net.Conn) {
	// 延时关闭
	defer conn.Close()
	//调用主控，创建一个process实例
	processor := &Processor{
		Conn: conn,
	}
	err := processor.main_process()
	if err != nil {
		fmt.Println("客户端和服务器通信协程错误, err=", err)
		return
	}
}

// 编写函数，完成对UserDao的初始化工作
func initUserDao() {
	// 注意初始化的顺序问题，先initPool, 后initUserDao
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	// 当服务器启动时，我们就初始化我们redis的连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
}

func main() {
	// 提示信息
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	// 监听成功
	for {
		fmt.Println("等待客户端来连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		// 连接成功,启动一个协程和客户端保持通信
		go run(conn)
	}
}
