package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"kkblogs/blogs-api/model"
	"kkblogs/blogs-api/service/app"
	"log"
	"net/http"
)

type Message struct {
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsHandle @Title 启动服务
func (msg *Message) WsHandle(c *gin.Context) {
	//升级get请求为webSocket协议
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}
	client := &model.Client{Socket: conn}
	//defer conn.Close()
	go msg.reader(client)
	//m.send(conn)
	//for {
	//	//读取ws中的数据
	//	mt, message, err := ws.ReadMessage()
	//	if err != nil {
	//		c.Writer.Write([]byte(err.Error()))
	//		break
	//	}
	//	fmt.Println("client message " + fmt.Sprintf("%d,%s", mt, string(message)))
	//	//写入ws数据
	//	err = ws.WriteMessage(mt, message)
	//	if err != nil {
	//		break
	//	}
	//	fmt.Println("system message " + fmt.Sprintf("%d,%s", mt, string(message)))
	//}
}

// reader @Title 获取消息
func (*Message) reader(client *model.Client) error {
	// 读取可以客户端连接发送的数据
	var receiveMsg model.ReceiveMessage
	for {
		mt, message, err := client.Socket.ReadMessage()
		if err != nil {
			client.Socket.Close()
			app.MessageLogic.DeleteClient(client)
			log.Println(err)
			return err
		}
		fmt.Println(mt)
		err = json.Unmarshal(message, &receiveMsg)
		fmt.Println("client message: " + string(message))
		switch receiveMsg.Type {
		case model.SystemMessage:
			app.MessageLogic.RegisterClient(client, receiveMsg)
		}
	}
}

// send @Title 发送消息
func (*Message) send(client *model.Client) {
	err := client.Socket.WriteMessage(websocket.TextMessage, []byte("Hello"))
	if err != nil {
		log.Println(err)
		return
	}
}
