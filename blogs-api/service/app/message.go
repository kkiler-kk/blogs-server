package app

import "kkblogs/blogs-api/model"

var MessageLogic = &messageLogic{}

type messageLogic struct {
}

var manager = model.ClientManager{
	Clients: make(map[int64]*model.Client, 1000),
}

// RegisterClient @Title 添加添加client
func (*messageLogic) RegisterClient(client *model.Client, receiveMsg model.ReceiveMessage) {
	manager.Lock()
	defer manager.Unlock()
	client.UserId = receiveMsg.UserID
	manager.Clients[client.UserId] = client
}

// DeleteClient @Title 删除client
func (*messageLogic) DeleteClient(client *model.Client) {
	manager.Lock()
	defer manager.Unlock()
	if _, ok := manager.Clients[client.UserId]; ok {
		delete(manager.Clients, client.UserId)
	}
}
