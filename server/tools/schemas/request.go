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
	Page     *int    `form:"page,default=0" binding:"omitempty,gte=1"`
	PageSize *int    `form:"pageSize,default=10" binding:"omitempty,gte=1,lte=1000"`
	Keyword  *string `form:"keyword" binding:"omitempty" `
	Order    *string `form:"order,default=desc" binding:"omitempty,oneof=asc desc"`
	OrderBy  *string `form:"orderBy,default=createdAt" binding:"omitempty"`
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
		if r.OrderBy != nil && r.Order == nil {
			order := ASC
			r.Order = &order
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
