package model
import "time"

type SysMenuModel struct { 
    ID            int64        `form:"id" gorm:"column:id" json:"id"`                                // 
    Name          string       `form:"name" gorm:"column:name" json:"name"`                          // 路由name
    Level         int          `form:"level" gorm:"column:level" json:"level"`                       // 菜单等级，0表示顶级菜单
    ParentID      int64        `form:"parent_id" gorm:"column:parent_id" json:"parentID"`            // 父菜单ID
    AllParentID   string       `form:"all_parent_id" gorm:"column:all_parent_id" json:"allParentID"` // 所有父级id
    Path          string       `form:"path" gorm:"column:path" json:"path"`                          // 路由path
    Hidden        int          `form:"hidden" gorm:"column:hidden" json:"hidden"`                    // 是否显示在菜单中显示路由（默认值：false）
    TypeOf        string       `form:"type_of" gorm:"column:type_of" json:"typeOf"`                  // 类型 page/outlink/nav
    Redirect      string       `form:"redirect" gorm:"column:redirect" json:"redirect"`              // 跳转路径（类型 page）
    Component     string       `form:"component" gorm:"column:component" json:"component"`           // 对应前端文件路径
    Title         string       `form:"title" gorm:"column:title" json:"title"`                       // 附加属性:标题
    Description   string       `form:"description" gorm:"column:description" json:"description"`     // 附加属性:菜单描述
    Icon          string       `form:"icon" gorm:"column:icon" json:"icon"`                          // 附加属性:icon
    KeepAlive     int          `form:"keep_alive" gorm:"column:keep_alive" json:"keepAlive"`         // 附加属性:当前路由是否缓存（默认值：true）
    Disable       int          `form:"disable" gorm:"column:disable" json:"disable"`                 // 是否禁用，0：否，1：是
    Sort          int          `form:"sort" gorm:"column:sort" json:"sort"`                          // 排序标记
    Locale        string       `form:"locale" gorm:"column:locale" json:"locale"`                    // 语言切换英文标识
    CreatedAt     *time.Time   `form:"created_at" gorm:"column:created_at" json:"createdAt"`         // 
    UpdatedAt     *time.Time   `form:"updated_at" gorm:"column:updated_at" json:"updatedAt"`         // 
    DeletedAt     *time.Time   `form:"deleted_at" gorm:"column:deleted_at" json:"deletedAt"`         // 删除时间 null未删除
    SystemCode    int          `form:"system_code" gorm:"column:system_code" json:"systemCode"`      // 
}

func (model *SysMenuModel) TableName() string {
	return "sys_menus"
}

var SysMenuColumn = struct {
    ID                   string
	Name                 string
	Level                string
	ParentID             string
	AllParentID          string
	Path                 string
	Hidden               string
	TypeOf               string
	Redirect             string
	Component            string
	Title                string
	Description          string
	Icon                 string
	KeepAlive            string
	Disable              string
	Sort                 string
	Locale               string
	CreatedAt            string
	UpdatedAt            string
	DeletedAt            string
	SystemCode           string
	
}{
    ID                  : "id",
    Name                : "name",
    Level               : "level",
    ParentID            : "parent_id",
    AllParentID         : "all_parent_id",
    Path                : "path",
    Hidden              : "hidden",
    TypeOf              : "type_of",
    Redirect            : "redirect",
    Component           : "component",
    Title               : "title",
    Description         : "description",
    Icon                : "icon",
    KeepAlive           : "keep_alive",
    Disable             : "disable",
    Sort                : "sort",
    Locale              : "locale",
    CreatedAt           : "created_at",
    UpdatedAt           : "updated_at",
    DeletedAt           : "deleted_at",
    SystemCode          : "system_code",
    
}

type SysMenuModelOptions struct {
    apply func(*SysMenuModel)
}

func NewSysMenuModel(opts ...*SysMenuModelOptions) *SysMenuModel {
    m := &SysMenuModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysMenu_ID(id int64) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.ID = id
    }}
}

func SysMenu_Name(name string) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.Name = name
    }}
}

func SysMenu_Level(level int) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.Level = level
    }}
}

func SysMenu_ParentID(parent_id int64) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.ParentID = parent_id
    }}
}

func SysMenu_AllParentID(all_parent_id string) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.AllParentID = all_parent_id
    }}
}

