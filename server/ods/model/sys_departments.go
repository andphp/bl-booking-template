package model


type SysDepartmentModel struct { 
    ID              int64        `gorm:"column:id" json:"id"`                          // 
    DepartmentName  string       `gorm:"column:department_name" json:"departmentName"` // 
    ParentID        int64        `gorm:"column:parent_id" json:"parentID"`             // 
    DepartmentCode  string       `gorm:"column:department_code" json:"departmentCode"` // 
    DepartmentType  int          `gorm:"column:department_type" json:"departmentType"` // 部门类型：跨组织（区域）
}

func (model *SysDepartmentModel) TableName() string {
	return "sys_departments"
}

var ColumnSysDepartment = struct {
    ID                   string
	DepartmentName       string
	ParentID             string
	DepartmentCode       string
	DepartmentType       string
	
}{
    ID                  : "id",
    DepartmentName      : "department_name",
    ParentID            : "parent_id",
    DepartmentCode      : "department_code",
    DepartmentType      : "department_type",
    
}

type SysDepartmentModelOptions struct {
    apply func(*SysDepartmentModel)
}

func NewSysDepartmentModel(opts ...*SysDepartmentModelOptions) *SysDepartmentModel {
    m := &SysDepartmentModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysDepartment_ID(id int64) *SysDepartmentModelOptions {
    return &SysDepartmentModelOptions{func(model *SysDepartmentModel) {
        model.ID = id
    }}
}

func SysDepartment_DepartmentName(department_name string) *SysDepartmentModelOptions {
    return &SysDepartmentModelOptions{func(model *SysDepartmentModel) {
        model.DepartmentName = department_name
    }}
}

func SysDepartment_ParentID(parent_id int64) *SysDepartmentModelOptions {
    return &SysDepartmentModelOptions{func(model *SysDepartmentModel) {
        model.ParentID = parent_id
    }}
}

func SysDepartment_DepartmentCode(department_code string) *SysDepartmentModelOptions {
    return &SysDepartmentModelOptions{func(model *SysDepartmentModel) {
        model.DepartmentCode = department_code
    }}
}

func SysDepartment_DepartmentType(department_type int) *SysDepartmentModelOptions {
    return &SysDepartmentModelOptions{func(model *SysDepartmentModel) {
        model.DepartmentType = department_type
    }}
}


func(model *SysDepartmentModel) WithID(id int64) *SysDepartmentModel  {
    model.ID = id
    return model
}

func(model *SysDepartmentModel) WithDepartmentName(department_name string) *SysDepartmentModel  {
    model.DepartmentName = department_name
    return model
}

func(model *SysDepartmentModel) WithParentID(parent_id int64) *SysDepartmentModel  {
    model.ParentID = parent_id
    return model
}

func(model *SysDepartmentModel) WithDepartmentCode(department_code string) *SysDepartmentModel  {
    model.DepartmentCode = department_code
    return model
}

func(model *SysDepartmentModel) WithDepartmentType(department_type int) *SysDepartmentModel  {
    model.DepartmentType = department_type
    return model
}









