package model


type SysUserGroupRelationModel struct { 
    UserID        int64  `gorm:"column:user_id" json:"userID"`             // 用户ID
    UserGroupID   int64  `gorm:"column:user_group_id" json:"userGroupID"`  // 角色ID
}

func (model *SysUserGroupRelationModel) TableName() string {
	return "sys_user_group_relations"
}

var ColumnSysUserGroupRelation = struct {
    UserID               string
	UserGroupID          string
	
}{
    UserID              : "user_id",
    UserGroupID         : "user_group_id",
    
}

type SysUserGroupRelationModelOptions struct {
    apply func(*SysUserGroupRelationModel)
}

func NewSysUserGroupRelationModel(opts ...*SysUserGroupRelationModelOptions) *SysUserGroupRelationModel {
    m := &SysUserGroupRelationModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysUserGroupRelation_UserID(user_id int64) *SysUserGroupRelationModelOptions {
    return &SysUserGroupRelationModelOptions{func(model *SysUserGroupRelationModel) {
        model.UserID = user_id
    }}
}

func SysUserGroupRelation_UserGroupID(user_group_id int64) *SysUserGroupRelationModelOptions {
    return &SysUserGroupRelationModelOptions{func(model *SysUserGroupRelationModel) {
        model.UserGroupID = user_group_id
    }}
}


func(model *SysUserGroupRelationModel) WithUserID(user_id int64) *SysUserGroupRelationModel  {
    model.UserID = user_id
    return model
}

func(model *SysUserGroupRelationModel) WithUserGroupID(user_group_id int64) *SysUserGroupRelationModel  {
    model.UserGroupID = user_group_id
    return model
}









