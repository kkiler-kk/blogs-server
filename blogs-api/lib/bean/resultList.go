package bean

type ResultLists struct {
	List  interface{} `json:"data"`
	Total int64       `json:"total"`
}
