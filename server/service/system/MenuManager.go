package system

import (
	"server/ods/getter"
	"server/ods/modelExtend"

	"github.com/andphp/go-gin/goby/echo"
)

type MenuService struct{}

func (ain *MenuService) GetCurrentMenu(roleId int) *echo.ErrorResult {
	allMenus := getter.Getter.SysMenu.CurrentMenu(roleId).Unwrap()
	treeMap := map[int64][]modelExtend.CurrentMenuModel{}
	menus, _ := allMenus.([]modelExtend.CurrentMenuModel)
	for _, v := range menus {
		treeMap[v.ParentID] = append(treeMap[v.ParentID], v)
	}
	menuTree := treeMap[0]
	var err error = nil
	for i := 0; i < len(menuTree); i++ {
		err = ain.getChildrenList(&menuTree[i], treeMap)
		if err != nil {
			return echo.Result(nil, err)
		}
	}
	return echo.Result(menuTree, nil)
}

func (ain *MenuService) getChildrenList(menu *modelExtend.CurrentMenuModel, treeMap map[int64][]modelExtend.CurrentMenuModel) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = ain.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
