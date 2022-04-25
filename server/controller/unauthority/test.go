package unauthority

import (
	// "server/boot"

	"server/ods/getter"

	"github.com/andphp/go-gin/goby/echo"
	"github.com/gin-gonic/gin"
	// "go.uber.org/zap"
)

type Test struct {
}

func (t *Test) GetIndex(c *gin.Context) {
	// boot.GLOBAL_LOG.Warn("testsss", zap.String("testszap", "testszap"), zap.String("testszap1", "testszap1"))

	// sysUser := model.NewSysUserModel()
	// var allMenus []model.SysUserModel
	// goby.GOBY_DB.Table(sysUser.TableName()).Find(&allMenus)

	echo.OutputOkData(c)(getter.Getter.SysUser.PageList())
	// echo.OutputFailMsg(c)("密码错误")
}
