package model

import (
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
)

// 我们在服务器启动时， 就初始化一个userDao实例
// 把它当作全局变量， 在需要和redis操作时， 就直接使用即可
var (
	MyUserDao *UserDao
)

// 完成对User结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 1. 根据用户id返回一个User实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	// 通过给定给id取redis查询
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		// 错误！
		if err == redis.ErrNil { // 表示在users中没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	// 这里我们需要把res反序列化成一个user实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// 完成登录的校验
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	// 先从UserDao取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	// 只是证明用户存在，此时还需要对密码进行验证
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

// 完成用户的注册
func (this *UserDao) Register(user *message.User) (err error) {
	// 先从UserDao取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	// 这时，说明id在redis中没有，完成注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	// 入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误，err=",err)
		return
	}
	return
}