package reponse

type FollowFansRep struct {
	FollowCount   int `json:"followCount"`
	FansCount     int `json:"fansCount"`
	ArticlesCount int `json:"articlesCount"`
}
type FollowRep struct {
	Id []int `json:"id"`
}
