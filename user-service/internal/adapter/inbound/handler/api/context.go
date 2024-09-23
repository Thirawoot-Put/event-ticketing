package api

import (
	"github.com/Thirawoot-Put/event-ticketing/user-service/internal/core/port"
	"github.com/gin-gonic/gin"
)

type ApiContext struct {
	*gin.Context
}

func NewApiContext(ctx *gin.Context) *ApiContext {
	return &ApiContext{Context: ctx}
}

func (c *ApiContext) JSON(statusCode int, response interface{}) {
	c.Context.JSON(statusCode, response)
}

func GinHandler(handler func(port.HTTPContext)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewApiContext(c))
	}
}
