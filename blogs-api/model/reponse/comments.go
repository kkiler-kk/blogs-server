package reponse

type ListCommentsRep struct {
	Id         int64             `json:"id"`
	Author     SysUser           `json:"author"`
	Content    string            `json:"content"`
	Children   []ListCommentsRep `json:"childrens"`
	CreateDate string            `json:"createDate"`
	Level      uint              `json:"level"`
	ToUser     SysUser           `json:"toUser"`
}
