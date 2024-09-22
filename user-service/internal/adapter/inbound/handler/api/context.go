package api

import "github.com/gin-gonic/gin"

type ApiContext struct {
	*gin.Context
}

func NewApiContext(ctx *gin.Context) ApiContext {
	return ApiContext{Context: ctx}
}

func (c *ApiContext) JSON(statusCode int, response interface{}) {
	c.Context.JSON(statusCode, response)
}
