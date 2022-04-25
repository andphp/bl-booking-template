package model


type SysAPIMenuRelationModel struct { 
    MenuID        int64        `form:"menu_id" gorm:"column:menu_id" json:"menuID"`                  // 菜单ID
    APIID         int64        `form:"api_id" gorm:"column:api_id" json:"apiID"`                     // 请求接口ID
    AbilityValue  string       `form:"ability_value" gorm:"column:ability_value" json:"abilityValue"` // 查询，添加，更新
    Path          string       `form:"path" gorm:"column:path" json:"path"`                          // 接口路由
    SystemCode    int          `form:"system_code" gorm:"column:system_code" json:"systemCode"`      // 
}

func (model *SysAPIMenuRelationModel) TableName() string {
	return "sys_api_menu_relations"
}

var SysAPIMenuRelationColumn = struct {
    MenuID               string
	APIID                string
	AbilityValue         string
	Path                 string
	SystemCode           string
	
}{
    MenuID              : "menu_id",
    APIID               : "api_id",
    AbilityValue        : "ability_value",
    Path                : "path",
    SystemCode          : "system_code",
    
}

type SysAPIMenuRelationModelOptions struct {
    apply func(*SysAPIMenuRelationModel)
}

func NewSysAPIMenuRelationModel(opts ...*SysAPIMenuRelationModelOptions) *SysAPIMenuRelationModel {
    m := &SysAPIMenuRelationModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysAPIMenuRelation_MenuID(menu_id int64) *SysAPIMenuRelationModelOptions {
    return &SysAPIMenuRelationModelOptions{func(model *SysAPIMenuRelationModel) {
        model.MenuID = menu_id
    }}
}

func SysAPIMenuRelation_APIID(api_id int64) *SysAPIMenuRelationModelOptions {
    return &SysAPIMenuRelationModelOptions{func(model *SysAPIMenuRelationModel) {
        model.APIID = api_id
    }}
}

func SysAPIMenuRelation_AbilityValue(ability_value string) *SysAPIMenuRelationModelOptions {
    return &SysAPIMenuRelationModelOptions{func(model *SysAPIMenuRelationModel) {
        model.AbilityValue = ability_value
    }}
}

func SysAPIMenuRelation_Path(path string) *SysAPIMenuRelationModelOptions {
    return &SysAPIMenuRelationModelOptions{func(model *SysAPIMenuRelationModel) {
        model.Path = path
    }}
}

func SysAPIMenuRelation_SystemCode(system_code int) *SysAPIMenuRelationModelOptions {
    return &SysAPIMenuRelationModelOptions{func(model *SysAPIMenuRelationModel) {
        model.SystemCode = system_code
    }}
}


func(model *SysAPIMenuRelationModel) WithMenuID(menu_id int64) *SysAPIMenuRelationModel  {
    model.MenuID = menu_id
    return model
}

func(model *SysAPIMenuRelationModel) WithAPIID(api_id int64) *SysAPIMenuRelationModel  {
    model.APIID = api_id
    return model
}

func(model *SysAPIMenuRelationModel) WithAbilityValue(ability_value string) *SysAPIMenuRelationModel  {
    model.AbilityValue = ability_value
    return model
}

func(model *SysAPIMenuRelationModel) WithPath(path string) *SysAPIMenuRelationModel  {
    model.Path = path
    return model
}

func(model *SysAPIMenuRelationModel) WithSystemCode(system_code int) *SysAPIMenuRelationModel  {
    model.SystemCode = system_code
    return model
}









