package base

import (
	"eshort/pkg/pagination"
)

type BaseController struct {
}

func (a *BaseController) NormalPage(list interface{}, pageData pagination.Pagination) map[string]interface{} {
	m := map[string]interface{}{
		"list": list,
		"page": pageData,
	}
	return m
}
