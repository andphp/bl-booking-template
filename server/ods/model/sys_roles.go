package model

import (
    "github.com/andphp/go-gin/goby/utils"
    "time"
)

type SysRoleModel struct { 
    ID                 int64        `gorm:"column:id" json:"id"`                                // 
    RoleName           string       `gorm:"column:role_name" json:"roleName"`                   // 角色名
    RoleNameEn         string       `gorm:"column:role_name_en" json:"roleNameEn"`              // 角色名_英文
    RoleParentID       string       `gorm:"column:role_parent_id" json:"roleParentID"`          // 父角色ID（父级包含子级角色权限）
    RoleDefaultPath    string       `gorm:"column:role_default_path" json:"roleDefaultPath"`    // 默认菜单
    RoleRemark         string       `gorm:"column:role_remark" json:"roleRemark"`               // 备注
    RoleAllParentID    string       `gorm:"column:role_all_parent_id" json:"roleAllParentID"`   // 所有上级id
    CreatedAt          utils.JSONTime `gorm:"column:created_at" json:"createdAt"`                 // 
    UpdatedAt          utils.JSONTime `gorm:"column:updated_at" json:"updatedAt"`                 // 
    DeletedAt          utils.JSONTime `gorm:"column:deleted_at" json:"deletedAt"`                 // 删除
    SystemCode         int          `gorm:"column:system_code" json:"systemCode"`               // 
}

func (model *SysRoleModel) TableName() string {
	return "sys_roles"
}

var ColumnSysRole = struct {
    ID                   string
	RoleName             string
	RoleNameEn           string
	RoleParentID         string
	RoleDefaultPath      string
	RoleRemark           string
	RoleAllParentID      string
	CreatedAt            string
	UpdatedAt            string
	DeletedAt            string
	SystemCode           string
	
}{
    ID                  : "id",
    RoleName            : "role_name",
    RoleNameEn          : "role_name_en",
    RoleParentID        : "role_parent_id",
    RoleDefaultPath     : "role_default_path",
    RoleRemark          : "role_remark",
    RoleAllParentID     : "role_all_parent_id",
    CreatedAt           : "created_at",
    UpdatedAt           : "updated_at",
    DeletedAt           : "deleted_at",
    SystemCode          : "system_code",
    
}

type SysRoleModelOptions struct {
    apply func(*SysRoleModel)
}

func NewSysRoleModel(opts ...*SysRoleModelOptions) *SysRoleModel {
    m := &SysRoleModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysRole_ID(id int64) *SysRoleModelOptions {
    return &SysRoleModelOptions{func(model *SysRoleModel) {
        model.ID = id
    }}
}

func SysRole_RoleName(role_name string) *SysRoleModelOptions {
    return &SysRoleModelOptions{func(model *SysRoleModel) {
        model.RoleName = role_name
    }}
}

func SysRole_RoleNameEn(role_name_en string) *SysRoleModelOptions {
    return &SysRoleModelOptions{func(model *SysRoleModel) {
        model.RoleNameEn = role_name_en
    }}
}

func SysRole_RoleParentID(role_parent_id string) *SysRoleModelOptions {
    return &SysRoleModelOptions{func(model *SysRoleModel) {
        model.RoleParentID = role_parent_id
    }}
}

func SysRole_RoleDefaultPath(role_default_path string) *SysRoleModelOptions {
    return &SysRoleModelOptions{func(model *SysRoleModel) {
        model.RoleDefaultPath = role_default_path
    }}
}

func SysRole_RoleRemark(role_remark string) *SysRoleModelOptions {
    return &SysRoleModelOptions{func(model *SysRoleModel) {
        model.RoleRemark = role_remark
    }}
}

func SysRole_RoleAllParentID(role_all_parent_id string) *SysRoleModelOptions {
    return &SysRoleModelOptions{func(model *SysRoleModel) {
        model.RoleAllParentID = role_all_parent_id
    }}
}

func SysRole_CreatedAt(created_at time.Time) *SysRoleModelOptions {
    return &SysRoleModelOptions{func(model *SysRoleModel) {
        model.CreatedAt = utils.JSONTime(created_at)
    }}
}

func SysRole_UpdatedAt(updated_at time.Time) *SysRoleModelOptions {
    return &SysRoleModelOptions{func(model *SysRoleModel) {
        model.UpdatedAt = utils.JSONTime(updated_at)
    }}
}

func SysRole_DeletedAt(deleted_at time.Time) *SysRoleModelOptions {
    return &SysRoleModelOptions{func(model *SysRoleModel) {
        model.DeletedAt = utils.JSONTime(deleted_at)
    }}
}

func SysRole_SystemCode(system_code int) *SysRoleModelOptions {
    return &SysRoleModelOptions{func(model *SysRoleModel) {
        model.SystemCode = system_code
    }}
}


func(model *SysRoleModel) WithID(id int64) *SysRoleModel  {
    model.ID = id
    return model
}

func(model *SysRoleModel) WithRoleName(role_name string) *SysRoleModel  {
    model.RoleName = role_name
    return model
}

func(model *SysRoleModel) WithRoleNameEn(role_name_en string) *SysRoleModel  {
    model.RoleNameEn = role_name_en
    return model
}

func(model *SysRoleModel) WithRoleParentID(role_parent_id string) *SysRoleModel  {
    model.RoleParentID = role_parent_id
    return model
}

func(model *SysRoleModel) WithRoleDefaultPath(role_default_path string) *SysRoleModel  {
    model.RoleDefaultPath = role_default_path
    return model
}

func(model *SysRoleModel) WithRoleRemark(role_remark string) *SysRoleModel  {
    model.RoleRemark = role_remark
    return model
}

func(model *SysRoleModel) WithRoleAllParentID(role_all_parent_id string) *SysRoleModel  {
    model.RoleAllParentID = role_all_parent_id
    return model
}

func(model *SysRoleModel) WithCreatedAt(created_at time.Time) *SysRoleModel  {
    model.CreatedAt = utils.JSONTime(created_at)
    return model
}

func(model *SysRoleModel) WithUpdatedAt(updated_at time.Time) *SysRoleModel  {
    model.UpdatedAt = utils.JSONTime(updated_at)
    return model
}

func(model *SysRoleModel) WithDeletedAt(deleted_at time.Time) *SysRoleModel  {
    model.DeletedAt = utils.JSONTime(deleted_at)
    return model
}

func(model *SysRoleModel) WithSystemCode(system_code int) *SysRoleModel  {
    model.SystemCode = system_code
    return model
}









