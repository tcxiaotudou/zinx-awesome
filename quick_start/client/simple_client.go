package main

import (
	"fmt"
	"github.com/aceld/zinx/znet"
	"io"
	"net"
	"time"
)

func main()  {
	fmt.Println("Client Test ... start")
	// 3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for n := 0; n < 3; n++ {
		dp := znet.NewDataPack()
		msg, _ := dp.Pack(znet.NewMsgPackage(0, []byte("Zinx Client Test Message")))
		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData)
		if err != nil {
			fmt.Println("read head error")
			break
		}

		// 将head字节流拆包到msg中
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err: ", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			// msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*znet.Message)

			fmt.Println("Msg: ID=", msg.GetMsgId(), ", len = ", msg.GetDataLen(), ", data = ", string(msg.GetData()))

			msg.Data = make([]byte, msg.GetDataLen())

			// 根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
			fmt.Println("==> Recv Msg: ID=", msg.GetMsgId(), ", len = ", msg.GetDataLen(), ", data = ", string(msg.GetData()))
		}
		time.Sleep(1*time.Second)
	}
}
