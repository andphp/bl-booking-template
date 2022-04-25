package model


type SysPositionModel struct { 
    ID            int64        `gorm:"column:id" json:"id"`                      // 
    PositionName  string       `gorm:"column:position_name" json:"positionName"` // 职位名称
    DepartmentID  int64        `gorm:"column:department_id" json:"departmentID"` // 部门ID
    PositionCode  string       `gorm:"column:position_code" json:"positionCode"` // 部门编码
    ParentID      int64        `gorm:"column:parent_id" json:"parentID"`         // 
}

func (model *SysPositionModel) TableName() string {
	return "sys_positions"
}

var ColumnSysPosition = struct {
    ID                   string
	PositionName         string
	DepartmentID         string
	PositionCode         string
	ParentID             string
	
}{
    ID                  : "id",
    PositionName        : "position_name",
    DepartmentID        : "department_id",
    PositionCode        : "position_code",
    ParentID            : "parent_id",
    
}

type SysPositionModelOptions struct {
    apply func(*SysPositionModel)
}

func NewSysPositionModel(opts ...*SysPositionModelOptions) *SysPositionModel {
    m := &SysPositionModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysPosition_ID(id int64) *SysPositionModelOptions {
    return &SysPositionModelOptions{func(model *SysPositionModel) {
        model.ID = id
    }}
}

func SysPosition_PositionName(position_name string) *SysPositionModelOptions {
    return &SysPositionModelOptions{func(model *SysPositionModel) {
        model.PositionName = position_name
    }}
}

func SysPosition_DepartmentID(department_id int64) *SysPositionModelOptions {
    return &SysPositionModelOptions{func(model *SysPositionModel) {
        model.DepartmentID = department_id
    }}
}

func SysPosition_PositionCode(position_code string) *SysPositionModelOptions {
    return &SysPositionModelOptions{func(model *SysPositionModel) {
        model.PositionCode = position_code
    }}
}

func SysPosition_ParentID(parent_id int64) *SysPositionModelOptions {
    return &SysPositionModelOptions{func(model *SysPositionModel) {
        model.ParentID = parent_id
    }}
}


func(model *SysPositionModel) WithID(id int64) *SysPositionModel  {
    model.ID = id
    return model
}

func(model *SysPositionModel) WithPositionName(position_name string) *SysPositionModel  {
    model.PositionName = position_name
    return model
}

func(model *SysPositionModel) WithDepartmentID(department_id int64) *SysPositionModel  {
    model.DepartmentID = department_id
    return model
}

func(model *SysPositionModel) WithPositionCode(position_code string) *SysPositionModel  {
    model.PositionCode = position_code
    return model
}

func(model *SysPositionModel) WithParentID(parent_id int64) *SysPositionModel  {
    model.ParentID = parent_id
    return model
}









