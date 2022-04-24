package utils

import (
	"strconv"

	"github.com/Investly-id/common-go/v3/payload"
	"github.com/labstack/echo/v4"
)

func GetPagination(c echo.Context) *payload.PaginationRequest {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perPage, _ := strconv.Atoi(c.QueryParam("perpage"))

	if page == 0 || page < 1 {
		page = 1
	}

	if perPage == 0 || perPage < 1 {
		perPage = 1
	}

	var offset int

	if page == 1 {
		offset = 0
	} else {
		offset = page*perPage - 1
	}

	limit := perPage

	return &payload.PaginationRequest{
		Page:    int64(page),
		PerPage: int64(perPage),
		Limit:   int64(limit),
		Offset:  int64(offset),
	}
}

func CalculateMetaPagination(totalData int64, p *payload.PaginationRequest) *payload.Pagination {

	lastPage := (totalData / p.PerPage)
	if lastPage == 0 {
		lastPage = 1
	}

	isLoadMore := p.Page != lastPage

	if p.Page > lastPage {
		isLoadMore = false
	}

	return &payload.Pagination{
		PerPage:     p.PerPage,
		CurrentPage: p.Page,
		IsLoadMore:  isLoadMore,
		LastPage:    lastPage,
	}
}
