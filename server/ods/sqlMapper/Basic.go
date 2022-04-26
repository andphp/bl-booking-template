package sqlMapper

import (
	"github.com/Masterminds/squirrel"
	"github.com/andphp/go-gin/goby/mapper"
	"github.com/andphp/go-gin/goby/utils"
)

type BasicMapper struct {
	columns []string
	limit uint64
	offset uint64
	pagination bool
	squirrelQuery  squirrel.SelectBuilder
}


func (s *BasicMapper) Select(modelStruct interface{}) *BasicMapper  {
	columns, _ := utils.Struct2ColumnsValues(modelStruct,"from")
	s.columns= columns
	return s
}

func (s *BasicMapper) GetList(param interface{}) *mapper.SqlMapper {
	return mapper.Mapper(s.squirrelQuery.ToSql())
}

func (s *BasicMapper) SelectCount() *BasicMapper  {
	s.columns= []string{"count(1) as total "}
	return s
}

func (s *BasicMapper) SelectPage(current ,pageSize uint64) *BasicMapper  {
	s.limit = pageSize
	s.offset = (current - 1) * pageSize
	s.pagination = true
	return s
}
