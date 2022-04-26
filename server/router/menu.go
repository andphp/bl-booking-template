package router

import (
	"github.com/gin-gonic/gin"
	"server/controller"
)

func Menu(g *gin.RouterGroup) {
	manage := g.Group("menu")

	manage.GET("/list", controller.ApiGroupApp.System.Menu.List)

}
