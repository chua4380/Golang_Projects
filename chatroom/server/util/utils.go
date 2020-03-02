package util

import (
	"net"
	"go_code/chatroom/common/message"
	"fmt"
	"encoding/binary"
	"encoding/json"
)

// 将这些方法关联到结构体
type Transfer struct {
	Conn net.Conn
	Buf [8096]byte  // 传输时使用缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("读取客户端发送的数据...")
	//buf := make([]byte, 8096)
	// conn.Read在conn没有被关闭的情况下才会阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err = errors.New("read pkg head error")
		return
	}
	fmt.Println("读到的buf=", this.Buf[:4])
	// 将buf[:4]转成一个uint32函数
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	// 根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	// 先发送数据包的长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	// 发送长度
	_, err = this.Conn.Write(this.Buf[0:4])

	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	//fmt.Printf("发送信息的长度=%d 内容=%s\n", len(data), string(data))

	// 发送data本身
	n, err := this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	return
}