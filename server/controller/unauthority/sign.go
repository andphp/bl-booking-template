package unauthority

import (
	"fmt"
	"server/ods/modelExtend"
	"server/service"
	"server/tools"
	"time"

	"github.com/andphp/go-gin/goby"
	"github.com/andphp/go-gin/goby/echo"
	"github.com/andphp/go-gin/goby/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Sign struct {
}

func (s *Sign) SignIn(c *gin.Context) {

	// 参数验证
	var l modelExtend.Login
	echo.Result(c.ShouldBind(&l)).Unwrap(echo.PARAMS_ERROR)

	// 账号查询
	loginResult := service.Service.System.SysUser.Login(&l).Unwrap(echo.PARAMS_ERROR)

	accountLogin, _ := loginResult.(modelExtend.Login)

	// 验证密码
	if !utils.PasswordVerify(accountLogin.Password, l.Password) {
		echo.OutputFailMsg(c)("密码错误")
		return
	}

	// 更新数据库登录状态信息
	utils.Task(s.updateLoginInfo, nil, accountLogin)

	// 生产token
	jwt := tools.NewJWT()
	claims := jwt.CreateClaims(modelExtend.Session{
		ID:          accountLogin.ID,
		AccountName: accountLogin.AccountName,
		NickName:    accountLogin.NickName,
		RoleID:      accountLogin.RoleID,
	})
	token, err := jwt.CreateToken(claims)
	if err != nil {
		goby.GOBY_LOG.Error("获取token失败!", zap.Error(err))
		echo.OutputFailMsg(c)("获取token失败")
		return
	}
	echo.OutputOkData(c)(&map[string]interface{}{"accessToken": token, "expireToken": time.Now().Unix() + goby.GOBY_CONFIG.JWT.ExpiresTime})

}

func (s *Sign) SignUp(c *gin.Context) {
	echo.OutputFailMsg(c)("密码SignUp")
}

func (s *Sign) updateLoginInfo(p ...interface{}) {
	fmt.Println("更新数据库登录状态信息:")
}
