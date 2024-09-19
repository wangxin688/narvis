package schemas

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Order *string

const (
	ASC  string = "asc"
	DESC string = "desc"
)

type PageInfo struct {
	Page     *int    `form:"page" binding:"omitempty,gte=1"`
	PageSize *int    `form:"pageSize" binding:"omitempty,gte=1,lte=1000"`
	Keyword  *string `form:"keyword" binding:"omitempty" `
	Order    *string `form:"order,default=desc" binding:"omitempty,oneof=asc desc"`
	OrderBy  *string `form:"orderBy,default=createdAt" binding:"omitempty"`
}

func (r *PageInfo) Pagination() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Page == nil && r.PageSize == nil {
			return db.Offset(0).Limit(10)
		}
		if r.Page == nil {
			initPage := 1
			r.Page = &initPage
		}
		if r.PageSize == nil {
			initPageSize := 10
			r.PageSize = &initPageSize
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

func (r *PageInfo) IsSearchable() bool {
	return r.Keyword != nil && *r.Keyword != ""
}
