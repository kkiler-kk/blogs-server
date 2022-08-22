package reponse

type LoginUser struct {
	Id        int64  `json:"id"`
	Account   string `json:"account"`
	NikeName  string `json:"nikeName"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`
}
