package router

import "github.com/gin-gonic/gin"

func Router(route *gin.Engine) *gin.Engine {
	// 授权群组
	auth := route.Group("/auth")
	AuthRoute(auth)
	// 校验群组
	code := route.Group("/code")
	CodeRoute(code)
	return route
}
