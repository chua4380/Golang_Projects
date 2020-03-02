package process

import (
	"net"
	"go_code/chatroom/common/message"
	"encoding/json"
	"fmt"
	"go_code/chatroom/server/util"
	"go_code/chatroom/server/model"
)

type UserProcess struct {
	Conn net.Conn
	UserId int   // 表明该conn是哪个用户的
}

// 通知所有在线用户有新用户加入
func (this *UserProcess) NotifyOtherOnlineuser(userId int) {
	// 遍历onlineUsers，然后逐个发送NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		// 开始通知
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	// 组装NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	// 将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 将序列化后的NotifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)

	// 对mes再次序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 发送，创建一个transfer实例
	tf := &util.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error){
	// 1.先从mes中取出mes.data, 并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err=", err)
		return
	}

	// 1.先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 2.再声明一个 LoginResMes
	var loginResMes message.LoginResMes

	// 3.连接redis数据库进行数据验证
	// 1)使用MyUserDao到redis
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			// 返回具体的错误信息
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {  // redis验证客户端登陆成功
		loginResMes.Code = 200
		// 将登录成功的用户id赋给this
		this.UserId = loginMes.UserId
		// 这里用户登录成功，我们就该把登录成功的用户放到userMgr中
		userMgr.AddOnlineUser(this)
		// 通知其他的在线用户， 我上线了
		this.NotifyOtherOnlineuser(this.UserId)
		// 将当前在线用户的id 放入到loginResMes.UsersId中
		// 遍历 userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Println(user.UserId, "登录成功")
	}
	/*
	// 判断用户是否合法
	if loginMes.UserId == 112 && loginMes.UserPwd == "123456" {
		// 合法
		loginResMes.Code = 200
	} else {
		// 不合法
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用"
	}
	*/

	// 3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}

	// 4.将data赋给resMes
	resMes.Data = string(data)

	// 5.对resMes进行序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}

	// 6. 发送data, 将其封装到readPkg函数
	// 先创建一个Transfer实例
	tf := &util.Transfer{
		Conn: this.Conn,
	}
	tf.WritePkg(data)
	return
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 1.先从mes中取出mes.data, 并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err=", err)
		return
	}
	// 2.先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	// 3.连接redis数据库完成注册
	// 1)使用MyUserDao到redis
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 400
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生位置错误"
		}
	} else {
		registerResMes.Code = 200
	}

	// 3.将registerResMes序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}
	// 4.将data赋给resMes
	resMes.Data = string(data)

	// 5.对resMes进行序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}

	// 6. 发送data, 将其封装到readPkg函数
	// 先创建一个Transfer实例
	tf := &util.Transfer{
		Conn: this.Conn,
	}
	tf.WritePkg(data)
	return
}