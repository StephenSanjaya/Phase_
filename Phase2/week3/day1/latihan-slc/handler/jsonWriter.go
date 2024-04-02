package handler

import (
	"Phase2/week3/day1/latihan-slc/utils"

	"github.com/gin-gonic/gin"
)

func ErrJsonWriter(c *gin.Context, err utils.ErrResponse, data any) {
	if data == nil && err.Details == nil {
		err.Details = struct{}{}
		c.JSON(err.Code, err)
		return
	}

	if err.Details != nil {
		c.JSON(err.Code, err)
		return
	}

	err.Details = data
	c.JSON(err.Code, err)
}
