package system

import (
	"server/ods/getter"
	"server/ods/modelExtend"

	"github.com/andphp/go-gin/goby/echo"
)

type SysUser struct{}

func (s *SysUser) Login(l *modelExtend.Login) *echo.ErrorResult {
	return getter.Getter.SysUser.ByAccountName(l)
}
