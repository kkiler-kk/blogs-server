package reponse

type SysUser struct {
	Id       int64  `json:"id"`
	Avatar   string `json:"avatar"`
	NickName string `json:"nickName"`
}
