package utils

import "fmt"

type Pagination struct {
	Page    int    `json:"page" form:"page"`
	Limit   int    `json:"limit" form:"limit"`
	Offset  int    `json:"offset" form:"offset"`
	Search  string `json:"search" form:"search"`
	OrderBy string `json:"order_by" form:"order"`
	Sort    string `json:"sort" form:"sort"`
}

func Paginate(p Pagination) (query string, cond string) {
	query += " "
	if p.Limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", p.Limit)
	}
	if p.Offset != 0 {
		query += fmt.Sprintf(" OFFSET %d", p.Offset)
	}
	if p.OrderBy != "" {
		query += fmt.Sprintf(" OFFSET %s", p.OrderBy)
	}
	if p.Sort != "" {
		query += fmt.Sprintf(" OFFSET %s", p.Sort)
	}
	if p.Search != "" {
		cond += " name like :search"
	}
	return
}
