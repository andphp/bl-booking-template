package modelExtend

import "github.com/andphp/go-gin/goby/utils"

type CurrentMenuTree struct {
	ID          int64              `form:"id" json:"id"`                     //
	Name        string             `form:"name" json:"name"`                 // 路由name
	Level       int                `form:"level"  json:"level"`              // 菜单等级，0表示顶级菜单
	ParentID    int64              `form:"parent_id"  json:"parentID"`       // 父菜单ID
	AllParentID string             `form:"all_parent_id" json:"allParentID"` // 所有父级id
	Path        string             `form:"path"  json:"path"`                // 路由path
	Hidden      int                `form:"hidden"   json:"hidden"`           // 是否显示在菜单中显示路由（默认值：false）
	TypeOf      string             `form:"type_of"   json:"typeOf"`          // 类型 page/outlink/nav
	Redirect    string             `form:"redirect"   json:"redirect"`       // 跳转路径（类型 page）
	Component   string             `form:"component"  json:"component"`      // 对应前端文件路径
	Title       string             `form:"title"  json:"title"`              // 附加属性:标题
	Description string             `form:"description" json:"description"`   // 附加属性:菜单描述
	Icon        string             `form:"icon"  json:"icon"`                // 附加属性:icon
	KeepAlive   int                `form:"keep_alive" json:"keepAlive"`      // 附加属性:当前路由是否缓存（默认值：true）
	Disable     int                `form:"disable" json:"disable"`           // 是否禁用，0：否，1：是
	Sort        int                `form:"sort"  json:"sort"`                // 排序标记
	Locale      string             `form:"locale"  json:"locale"`            // 语言切换英文标识
	Ability     utils.JSON         `form:"ability" json:"ability"`
	SystemCode  int                `form:"system_code" gorm:"column:system_code" json:"systemCode"` //
	Children    []CurrentMenuTree `form:"children" json:"children" gorm:"-"`
}

type MenuRequestParams struct {
	ID          int64              `form:"id" `                     //
	Name        string             `form:"name" `                 // 路由name
}