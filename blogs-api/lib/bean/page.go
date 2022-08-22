package bean

const (
	LIMIT = 6 // 默认分页
)

type Page struct {
	Page     int //`json:"page"`  // 分页
	PageSize int //`json:"limit"` // 分页大小
}

type pageConfig struct {
	minLimit     int
	maxLimit     int
	defaultLimit int
}

var pConfig = &pageConfig{}

func InitPage(minLimit, maxLimit, defaultLimit int) {
	pConfig.minLimit = minLimit
	pConfig.maxLimit = maxLimit
	pConfig.defaultLimit = defaultLimit
}

// @Title 获取页数
func (p *Page) GetPage() int {
	if p.Page > 0 {
		return p.Page
	}
	return 1
}

// @Title 获取分页大小
func (p *Page) GetLimit() int {
	if p.PageSize >= pConfig.minLimit && p.PageSize <= pConfig.maxLimit {
		return p.PageSize
	}
	return pConfig.defaultLimit
}

// @Title 获取开始坐标
func (p *Page) GetStart() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Page) HasPage(total int64) bool {
	return total > int64(p.GetStart())
}
