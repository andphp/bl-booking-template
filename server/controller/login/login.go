package login

import (
	"server/ods/getter"
	"server/service"
	"server/tools"

	"github.com/andphp/go-gin/goby/echo"
	"github.com/gin-gonic/gin"
)

type Login struct{}

func (ain *Login) CurrentMenu(c *gin.Context) {
	echo.OutputOkData(c)(service.Service.System.GetCurrentMenu(tools.GetUserRoleId(c)).Unwrap())
}

func (ain *Login) CurrentUser(c *gin.Context) {
	echo.OutputOkData(c)(getter.Getter.SysUser.ByID(tools.GetUserID(c)).Unwrap())
}
