package model

import "github.com/andphp/go-gin/goby/utils"

type SysMenuRoleRelationModel struct {
    RoleID      int64   `gorm:"column:role_id" json:"roleID"`         // 角色ID
    MenuID      int64   `gorm:"column:menu_id" json:"menuID"`         // 菜单ID
    Ability     utils.JSON `gorm:"column:ability" json:"ability"`        // 权限：序列化["READ","WRITE","UPDATE","DELETE","IMPORT","EXPRORT","BATCH_WRITE","BATCH_UPDATE","BATCH_DELETE"]
    SystemCode  int     `gorm:"column:system_code" json:"systemCode"` // 
}

func (model *SysMenuRoleRelationModel) TableName() string {
	return "sys_menu_role_relations"
}

var ColumnSysMenuRoleRelation = struct {
    RoleID               string
	MenuID               string
	Ability              string
	SystemCode           string
	
}{
    RoleID              : "role_id",
    MenuID              : "menu_id",
    Ability             : "ability",
    SystemCode          : "system_code",
    
}

type SysMenuRoleRelationModelOptions struct {
    apply func(*SysMenuRoleRelationModel)
}

func NewSysMenuRoleRelationModel(opts ...*SysMenuRoleRelationModelOptions) *SysMenuRoleRelationModel {
    m := &SysMenuRoleRelationModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysMenuRoleRelation_RoleID(role_id int64) *SysMenuRoleRelationModelOptions {
    return &SysMenuRoleRelationModelOptions{func(model *SysMenuRoleRelationModel) {
        model.RoleID = role_id
    }}
}

func SysMenuRoleRelation_MenuID(menu_id int64) *SysMenuRoleRelationModelOptions {
    return &SysMenuRoleRelationModelOptions{func(model *SysMenuRoleRelationModel) {
        model.MenuID = menu_id
    }}
}

func SysMenuRoleRelation_Ability(ability utils.JSON) *SysMenuRoleRelationModelOptions {
    return &SysMenuRoleRelationModelOptions{func(model *SysMenuRoleRelationModel) {
        model.Ability = utils.StructToJSON(ability)
    }}
}

func SysMenuRoleRelation_SystemCode(system_code int) *SysMenuRoleRelationModelOptions {
    return &SysMenuRoleRelationModelOptions{func(model *SysMenuRoleRelationModel) {
        model.SystemCode = system_code
    }}
}


func(model *SysMenuRoleRelationModel) WithRoleID(role_id int64) *SysMenuRoleRelationModel  {
    model.RoleID = role_id
    return model
}

func(model *SysMenuRoleRelationModel) WithMenuID(menu_id int64) *SysMenuRoleRelationModel  {
    model.MenuID = menu_id
    return model
}

func(model *SysMenuRoleRelationModel) WithAbility(ability utils.JSON) *SysMenuRoleRelationModel  {
    model.Ability = utils.StructToJSON(ability)
    return model
}

func(model *SysMenuRoleRelationModel) WithSystemCode(system_code int) *SysMenuRoleRelationModel  {
    model.SystemCode = system_code
    return model
}









