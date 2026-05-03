package dto

import "github.com/gin-gonic/gin"

type ContainerID struct {
	ID string `uri:"id" binding:"required"`
}

func (r *ContainerID) FromReq(c *gin.Context) error {
	if err := c.BindUri(r); err != nil {
		return err
	}
	return nil
}


type ContainersReq struct {
	All bool `form:"all"`
}

func (r *ContainersReq) FromReq(c *gin.Context) error {
	if err := c.BindQuery(r); err != nil {
		return err
	}
	return nil
}
