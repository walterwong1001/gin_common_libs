package endpoints

import (
	"context"
	"errors"
	"github.com/walterwong1001/gin_common_libs/page"
	"github.com/walterwong1001/gin_common_libs/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PaginationHelper[T, S any] interface {
	Pagination(ctx context.Context, p page.Paginator[T], filter S) error
}

func PathParamAsInt(c *gin.Context, key string) (uint64, error) {
	return strconv.ParseUint(c.Param(key), 10, 64)
}

func Abort(c *gin.Context, err error) {
	AbortWithMessage(c, err, "")
}

func AbortWithMessage(c *gin.Context, err error, text string) {
	if text != "" {
		err = errors.New(text)
	}
	_ = c.Error(err)
	c.Abort()
}

func CreateTime() int64 {
	return time.Now().UnixMilli()
}

func Render(c *gin.Context, data any) {
	c.Set(response.DATA_KEY, data)
}

// Paginator 分页器
func Paginator[T any](c *gin.Context) (page.Paginator[T], error) {
	current, err := PathParamAsInt(c, "current")
	if err != nil {
		Abort(c, err)
		return nil, err
	}
	size, err := PathParamAsInt(c, "size")
	if err != nil {
		Abort(c, err)
		return nil, err
	}
	return page.NewPagination[T](int(current), int(size)), nil
}

// QueryParams 查询参数
func QueryParams[T any](c *gin.Context) (T, error) {
	var filter T
	if err := c.ShouldBindQuery(&filter); err != nil {
		Abort(c, err)
		return filter, err
	}
	return filter, nil
}

// Pagination 抽象分页逻辑
func Pagination[T, F any](c *gin.Context, helper PaginationHelper[T, F], callback func(c *gin.Context, p page.Paginator[T]) any) {
	// 分页信息
	p, err := Paginator[T](c)
	if err != nil {
		Abort(c, err)
		return
	}
	// 过滤参数
	filter, err := QueryParams[F](c)
	if err != nil {
		Abort(c, err)
		return
	}

	err = helper.Pagination(c.Request.Context(), p, filter)
	if err != nil {
		Abort(c, err)
		return
	}
	if callback != nil {
		r := callback(c, p)
		Render(c, r)
		return
	}
	Render(c, p)
}
