package routes

import (
	. "github.com/gtioi/chatgpt-html/app/http/controllers"
	"github.com/gtioi/chatgpt-html/app/middlewares"
	"github.com/gin-gonic/gin"
)

var chatController = NewChatController()

// RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.Engine) {
	router.Use(middlewares.Cors())
	router.GET("/", chatController.Index)
	router.POST("/completion", chatController.Completion)
}
