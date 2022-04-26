package system

import (
	"fmt"
	"server/ods/modelExtend"
	"server/ods/sqlMapper"

	"github.com/andphp/go-gin/goby/echo"
)

type SysMenuGetter struct {
	sysMenuMapper *sqlMapper.SysMenuMapper
}

func (s *SysMenuGetter) CurrentMenu(roleId int) *echo.ErrorResult {
	currentAllMenu := s.sysMenuMapper.GetCurrentAllMenu(roleId)
	var sysMenuList []modelExtend.CurrentMenuTree
	db := currentAllMenu.Query().Find(&sysMenuList)
	if db.Error != nil {
		return echo.Result(nil, db.Error)
	}
	return echo.Result(sysMenuList, nil)
}

func (s *SysMenuGetter) List(p modelExtend.MenuRequestParams) *echo.ErrorResult {
	//var sysApiList []model.SysMenuModel
	var total int64

	// 统计sql
	sqlCountMapper := s.sysMenuMapper.SelectCount().GetList(p)
	fmt.Println(sqlCountMapper)
	db := sqlCountMapper.Query().Find(&total)
	if db.Error != nil {
		return echo.Result(nil, db.Error)
	}

	// 列表sql
	//sqlMapper := s.sysMenuMapper.Select(model.ColumnSysAPI).Page(p.Page,p.PageSize).GetApiList(p)
	//db = sqlMapper.Query().Find(&sysApiList)
	//if db.Error != nil {
	//	return echo.Result(nil, db.Error)
	//}
	return echo.Result(echo.PageResult{
		//List:     sysApiList,
		Total:    total,
		//Page:     p.Page,
		//PageSize: p.PageSize,
		Success:  true,
	}, nil)
}