package model


type SysUserDepartmentRelationModel struct { 
    UserID          int64   `gorm:"column:user_id" json:"userID"`                 // 用户ID
    DepartmentID    int64   `gorm:"column:department_id" json:"departmentID"`     // 角色ID
    DepartmentType  int     `gorm:"column:department_type" json:"departmentType"` // 
}

func (model *SysUserDepartmentRelationModel) TableName() string {
	return "sys_user_department_relations"
}

var ColumnSysUserDepartmentRelation = struct {
    UserID               string
	DepartmentID         string
	DepartmentType       string
	
}{
    UserID              : "user_id",
    DepartmentID        : "department_id",
    DepartmentType      : "department_type",
    
}

type SysUserDepartmentRelationModelOptions struct {
    apply func(*SysUserDepartmentRelationModel)
}

func NewSysUserDepartmentRelationModel(opts ...*SysUserDepartmentRelationModelOptions) *SysUserDepartmentRelationModel {
    m := &SysUserDepartmentRelationModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysUserDepartmentRelation_UserID(user_id int64) *SysUserDepartmentRelationModelOptions {
    return &SysUserDepartmentRelationModelOptions{func(model *SysUserDepartmentRelationModel) {
        model.UserID = user_id
    }}
}

func SysUserDepartmentRelation_DepartmentID(department_id int64) *SysUserDepartmentRelationModelOptions {
    return &SysUserDepartmentRelationModelOptions{func(model *SysUserDepartmentRelationModel) {
        model.DepartmentID = department_id
    }}
}

func SysUserDepartmentRelation_DepartmentType(department_type int) *SysUserDepartmentRelationModelOptions {
    return &SysUserDepartmentRelationModelOptions{func(model *SysUserDepartmentRelationModel) {
        model.DepartmentType = department_type
    }}
}


func(model *SysUserDepartmentRelationModel) WithUserID(user_id int64) *SysUserDepartmentRelationModel  {
    model.UserID = user_id
    return model
}

func(model *SysUserDepartmentRelationModel) WithDepartmentID(department_id int64) *SysUserDepartmentRelationModel  {
    model.DepartmentID = department_id
    return model
}

func(model *SysUserDepartmentRelationModel) WithDepartmentType(department_type int) *SysUserDepartmentRelationModel  {
    model.DepartmentType = department_type
    return model
}









