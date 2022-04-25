package model


type SysRoleUserGroupRelationModel struct { 
    UserGroupID   int64   `gorm:"column:user_group_id" json:"userGroupID"`  // 用户ID
    RoleID        int64   `gorm:"column:role_id" json:"roleID"`             // 角色ID
    SystemCode    int     `gorm:"column:system_code" json:"systemCode"`     // 
}

func (model *SysRoleUserGroupRelationModel) TableName() string {
	return "sys_role_user_group_relations"
}

var ColumnSysRoleUserGroupRelation = struct {
    UserGroupID          string
	RoleID               string
	SystemCode           string
	
}{
    UserGroupID         : "user_group_id",
    RoleID              : "role_id",
    SystemCode          : "system_code",
    
}

type SysRoleUserGroupRelationModelOptions struct {
    apply func(*SysRoleUserGroupRelationModel)
}

func NewSysRoleUserGroupRelationModel(opts ...*SysRoleUserGroupRelationModelOptions) *SysRoleUserGroupRelationModel {
    m := &SysRoleUserGroupRelationModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysRoleUserGroupRelation_UserGroupID(user_group_id int64) *SysRoleUserGroupRelationModelOptions {
    return &SysRoleUserGroupRelationModelOptions{func(model *SysRoleUserGroupRelationModel) {
        model.UserGroupID = user_group_id
    }}
}

func SysRoleUserGroupRelation_RoleID(role_id int64) *SysRoleUserGroupRelationModelOptions {
    return &SysRoleUserGroupRelationModelOptions{func(model *SysRoleUserGroupRelationModel) {
        model.RoleID = role_id
    }}
}

func SysRoleUserGroupRelation_SystemCode(system_code int) *SysRoleUserGroupRelationModelOptions {
    return &SysRoleUserGroupRelationModelOptions{func(model *SysRoleUserGroupRelationModel) {
        model.SystemCode = system_code
    }}
}


func(model *SysRoleUserGroupRelationModel) WithUserGroupID(user_group_id int64) *SysRoleUserGroupRelationModel  {
    model.UserGroupID = user_group_id
    return model
}

func(model *SysRoleUserGroupRelationModel) WithRoleID(role_id int64) *SysRoleUserGroupRelationModel  {
    model.RoleID = role_id
    return model
}

func(model *SysRoleUserGroupRelationModel) WithSystemCode(system_code int) *SysRoleUserGroupRelationModel  {
    model.SystemCode = system_code
    return model
}









