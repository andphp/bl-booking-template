package model
import "time"

type SysAPIModel struct { 
    ID          int64        `form:"id" gorm:"column:id" json:"id"`                          // 
    Name        string       `form:"name" gorm:"column:name" json:"name"`                    // 名称（查询、添加、更新、删除、批量添加、批量更新、批量删除）
    Path        string       `form:"path" gorm:"column:path" json:"path"`                    // api路径
    Description string       `form:"description" gorm:"column:description" json:"description"` // api中文描述
    APIGroup    string       `form:"api_group" gorm:"column:api_group" json:"apiGroup"`      // api组
    Method      string       `form:"method" gorm:"column:method" json:"method"`              // 请求方式
    CreatedAt   *time.Time   `form:"created_at" gorm:"column:created_at" json:"createdAt"`   // 
    UpdatedAt   *time.Time   `form:"updated_at" gorm:"column:updated_at" json:"updatedAt"`   // 
    DeletedAt   *time.Time   `form:"deleted_at" gorm:"column:deleted_at" json:"deletedAt"`   // 删除时间 null未删除
    SystemCode  int          `form:"system_code" gorm:"column:system_code" json:"systemCode"` // 
}

func (model *SysAPIModel) TableName() string {
	return "sys_apis"
}

var SysAPIColumn = struct {
    ID                   string
	Name                 string
	Path                 string
	Description          string
	APIGroup             string
	Method               string
	CreatedAt            string
	UpdatedAt            string
	DeletedAt            string
	SystemCode           string
	
}{
    ID                  : "id",
    Name                : "name",
    Path                : "path",
    Description         : "description",
    APIGroup            : "api_group",
    Method              : "method",
    CreatedAt           : "created_at",
    UpdatedAt           : "updated_at",
    DeletedAt           : "deleted_at",
    SystemCode          : "system_code",
    
}

type SysAPIModelOptions struct {
    apply func(*SysAPIModel)
}

func NewSysAPIModel(opts ...*SysAPIModelOptions) *SysAPIModel {
    m := &SysAPIModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysAPI_ID(id int64) *SysAPIModelOptions {
    return &SysAPIModelOptions{func(model *SysAPIModel) {
        model.ID = id
    }}
}

func SysAPI_Name(name string) *SysAPIModelOptions {
    return &SysAPIModelOptions{func(model *SysAPIModel) {
        model.Name = name
    }}
}

func SysAPI_Path(path string) *SysAPIModelOptions {
    return &SysAPIModelOptions{func(model *SysAPIModel) {
        model.Path = path
    }}
}

func SysAPI_Description(description string) *SysAPIModelOptions {
    return &SysAPIModelOptions{func(model *SysAPIModel) {
        model.Description = description
    }}
}

func SysAPI_APIGroup(api_group string) *SysAPIModelOptions {
    return &SysAPIModelOptions{func(model *SysAPIModel) {
        model.APIGroup = api_group
    }}
}

func SysAPI_Method(method string) *SysAPIModelOptions {
    return &SysAPIModelOptions{func(model *SysAPIModel) {
        model.Method = method
    }}
}

func SysAPI_CreatedAt(created_at time.Time) *SysAPIModelOptions {
    return &SysAPIModelOptions{func(model *SysAPIModel) {
        model.CreatedAt = &created_at
    }}
}

func SysAPI_UpdatedAt(updated_at time.Time) *SysAPIModelOptions {
    return &SysAPIModelOptions{func(model *SysAPIModel) {
        model.UpdatedAt = &updated_at
    }}
}

func SysAPI_DeletedAt(deleted_at time.Time) *SysAPIModelOptions {
    return &SysAPIModelOptions{func(model *SysAPIModel) {
        model.DeletedAt = &deleted_at
    }}
}

func SysAPI_SystemCode(system_code int) *SysAPIModelOptions {
    return &SysAPIModelOptions{func(model *SysAPIModel) {
        model.SystemCode = system_code
    }}
}


func(model *SysAPIModel) WithID(id int64) *SysAPIModel  {
    model.ID = id
    return model
}

func(model *SysAPIModel) WithName(name string) *SysAPIModel  {
    model.Name = name
    return model
}

func(model *SysAPIModel) WithPath(path string) *SysAPIModel  {
    model.Path = path
    return model
}

func(model *SysAPIModel) WithDescription(description string) *SysAPIModel  {
    model.Description = description
    return model
}

func(model *SysAPIModel) WithAPIGroup(api_group string) *SysAPIModel  {
    model.APIGroup = api_group
    return model
}

func(model *SysAPIModel) WithMethod(method string) *SysAPIModel  {
    model.Method = method
    return model
}

func(model *SysAPIModel) WithCreatedAt(created_at time.Time) *SysAPIModel  {
    model.CreatedAt = &created_at
    return model
}

func(model *SysAPIModel) WithUpdatedAt(updated_at time.Time) *SysAPIModel  {
    model.UpdatedAt = &updated_at
    return model
}

func(model *SysAPIModel) WithDeletedAt(deleted_at time.Time) *SysAPIModel  {
    model.DeletedAt = &deleted_at
    return model
}

func(model *SysAPIModel) WithSystemCode(system_code int) *SysAPIModel  {
    model.SystemCode = system_code
    return model
}









