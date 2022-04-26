package sqlMapper

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/andphp/go-gin/goby/mapper"
	"server/ods/modelExtend"
)

type SysMenuMapper struct{
	*BasicMapper
}

func (*SysMenuMapper) GetCurrentAllMenu(roleId int) *mapper.SqlMapper {
	return mapper.Mapper(squirrel.Select("*").RightJoin("sys_menus on sys_menu_role_relations.menu_id = sys_menus.id").Where("sys_menu_role_relations.role_id = ?", roleId).Where(squirrel.Eq{"sys_menus.deleted_at": nil}).From("sys_menu_role_relations").ToSql())
}


func (s *SysMenuMapper) GetList(param modelExtend.MenuRequestParams) *mapper.SqlMapper {
	s.squirrelQuery = squirrel.Select(s.columns...).
		Where(squirrel.Eq{"deleted_at": nil}).From("sys_apis")


	if param.Name != "" {
		//squirrelQuery = squirrelQuery.Where("name LIKE ?", fmt.Sprint("%", param.Name, "%"))
		s.squirrelQuery = s.squirrelQuery.Where(squirrel.Like{"name": fmt.Sprint("%", param.Name, "%")})
	}

	if s.pagination == true {
		s.squirrelQuery = s.squirrelQuery.Limit(s.limit).Offset(s.offset)
	}
	return mapper.Mapper(s.squirrelQuery.ToSql())
}