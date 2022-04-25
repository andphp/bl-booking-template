package model


type SysRoleDepartmentRelationModel struct { 
    DepartmentID  int64   `gorm:"column:department_id" json:"departmentID"` // 
    RoleID        int64   `gorm:"column:role_id" json:"roleID"`             // 
    SystemCode    int     `gorm:"column:system_code" json:"systemCode"`     // 
}

func (model *SysRoleDepartmentRelationModel) TableName() string {
	return "sys_role_department_relations"
}

var ColumnSysRoleDepartmentRelation = struct {
    DepartmentID         string
	RoleID               string
	SystemCode           string
	
}{
    DepartmentID        : "department_id",
    RoleID              : "role_id",
    SystemCode          : "system_code",
    
}

type SysRoleDepartmentRelationModelOptions struct {
    apply func(*SysRoleDepartmentRelationModel)
}

func NewSysRoleDepartmentRelationModel(opts ...*SysRoleDepartmentRelationModelOptions) *SysRoleDepartmentRelationModel {
    m := &SysRoleDepartmentRelationModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysRoleDepartmentRelation_DepartmentID(department_id int64) *SysRoleDepartmentRelationModelOptions {
    return &SysRoleDepartmentRelationModelOptions{func(model *SysRoleDepartmentRelationModel) {
        model.DepartmentID = department_id
    }}
}

func SysRoleDepartmentRelation_RoleID(role_id int64) *SysRoleDepartmentRelationModelOptions {
    return &SysRoleDepartmentRelationModelOptions{func(model *SysRoleDepartmentRelationModel) {
        model.RoleID = role_id
    }}
}

func SysRoleDepartmentRelation_SystemCode(system_code int) *SysRoleDepartmentRelationModelOptions {
    return &SysRoleDepartmentRelationModelOptions{func(model *SysRoleDepartmentRelationModel) {
        model.SystemCode = system_code
    }}
}


func(model *SysRoleDepartmentRelationModel) WithDepartmentID(department_id int64) *SysRoleDepartmentRelationModel  {
    model.DepartmentID = department_id
    return model
}

func(model *SysRoleDepartmentRelationModel) WithRoleID(role_id int64) *SysRoleDepartmentRelationModel  {
    model.RoleID = role_id
    return model
}

func(model *SysRoleDepartmentRelationModel) WithSystemCode(system_code int) *SysRoleDepartmentRelationModel  {
    model.SystemCode = system_code
    return model
}









