package schemas

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Order *string

const (
	ASC  string = "asc"
	DESC string = "desc"
)

type PageInfo struct {
	Page     *int    `json:"page" binding:"omitempty,gte=1"`
	PageSize *int    `json:"page_size" binding:"omitempty,gte=1,lte=1000"`
	Keyword  *string `json:"keyword" binding:"omitempty" `
	Order    *string `json:"order" binding:"omitempty,oneof=asc desc"`
	OrderBy  *string `json:"order_by" binding:"omitempty"`
}

func (r *PageInfo) LimitOffset() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Page == nil || r.PageSize == nil {
			return db.Offset(0).Limit(10)
		}
		return db.Offset((*r.Page - 1) * *r.PageSize).Limit(*r.PageSize)
	}
}

func (r *PageInfo) OrderByField() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.OrderBy == nil {
			return db
		}
		if r.Order != nil {
			return db.Order(clause.OrderByColumn{Column: clause.Column{Name: *r.OrderBy}, Desc: *r.Order == DESC})
		}
		return db
	}
}

func (r *PageInfo) Search(searchField []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Keyword == nil {
			return db
		}
		for _, field := range searchField {
			db = db.Where(fmt.Sprintf("%s like ?", field), fmt.Sprintf("%%%s%%", *r.Keyword))
		}
		return db
	}
}
