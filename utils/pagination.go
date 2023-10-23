package utils

import (
	"math"
)

type PaginationMeta struct {
	Count     int64            `json:"count"`
	Page      int              `json:"page"`
	Limit     int              `json:"limit"`
	TotalPage int              `json:"totalPage"`
	PrevPage  *int             `json:"prevPage"`
	NextPage  *int             `json:"nextPage"`
	From      int              `json:"from"`
	To        int              `json:"to"`
	Links     []PaginationLink `json:"links"`
}

type PaginationLink struct {
	Page     int  `json:"page"`
	IsActive bool `json:"isActive"`
}

func GeneratePaginationMeta(count int64, page int, limit int, restMeta ...interface{}) PaginationMeta {
	totalPage := int(math.Ceil(float64(count) / float64(limit)))

	var prevPage *int
	if page > 1 {
		prev := page - 1
		prevPage = &prev
	}

	var nextPage *int
	if page < totalPage {
		next := page + 1
		nextPage = &next
	}

	from := (page-1)*limit + 1
	to := page * limit
	if to > int(count) {
		to = int(count)
	}

	links := make([]PaginationLink, totalPage)
	for i := 1; i <= totalPage; i++ {
		links[i-1] = PaginationLink{
			Page:     i,
			IsActive: i == page,
		}
	}

	meta := PaginationMeta{
		Count:     count,
		Page:      page,
		Limit:     limit,
		TotalPage: totalPage,
		PrevPage:  prevPage,
		NextPage:  nextPage,
		From:      from,
		To:        to,
		Links:     links,
	}

	return meta
}
