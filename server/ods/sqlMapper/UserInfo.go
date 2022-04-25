package sqlMapper

import (
	"server/ods/modelExtend"

	"github.com/Masterminds/squirrel"

	"github.com/andphp/go-gin/goby/mapper"
	"github.com/andphp/go-gin/goby/utils"
)

type SysUserMapper struct {
	columns []string
}

func (s *SysUserMapper) Select(modelStruct interface{}) *SysUserMapper {
	columns, _ := utils.Struct2ColumnsValues(modelStruct, "from")
	m := &SysUserMapper{
		columns: columns,
	}
	return m
}

func (s *SysUserMapper) GetUserList() *mapper.SqlMapper {
	return mapper.Mapper(squirrel.Select("account_name", "nick_name").From("sys_users").OrderBy("id asc").Limit(10).ToSql())
}

func (s *SysUserMapper) GetUserOfLogin(l *modelExtend.Login) *mapper.SqlMapper {
	return mapper.Mapper(squirrel.Select(s.columns...).Where("account_name = ?", l.AccountName).Where(squirrel.Eq{"deleted_at": nil}).From("sys_users").Limit(1).ToSql())
}

func (s *SysUserMapper) GetUserByID(uid int64) *mapper.SqlMapper {
	u := modelExtend.CurrentUser{}
	columns, _ := utils.Struct2ColumnsValues(u, "form")
	return mapper.Mapper(squirrel.Select(columns...).Where("id = ?", uid).From("sys_users").OrderBy("id asc").ToSql())
}
