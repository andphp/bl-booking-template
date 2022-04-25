package system

import (
	"fmt"
	"server/ods/model"
	"server/ods/modelExtend"
	mapper "server/ods/sqlMapper"

	"github.com/andphp/go-gin/goby/echo"
)

type SysUserGetterInterface interface {
	PageList() []*model.SysUserModel
}

type SysUserGetter struct {
	sysUserMapper *mapper.SysUserMapper
}

func (s *SysUserGetter) PageList() (users []*model.SysUserModel) {
	sqlMapper := s.sysUserMapper.GetUserList()
	sqlMapper.Query().Find(&users)
	return
}

func (s *SysUserGetter) ByAccountName(l *modelExtend.Login) *echo.ErrorResult {

	account := modelExtend.Login{}
	db := s.sysUserMapper.Select(account).GetUserOfLogin(l).Query().Find(&account)
	if db.Error != nil {
		return echo.Result(nil, db.Error)
	}
	if account.AccountName == "" {
		return echo.Result(nil, fmt.Errorf("用户不存在"))
	}
	return echo.Result(account, nil)
}

func (s *SysUserGetter) ByID(uid int64) *echo.ErrorResult {
	userInfo := modelExtend.CurrentUser{}
	db := s.sysUserMapper.GetUserByID(uid).Query().Find(&userInfo)
	if db.Error != nil {
		return echo.Result(nil, db.Error)
	}
	return echo.Result(userInfo, nil)
}
