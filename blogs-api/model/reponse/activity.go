package reponse

type Articles struct {
	Id            int64        `json:"id"`
	Title         string       `json:"title"`
	Summary       string       `json:"summary"`
	CommentCounts uint         `json:"commentCounts"`
	ViewCount     uint         `json:"viewCount"`
	Weight        uint         `json:"weight"`
	CreateDate    string       `json:"createDate"`
	Author        string       `json:"author"`
	ActivityBody  ActivityBody `json:"body"`
	Tags          []ListTag    `json:"tags"`
	Category      Category     `json:"category"`
}
type Publish struct {
	Id int64 `json:"id"`
}
type HNArticles struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type ArticlesView struct {
	Id            int64        `json:"id"`
	Title         string       `json:"title"`
	Summary       string       `json:"summary"`
	CommentCounts uint         `json:"commentCounts"`
	ViewCount     uint         `json:"viewCounts"`
	Weight        uint         `json:"weight"`
	CreateDate    string       `json:"createDate"`
	Author        SysUser      `json:"author"`
	ActivityBody  ActivityBody `json:"body"`
	Tags          []ListTag    `json:"tags"`
	Category      Category     `json:"category"`
}

type ActivitySearchRep struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}
