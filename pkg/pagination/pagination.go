package pagination

import (
	"eshort/pkg/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
)

var LIMIT uint64 = 20
var PAGE_KEY_NAME = "page"
var LIMIT_KEY_NAME = "limit"

// ViewData 同视图渲染的数据
type ViewData struct {
	Current int
	// 数据库的内容总数量
	TotalCount int64
	// 总页数
	TotalPage int
}

// Pagination 分页对象
type Pagination struct {
	Limit   uint64 `json:"limit"`   //本次查询筛选多少条
	Current uint64 `json:"current"` //当前第几页
	Total   uint64 `json:"total"`   //共多少条
	Pages   uint64 `json:"pages"`   //一共多少页
	db      *gorm.DB
}

// New 分页对象构建器
// r —— 用来获取分页的 URL 参数，默认是 page，可通过 config/pagination.go 修改
// db —— GORM 查询句柄，用以查询数据集和获取数据总数
// baseURL —— 用以分页链接
// PerPage —— 每页条数，传参为小于或者等于 0 时为默认值  10，可通过 config/pagination.go 修改
func New(c *gin.Context, db *gorm.DB, data interface{}) (Pagination, error) {
	// 实例对象
	p := Pagination{
		Limit:   getLimit(c),
		Current: getCurren(c),
		Total:   getCount(db),
	}
	p.Pages = getPages(p.Total, p.Limit)
	err := getData(db, data, p.Current, p.Limit)
	return p, err
}

func getLimit(c *gin.Context) uint64 {
	limit := types.StringToUint64(c.Query(LIMIT_KEY_NAME))
	if limit > 1000 {
		return 1000
	}
	if limit <= 0 {
		return LIMIT
	}
	return limit
}

func getCurren(c *gin.Context) uint64 {
	curren := types.StringToUint64(c.Query(PAGE_KEY_NAME))
	if curren <= 0 {
		return 1
	}
	return curren
}

func getCount(db *gorm.DB) uint64 {
	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return 0
	}
	c := uint64(count)
	return c
}

func getPages(total uint64, limit uint64) uint64 {
	if total == 0 {
		return 0
	}
	nums := int64(math.Ceil(float64(total) / float64(limit)))
	if nums == 0 {
		nums = 1
	}
	return uint64(nums)
}

func getData(db *gorm.DB, data interface{}, curren uint64, limit uint64) error {
	offset := 0
	if curren > 1 {
		offset = (int(curren - 1)) * int(limit)
	}
	err := db.Preload(clause.Associations).Limit(int(limit)).Offset(offset).Find(data).Error
	if err != nil {
		return err
	}
	return nil
}
