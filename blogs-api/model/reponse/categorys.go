package reponse

type ListCategory struct {
	Id           int64  `json:"id"`
	Avatar       string `json:"avatar"`
	CategoryName string `json:"categoryName"`
	Description  string `json:"description"`
}
