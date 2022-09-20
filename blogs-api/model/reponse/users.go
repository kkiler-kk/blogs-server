package reponse

type UserRep struct {
	Id        int64  `json:"id"`
	Account   string `json:"account"`
	NikeName  string `json:"nikeName"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`
}
type UserInfoRep struct {
	Id        int64  `json:"id"`
	Account   string `json:"account"`
	NikeName  string `json:"nikeName"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`
	City      string `json:"city"`
	BirthDate int64  `json:"birthDate"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Gender    int    `json:"sex"`
}
