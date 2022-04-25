package model


type SysRolePositionRelationModel struct { 
    PositionID    int64   `form:"position_id" gorm:"column:position_id" json:"positionID"`      // 
    RoleID        int64   `form:"role_id" gorm:"column:role_id" json:"roleID"`                  // 
    SystemCode    int     `form:"system_code" gorm:"column:system_code" json:"systemCode"`      // 
    DeoartnebtID  int64   `form:"deoartnebt_id" gorm:"column:deoartnebt_id" json:"deoartnebtID"` // 
}

func (model *SysRolePositionRelationModel) TableName() string {
	return "sys_role_position_relations"
}

var SysRolePositionRelationColumn = struct {
    PositionID           string
	RoleID               string
	SystemCode           string
	DeoartnebtID         string
	
}{
    PositionID          : "position_id",
    RoleID              : "role_id",
    SystemCode          : "system_code",
    DeoartnebtID        : "deoartnebt_id",
    
}

type SysRolePositionRelationModelOptions struct {
    apply func(*SysRolePositionRelationModel)
}

func NewSysRolePositionRelationModel(opts ...*SysRolePositionRelationModelOptions) *SysRolePositionRelationModel {
    m := &SysRolePositionRelationModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysRolePositionRelation_PositionID(position_id int64) *SysRolePositionRelationModelOptions {
    return &SysRolePositionRelationModelOptions{func(model *SysRolePositionRelationModel) {
        model.PositionID = position_id
    }}
}

func SysRolePositionRelation_RoleID(role_id int64) *SysRolePositionRelationModelOptions {
    return &SysRolePositionRelationModelOptions{func(model *SysRolePositionRelationModel) {
        model.RoleID = role_id
    }}
}

func SysRolePositionRelation_SystemCode(system_code int) *SysRolePositionRelationModelOptions {
    return &SysRolePositionRelationModelOptions{func(model *SysRolePositionRelationModel) {
        model.SystemCode = system_code
    }}
}

func SysRolePositionRelation_DeoartnebtID(deoartnebt_id int64) *SysRolePositionRelationModelOptions {
    return &SysRolePositionRelationModelOptions{func(model *SysRolePositionRelationModel) {
        model.DeoartnebtID = deoartnebt_id
    }}
}


func(model *SysRolePositionRelationModel) WithPositionID(position_id int64) *SysRolePositionRelationModel  {
    model.PositionID = position_id
    return model
}

func(model *SysRolePositionRelationModel) WithRoleID(role_id int64) *SysRolePositionRelationModel  {
    model.RoleID = role_id
    return model
}

func(model *SysRolePositionRelationModel) WithSystemCode(system_code int) *SysRolePositionRelationModel  {
    model.SystemCode = system_code
    return model
}

func(model *SysRolePositionRelationModel) WithDeoartnebtID(deoartnebt_id int64) *SysRolePositionRelationModel  {
    model.DeoartnebtID = deoartnebt_id
    return model
}









