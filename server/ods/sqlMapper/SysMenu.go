package sqlMapper

import (
	"github.com/Masterminds/squirrel"
	"github.com/andphp/go-gin/goby/mapper"
)

type SysMenuMapper struct{}

func (*SysMenuMapper) GetCurrentAllMenu(roleId int) *mapper.SqlMapper {
	return mapper.Mapper(squirrel.Select("*").RightJoin("sys_menus on sys_menu_role_relations.menu_id = sys_menus.id").Where("sys_menu_role_relations.role_id = ?", roleId).Where(squirrel.Eq{"sys_menus.deleted_at": nil}).From("sys_menu_role_relations").ToSql())
}
