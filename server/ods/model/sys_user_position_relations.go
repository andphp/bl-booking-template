package model


type SysUserPositionRelationModel struct { 
    UserID      int64  `gorm:"column:user_id" json:"userID"`         // 用户ID
    PositionID  int64  `gorm:"column:position_id" json:"positionID"` // 角色ID
}

func (model *SysUserPositionRelationModel) TableName() string {
	return "sys_user_position_relations"
}

var ColumnSysUserPositionRelation = struct {
    UserID               string
	PositionID           string
	
}{
    UserID              : "user_id",
    PositionID          : "position_id",
    
}

type SysUserPositionRelationModelOptions struct {
    apply func(*SysUserPositionRelationModel)
}

func NewSysUserPositionRelationModel(opts ...*SysUserPositionRelationModelOptions) *SysUserPositionRelationModel {
    m := &SysUserPositionRelationModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysUserPositionRelation_UserID(user_id int64) *SysUserPositionRelationModelOptions {
    return &SysUserPositionRelationModelOptions{func(model *SysUserPositionRelationModel) {
        model.UserID = user_id
    }}
}

func SysUserPositionRelation_PositionID(position_id int64) *SysUserPositionRelationModelOptions {
    return &SysUserPositionRelationModelOptions{func(model *SysUserPositionRelationModel) {
        model.PositionID = position_id
    }}
}


func(model *SysUserPositionRelationModel) WithUserID(user_id int64) *SysUserPositionRelationModel  {
    model.UserID = user_id
    return model
}

func(model *SysUserPositionRelationModel) WithPositionID(position_id int64) *SysUserPositionRelationModel  {
    model.PositionID = position_id
    return model
}









