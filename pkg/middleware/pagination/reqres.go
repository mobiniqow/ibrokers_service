package pagination

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const UrlScheme = "http"

type pagination struct {
	Next      string `json:"next_page"`
	Prev      string `json:"prev_page"`
	Count     int64  `json:"total_records"`
	TotalPage int    `json:"total_pages"`
	Total     int    `json:"current_page"`
}

type Response struct {
	Pagination pagination  `json:"pagination"`
	Items      interface{} `json:"data"`
}

func GenerateResponse(limit, total int, count int64, ctx *gin.Context, items interface{}) Response {
	hasPrev, hasNext := Counter(count, total, limit)
	var next string
	var pre string
	if hasNext {
		next = NextUrl(ctx)
	}
	if hasPrev {
		pre = PrevUrl(ctx)
	}
	_pagiantion := pagination{
		Count:     count,
		Total:     total,
		Next:      next,
		Prev:      pre,
		TotalPage: int(math.Ceil(float64(count) / float64(limit)))}
	return Response{
		Items:      items,
		Pagination: _pagiantion,
	}
}

func CurrentUrl(ctx *gin.Context) string {
	host := ctx.Request.Host
	fullURL := fmt.Sprintf("%s://%s%s", UrlScheme, host, ctx.Request.URL)
	return fullURL
}

func NextUrl(ctx *gin.Context) string {
	host := ctx.Request.Host
	currentURL := ctx.Request.URL
	currentPage := currentURL.Query().Get("page")
	var fullURL string
	if currentPage == "" {
		currentPage = "1"
	} else {
		currentPageInt, err := strconv.Atoi(currentPage)
		if err != nil {
			currentPage = "1"
		} else {
			currentURL.Query().Set("page", strconv.Itoa(currentPageInt))
			fullURL = fmt.Sprintf("%s://%s%s", UrlScheme, host, currentURL)
			fullURL = strings.Replace(fullURL, fmt.Sprintf("page=%d", currentPageInt), fmt.Sprintf("page=%d", currentPageInt+1), 1)

		}
	}
	return fullURL
}

func PrevUrl(ctx *gin.Context) string {
	host := ctx.Request.Host
	currentURL := ctx.Request.URL
	currentPage := currentURL.Query().Get("page")
	var fullURL string
	if currentPage == "" {
		currentPage = "1"
	} else {
		currentPageInt, err := strconv.Atoi(currentPage)
		if err != nil {
			currentPage = "1"
		} else {
			currentURL.Query().Set("page", strconv.Itoa(currentPageInt))
			fullURL = fmt.Sprintf("%s://%s%s", UrlScheme, host, currentURL)
			fullURL = strings.Replace(fullURL, fmt.Sprintf("page=%d", currentPageInt), fmt.Sprintf("page=%d", currentPageInt-1), 1)

		}
	}
	return fullURL
}

func Counter(count int64, page, limit int) (hasPrev bool, hasNext bool) {

	hasNext = count > int64(page*limit) && count != 0
	hasPrev = (page > 1) && count > int64(limit) && count != 0
	return hasPrev, hasNext
}
