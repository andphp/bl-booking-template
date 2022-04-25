package middleware

import (
	"server/ods/model"

	"github.com/andphp/go-gin/goby/echo"
	"github.com/gin-gonic/gin"
)

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 参数验证
		var l model.SysUserModel
		echo.Result(c.ShouldBind(&l)).Unwrap(echo.PARAMS_ERROR)
		// do something
		c.Next()
	}
}
