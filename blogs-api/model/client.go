package model

import (
	"github.com/gorilla/websocket"
	"sync"
)

const (
	SystemMessage = iota
	Register
	ToMessage
	Broadcast
)

// Client @Title 单个客户端
type Client struct {
	ID      string // 连接id (唯一标识符)
	UserId  int64  // 用户Id
	Socket  *websocket.Conn
	message ReceiveMessage
}

// ClientManager @Title 所有客户端的管理者
type ClientManager struct {
	sync.RWMutex
	// 所有客户端客户
	Clients map[int64]*Client
}

// ReceiveMessage @Title 客户端上传的数据
type ReceiveMessage struct {
	ID     int64       `json:"ID"` // 唯一标识符
	UserID int64       `json:"userID"`
	Type   int         `json:"type"`   // 消息类型
	Data   interface{} `json:"data"`   // 具体数据
	toUser int64       `json:"toUser"` // 发送给谁
}
