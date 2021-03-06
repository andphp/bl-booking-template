package model

import (
    "github.com/andphp/go-gin/goby/utils"
    "time"
)

type SysUserGroupModel struct { 
    ID                int64        `gorm:"column:id" json:"id"`                              // 
    GroupName         string       `gorm:"column:group_name" json:"groupName"`               // 
    GroupCode         string       `gorm:"column:group_code" json:"groupCode"`               // 
    GroupDescription  string       `gorm:"column:group_description" json:"groupDescription"` // 
    CreatedAt         utils.JSONTime `gorm:"column:created_at" json:"createdAt"`               // 
    UpdatedAt         utils.JSONTime `gorm:"column:updated_at" json:"updatedAt"`               // 
    DeletedAt         utils.JSONTime `gorm:"column:deleted_at" json:"deletedAt"`               // 
    ExtendType        int          `gorm:"column:extend_type" json:"extendType"`             // 扩展类型
    ExtendID          int64        `gorm:"column:extend_id" json:"extendID"`                 // 扩展id
}

func (model *SysUserGroupModel) TableName() string {
	return "sys_user_groups"
}

var ColumnSysUserGroup = struct {
    ID                   string
	GroupName            string
	GroupCode            string
	GroupDescription     string
	CreatedAt            string
	UpdatedAt            string
	DeletedAt            string
	ExtendType           string
	ExtendID             string
	
}{
    ID                  : "id",
    GroupName           : "group_name",
    GroupCode           : "group_code",
    GroupDescription    : "group_description",
    CreatedAt           : "created_at",
    UpdatedAt           : "updated_at",
    DeletedAt           : "deleted_at",
    ExtendType          : "extend_type",
    ExtendID            : "extend_id",
    
}

type SysUserGroupModelOptions struct {
    apply func(*SysUserGroupModel)
}

func NewSysUserGroupModel(opts ...*SysUserGroupModelOptions) *SysUserGroupModel {
    m := &SysUserGroupModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysUserGroup_ID(id int64) *SysUserGroupModelOptions {
    return &SysUserGroupModelOptions{func(model *SysUserGroupModel) {
        model.ID = id
    }}
}

func SysUserGroup_GroupName(group_name string) *SysUserGroupModelOptions {
    return &SysUserGroupModelOptions{func(model *SysUserGroupModel) {
        model.GroupName = group_name
    }}
}

func SysUserGroup_GroupCode(group_code string) *SysUserGroupModelOptions {
    return &SysUserGroupModelOptions{func(model *SysUserGroupModel) {
        model.GroupCode = group_code
    }}
}

func SysUserGroup_GroupDescription(group_description string) *SysUserGroupModelOptions {
    return &SysUserGroupModelOptions{func(model *SysUserGroupModel) {
        model.GroupDescription = group_description
    }}
}

func SysUserGroup_CreatedAt(created_at time.Time) *SysUserGroupModelOptions {
    return &SysUserGroupModelOptions{func(model *SysUserGroupModel) {
        model.CreatedAt = utils.JSONTime(created_at)
    }}
}

func SysUserGroup_UpdatedAt(updated_at time.Time) *SysUserGroupModelOptions {
    return &SysUserGroupModelOptions{func(model *SysUserGroupModel) {
        model.UpdatedAt = utils.JSONTime(updated_at)
    }}
}

func SysUserGroup_DeletedAt(deleted_at time.Time) *SysUserGroupModelOptions {
    return &SysUserGroupModelOptions{func(model *SysUserGroupModel) {
        model.DeletedAt = utils.JSONTime(deleted_at)
    }}
}

func SysUserGroup_ExtendType(extend_type int) *SysUserGroupModelOptions {
    return &SysUserGroupModelOptions{func(model *SysUserGroupModel) {
        model.ExtendType = extend_type
    }}
}

func SysUserGroup_ExtendID(extend_id int64) *SysUserGroupModelOptions {
    return &SysUserGroupModelOptions{func(model *SysUserGroupModel) {
        model.ExtendID = extend_id
    }}
}


func(model *SysUserGroupModel) WithID(id int64) *SysUserGroupModel  {
    model.ID = id
    return model
}

func(model *SysUserGroupModel) WithGroupName(group_name string) *SysUserGroupModel  {
    model.GroupName = group_name
    return model
}

func(model *SysUserGroupModel) WithGroupCode(group_code string) *SysUserGroupModel  {
    model.GroupCode = group_code
    return model
}

func(model *SysUserGroupModel) WithGroupDescription(group_description string) *SysUserGroupModel  {
    model.GroupDescription = group_description
    return model
}

func(model *SysUserGroupModel) WithCreatedAt(created_at time.Time) *SysUserGroupModel  {
    model.CreatedAt = utils.JSONTime(created_at)
    return model
}

func(model *SysUserGroupModel) WithUpdatedAt(updated_at time.Time) *SysUserGroupModel  {
    model.UpdatedAt = utils.JSONTime(updated_at)
    return model
}

func(model *SysUserGroupModel) WithDeletedAt(deleted_at time.Time) *SysUserGroupModel  {
    model.DeletedAt = utils.JSONTime(deleted_at)
    return model
}

func(model *SysUserGroupModel) WithExtendType(extend_type int) *SysUserGroupModel  {
    model.ExtendType = extend_type
    return model
}

func(model *SysUserGroupModel) WithExtendID(extend_id int64) *SysUserGroupModel  {
    model.ExtendID = extend_id
    return model
}









