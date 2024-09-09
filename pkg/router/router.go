package router

import "github.com/gin-gonic/gin"

type Router interface {
	RegisterRoutes(route *gin.RouterGroup)
}
