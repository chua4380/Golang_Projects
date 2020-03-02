package process
import "fmt"
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 增加在线用户
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// 删除在线用户
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

// 返回当前所有在线的用户
func (this* UserMgr) GetAllOnlineUser() map[int]*UserProcess{
	return this.onlineUsers
}

// 根据userId返回对应的连接
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	// 待检测的方式
	up, ok := this.onlineUsers[userId]
	if !ok {   // 查找的用户当前不在线
		err = fmt.Errorf("用户%d不存在", userId)
		return
	}
	return
}