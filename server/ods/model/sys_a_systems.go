package model


type SysASystemModel struct { 
    ID          int64        `gorm:"column:id" json:"id"`                  // 
    SystemName  string       `gorm:"column:system_name" json:"systemName"` // 
    SystemCode  int          `gorm:"column:system_code" json:"systemCode"` // 
}

func (model *SysASystemModel) TableName() string {
	return "sys_a_systems"
}

var ColumnSysASystem = struct {
    ID                   string
	SystemName           string
	SystemCode           string
	
}{
    ID                  : "id",
    SystemName          : "system_name",
    SystemCode          : "system_code",
    
}

type SysASystemModelOptions struct {
    apply func(*SysASystemModel)
}

func NewSysASystemModel(opts ...*SysASystemModelOptions) *SysASystemModel {
    m := &SysASystemModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysASystem_ID(id int64) *SysASystemModelOptions {
    return &SysASystemModelOptions{func(model *SysASystemModel) {
        model.ID = id
    }}
}

func SysASystem_SystemName(system_name string) *SysASystemModelOptions {
    return &SysASystemModelOptions{func(model *SysASystemModel) {
        model.SystemName = system_name
    }}
}

func SysASystem_SystemCode(system_code int) *SysASystemModelOptions {
    return &SysASystemModelOptions{func(model *SysASystemModel) {
        model.SystemCode = system_code
    }}
}


func(model *SysASystemModel) WithID(id int64) *SysASystemModel  {
    model.ID = id
    return model
}

func(model *SysASystemModel) WithSystemName(system_name string) *SysASystemModel  {
    model.SystemName = system_name
    return model
}

func(model *SysASystemModel) WithSystemCode(system_code int) *SysASystemModel  {
    model.SystemCode = system_code
    return model
}









