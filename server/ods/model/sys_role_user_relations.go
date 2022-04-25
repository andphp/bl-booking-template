package model


type SysRoleUserRelationModel struct { 
    UserID      int64   `gorm:"column:user_id" json:"userID"`         // 用户ID
    RoleID      int64   `gorm:"column:role_id" json:"roleID"`         // 角色ID
    SystemCode  int     `gorm:"column:system_code" json:"systemCode"` // 
}

func (model *SysRoleUserRelationModel) TableName() string {
	return "sys_role_user_relations"
}

var ColumnSysRoleUserRelation = struct {
    UserID               string
	RoleID               string
	SystemCode           string
	
}{
    UserID              : "user_id",
    RoleID              : "role_id",
    SystemCode          : "system_code",
    
}

type SysRoleUserRelationModelOptions struct {
    apply func(*SysRoleUserRelationModel)
}

func NewSysRoleUserRelationModel(opts ...*SysRoleUserRelationModelOptions) *SysRoleUserRelationModel {
    m := &SysRoleUserRelationModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysRoleUserRelation_UserID(user_id int64) *SysRoleUserRelationModelOptions {
    return &SysRoleUserRelationModelOptions{func(model *SysRoleUserRelationModel) {
        model.UserID = user_id
    }}
}

func SysRoleUserRelation_RoleID(role_id int64) *SysRoleUserRelationModelOptions {
    return &SysRoleUserRelationModelOptions{func(model *SysRoleUserRelationModel) {
        model.RoleID = role_id
    }}
}

func SysRoleUserRelation_SystemCode(system_code int) *SysRoleUserRelationModelOptions {
    return &SysRoleUserRelationModelOptions{func(model *SysRoleUserRelationModel) {
        model.SystemCode = system_code
    }}
}


func(model *SysRoleUserRelationModel) WithUserID(user_id int64) *SysRoleUserRelationModel  {
    model.UserID = user_id
    return model
}

func(model *SysRoleUserRelationModel) WithRoleID(role_id int64) *SysRoleUserRelationModel  {
    model.RoleID = role_id
    return model
}

func(model *SysRoleUserRelationModel) WithSystemCode(system_code int) *SysRoleUserRelationModel  {
    model.SystemCode = system_code
    return model
}









