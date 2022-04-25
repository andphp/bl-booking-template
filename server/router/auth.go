package router

import (
	"server/controller"

	"github.com/gin-gonic/gin"
)

func Admin(g *gin.RouterGroup) {
	manage := g.Group("auth")
	manage.Handle("POST", "/sign/in", controller.ApiGroupApp.UnAuthorityGroup.Sign.SignIn)

	manage.GET("/menu", controller.ApiGroupApp.Login.CurrentMenu).
		GET("/user/info", controller.ApiGroupApp.Login.CurrentUser)

}
