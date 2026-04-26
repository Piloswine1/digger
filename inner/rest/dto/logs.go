package dto

import "github.com/gin-gonic/gin"

type GetLogs struct {
	ID     string `uri:"id" binding:"required"`
	Limit  int    `form:"limit"`
	StdErr bool   `form:"stderr"`
}


func Clamp[T ~uint | ~uint8 | ~uint32 | ~uint64 | ~int | ~int32 | ~int64](n, min, max T) T {
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}


func (r *GetLogs) FromReq(c *gin.Context) error {
	if err := c.BindUri(r); err != nil {
		return err
	}

	if err := c.BindQuery(r); err != nil {
		return err
	}

	r.Limit = Clamp(r.Limit, 1, 10000)
	return nil
}
