package utils

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GenericPaginated(slice interface{}, page, pageSize int) interface{} {
	s := reflect.ValueOf(slice)

	offset := (page - 1) * pageSize

	end := offset + pageSize

	if end > s.Len() {
		end = s.Len()
	}

	paginatedSlice := reflect.MakeSlice(reflect.SliceOf(s.Type().Elem()), end-offset, end-offset)

	for i := offset; i < end; i++ {
		paginatedSlice.Index(i - offset).Set(s.Index(i))
	}
	return paginatedSlice.Interface()
}

func ParserPageAndPageSize(context *gin.Context) (page, pageSize int) {
	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err = strconv.Atoi(context.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	return page, pageSize
}

func CalculateLinks(context *gin.Context, page, pageSize, totalItems int) (nextLink, prevLink string) {
	if page*pageSize < totalItems {
		nextLink = fmt.Sprintf("/feed?page=%d", page+1)
	}
	if page > 1 {
		prevLink = fmt.Sprintf("/feed?page=%d", page-1)
	}

	return nextLink, prevLink
}