func SysMenu_Path(path string) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.Path = path
    }}
}

func SysMenu_Hidden(hidden int) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.Hidden = hidden
    }}
}

func SysMenu_TypeOf(type_of string) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.TypeOf = type_of
    }}
}

func SysMenu_Redirect(redirect string) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.Redirect = redirect
    }}
}

func SysMenu_Component(component string) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.Component = component
    }}
}

func SysMenu_Title(title string) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.Title = title
    }}
}

func SysMenu_Description(description string) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.Description = description
    }}
}

func SysMenu_Icon(icon string) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.Icon = icon
    }}
}

func SysMenu_KeepAlive(keep_alive int) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.KeepAlive = keep_alive
    }}
}

func SysMenu_Disable(disable int) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.Disable = disable
    }}
}

func SysMenu_Sort(sort int) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.Sort = sort
    }}
}

func SysMenu_Locale(locale string) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.Locale = locale
    }}
}

func SysMenu_CreatedAt(created_at time.Time) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.CreatedAt = &created_at
    }}
}

func SysMenu_UpdatedAt(updated_at time.Time) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.UpdatedAt = &updated_at
    }}
}

func SysMenu_DeletedAt(deleted_at time.Time) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.DeletedAt = &deleted_at
    }}
}

func SysMenu_SystemCode(system_code int) *SysMenuModelOptions {
    return &SysMenuModelOptions{func(model *SysMenuModel) {
        model.SystemCode = system_code
    }}
}


func(model *SysMenuModel) WithID(id int64) *SysMenuModel  {
    model.ID = id
    return model
}

func(model *SysMenuModel) WithName(name string) *SysMenuModel  {
    model.Name = name
    return model
}

func(model *SysMenuModel) WithLevel(level int) *SysMenuModel  {
    model.Level = level
    return model
}

func(model *SysMenuModel) WithParentID(parent_id int64) *SysMenuModel  {
    model.ParentID = parent_id
    return model
}

func(model *SysMenuModel) WithAllParentID(all_parent_id string) *SysMenuModel  {
    model.AllParentID = all_parent_id
    return model
}

func(model *SysMenuModel) WithPath(path string) *SysMenuModel  {
    model.Path = path
    return model
}

func(model *SysMenuModel) WithHidden(hidden int) *SysMenuModel  {
    model.Hidden = hidden
    return model
}

func(model *SysMenuModel) WithTypeOf(type_of string) *SysMenuModel  {
    model.TypeOf = type_of
    return model
}

func(model *SysMenuModel) WithRedirect(redirect string) *SysMenuModel  {
    model.Redirect = redirect
    return model
}

func(model *SysMenuModel) WithComponent(component string) *SysMenuModel  {
    model.Component = component
    return model
}

func(model *SysMenuModel) WithTitle(title string) *SysMenuModel  {
    model.Title = title
    return model
}

func(model *SysMenuModel) WithDescription(description string) *SysMenuModel  {
    model.Description = description
    return model
}

func(model *SysMenuModel) WithIcon(icon string) *SysMenuModel  {
    model.Icon = icon
    return model
}

func(model *SysMenuModel) WithKeepAlive(keep_alive int) *SysMenuModel  {
    model.KeepAlive = keep_alive
    return model
}

func(model *SysMenuModel) WithDisable(disable int) *SysMenuModel  {
    model.Disable = disable
    return model
}

func(model *SysMenuModel) WithSort(sort int) *SysMenuModel  {
    model.Sort = sort
    return model
}

func(model *SysMenuModel) WithLocale(locale string) *SysMenuModel  {
    model.Locale = locale
    return model
}

func(model *SysMenuModel) WithCreatedAt(created_at time.Time) *SysMenuModel  {
    model.CreatedAt = &created_at
    return model
}

func(model *SysMenuModel) WithUpdatedAt(updated_at time.Time) *SysMenuModel  {
    model.UpdatedAt = &updated_at
    return model
}

func(model *SysMenuModel) WithDeletedAt(deleted_at time.Time) *SysMenuModel  {
    model.DeletedAt = &deleted_at
    return model
}

func(model *SysMenuModel) WithSystemCode(system_code int) *SysMenuModel  {
    model.SystemCode = system_code
    return model
}









