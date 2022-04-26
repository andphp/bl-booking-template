package system

import (
	"github.com/andphp/go-gin/goby/echo"
	"github.com/gin-gonic/gin"
	"server/ods/getter"
	"server/ods/modelExtend"
)

type Menu struct {}

func (ain *Menu) List(c *gin.Context)  {
	// 参数验证
	var p modelExtend.MenuRequestParams
	echo.Result(c.ShouldBindQuery(&p)).Unwrap(echo.PARAMS_ERROR)
	echo.OutputOkData(c)(getter.Getter.SysMenu.List(p).Unwrap())
}