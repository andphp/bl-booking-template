package system

import (
	"server/ods/modelExtend"
	"server/ods/sqlMapper"

	"github.com/andphp/go-gin/goby/echo"
)

type SysMenuGetter struct {
	sysMenuMapper *sqlMapper.SysMenuMapper
}

func (s *SysMenuGetter) CurrentMenu(roleId int) *echo.ErrorResult {
	sqlMapper := s.sysMenuMapper.GetCurrentAllMenu(roleId)
	var sysMenuList []modelExtend.CurrentMenuModel
	db := sqlMapper.Query().Find(&sysMenuList)
	if db.Error != nil {
		return echo.Result(nil, db.Error)
	}
	return echo.Result(sysMenuList, nil)
}
