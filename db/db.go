package db

import (
	"math"

	"github.com/GoRustNet/xurl/conf"
	"github.com/GoRustNet/xurl/defs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

const (
	DefaultPageSize int = 30
)

func Init(cfg *conf.PgConfig) (err error) {
	db, err = sqlx.Open("postgres", cfg.Dsn)
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	return nil
}

type Modeler interface {
	defs.User | defs.Url | defs.UserUrlVisitLiteView | defs.UserUrlLiteView
}

type Pagination[T Modeler] struct {
	Page        int  `json:"page"`
	PageSize    int  `json:"page_size"`
	RecordTotal int  `json:"record_total"`
	PageTotal   int  `json:"page_total"`
	Data        []*T `json:"data,omitempty"`
}

func NewPagination[T Modeler](page, pageSize, recordTotal int, data []*T) *Pagination[T] {
	pageTotal := int(math.Ceil(float64(recordTotal) / float64(pageSize)))
	return &Pagination[T]{
		Page:        page,
		PageSize:    pageSize,
		RecordTotal: recordTotal,
		Data:        data,
		PageTotal:   pageTotal,
	}
}

func (p *Pagination[T]) LastPage() int {
	return p.PageTotal - 1
}
func (p *Pagination[T]) HasNext() bool {
	return p.Page < p.LastPage()
}
func (p *Pagination[T]) HasPrev() bool {
	return p.Page > 0
}
func (p *Pagination[T]) IsCurrent(page int) bool {
	return p.Page == page
}

func PaginateSpecifyCount[T Modeler](selc, countSelc *Select, page int, params ...interface{}) (*Pagination[T], error) {
	count, err := Count(selc.ToCount(), params...)
	if err != nil {
		return nil, err
	}
	var data []*T
	if err := db.Select(&data, selc.Build(), params...); err != nil {
		return nil, err
	}
	p := NewPagination[T](page, DefaultPageSize, count, data)
	return p, nil
}
func Paginate[T Modeler](selc *Select, page int, params ...interface{}) (*Pagination[T], error) {
	return PaginateSpecifyCount[T](selc, selc.ToCount(), page, params...)
}

func Count(selc *Select, params ...interface{}) (int, error) {
	sqlstr := selc.Build()
	var count int
	if err := db.Get(&count, sqlstr, params...); err != nil {
		return 0, err
	}
	return count, nil
}

func Exists(selc *Select, params ...interface{}) (bool, error) {
	c, err := Count(selc, params...)
	if err != nil {
		return false, err
	}
	return c > 0, nil
}
