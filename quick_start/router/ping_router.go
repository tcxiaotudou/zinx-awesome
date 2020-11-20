package router

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

// 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// ping handle
func (this *PingRouter) Handle (req ziface.IRequest) {
	// 先读取客户端的数据
	fmt.Println("recv from client : msgId = ", req.GetMsgID(), ", data = ", string(req.GetData()))

	// 再回写ping...ping...ping
	err := req.GetConnection().SendBuffMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}




