package model
import "time"

type SysUserModel struct { 
    ID                int64         `form:"id" gorm:"column:id" json:"id"`                                            // 
    Email             string        `form:"email" gorm:"column:email" json:"email"`                                   // 邮箱
    Phone             string        `form:"phone" gorm:"column:phone" json:"phone"`                                   // 手机号
    EmailVerifiedAt   *time.Time    `form:"email_verified_at" gorm:"column:email_verified_at" json:"emailVerifiedAt"` // 邮箱验证时间
    AccountName       string        `form:"account_name" gorm:"column:account_name" json:"accountName"`               // 用户登录名
    NickName          string        `form:"nick_name" gorm:"column:nick_name" json:"nickName"`                        // 昵称
    FullName          string        `form:"full_name" gorm:"column:full_name" json:"fullName"`                        // 实名
    Password          string        `form:"password" gorm:"column:password" json:"password"`                          // 密码
    Avatar            string        `form:"avatar" gorm:"column:avatar" json:"avatar"`                                // 头像
    LastLoginAt       *time.Time    `form:"last_login_at" gorm:"column:last_login_at" json:"lastLoginAt"`             // 最后登录日期
    LastToken         string        `form:"last_token" gorm:"column:last_token" json:"lastToken"`                     // 最新登录token
    LastIP            string        `form:"last_ip" gorm:"column:last_ip" json:"lastIP"`                              // 最后登录IP
    RoleID            int           `form:"role_id" gorm:"column:role_id" json:"roleID"`                              // 角色ID
    Status            int           `form:"status" gorm:"column:status" json:"status"`                                // 状态: 1=启用 0=禁用
    DepartmentID      int           `form:"department_id" gorm:"column:department_id" json:"departmentID"`            // 部门ID
    CreatedAt         *time.Time    `form:"created_at" gorm:"column:created_at" json:"createdAt"`                     // 
    UpdatedAt         *time.Time    `form:"updated_at" gorm:"column:updated_at" json:"updatedAt"`                     // 
    DeletedAt         *time.Time    `form:"deleted_at" gorm:"column:deleted_at" json:"deletedAt"`                     // 删除时间 null未删除
    SideMode          string        `form:"side_mode" gorm:"column:side_mode" json:"sideMode"`                        // 用户侧边主题
    BaseColor         string        `form:"base_color" gorm:"column:base_color" json:"baseColor"`                     // 基础颜色
    ActiveColor       string        `form:"active_color" gorm:"column:active_color" json:"activeColor"`               // 活跃颜色
}

func (model *SysUserModel) TableName() string {
	return "sys_users"
}

var SysUserColumn = struct {
    ID                   string
	Email                string
	Phone                string
	EmailVerifiedAt      string
	AccountName          string
	NickName             string
	FullName             string
	Password             string
	Avatar               string
	LastLoginAt          string
	LastToken            string
	LastIP               string
	RoleID               string
	Status               string
	DepartmentID         string
	CreatedAt            string
	UpdatedAt            string
	DeletedAt            string
	SideMode             string
	BaseColor            string
	ActiveColor          string
	
}{
    ID                  : "id",
    Email               : "email",
    Phone               : "phone",
    EmailVerifiedAt     : "email_verified_at",
    AccountName         : "account_name",
    NickName            : "nick_name",
    FullName            : "full_name",
    Password            : "password",
    Avatar              : "avatar",
    LastLoginAt         : "last_login_at",
    LastToken           : "last_token",
    LastIP              : "last_ip",
    RoleID              : "role_id",
    Status              : "status",
    DepartmentID        : "department_id",
    CreatedAt           : "created_at",
    UpdatedAt           : "updated_at",
    DeletedAt           : "deleted_at",
    SideMode            : "side_mode",
    BaseColor           : "base_color",
    ActiveColor         : "active_color",
    
}

type SysUserModelOptions struct {
    apply func(*SysUserModel)
}

func NewSysUserModel(opts ...*SysUserModelOptions) *SysUserModel {
    m := &SysUserModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysUser_ID(id int64) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.ID = id
    }}
}

func SysUser_Email(email string) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.Email = email
    }}
}

func SysUser_Phone(phone string) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.Phone = phone
    }}
}

func SysUser_EmailVerifiedAt(email_verified_at time.Time) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.EmailVerifiedAt = &email_verified_at
    }}
}

func SysUser_AccountName(account_name string) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.AccountName = account_name
    }}
}

func SysUser_NickName(nick_name string) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.NickName = nick_name
    }}
}

func SysUser_FullName(full_name string) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.FullName = full_name
    }}
}

func SysUser_Password(password string) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.Password = password
    }}
}

func SysUser_Avatar(avatar string) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.Avatar = avatar
    }}
}

func SysUser_LastLoginAt(last_login_at time.Time) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.LastLoginAt = &last_login_at
    }}
}

func SysUser_LastToken(last_token string) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.LastToken = last_token
    }}
}

func SysUser_LastIP(last_ip string) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.LastIP = last_ip
    }}
}

func SysUser_RoleID(role_id int) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.RoleID = role_id
    }}
}

func SysUser_Status(status int) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.Status = status
    }}
}

func SysUser_DepartmentID(department_id int) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.DepartmentID = department_id
    }}
}

func SysUser_CreatedAt(created_at time.Time) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.CreatedAt = &created_at
    }}
}

func SysUser_UpdatedAt(updated_at time.Time) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.UpdatedAt = &updated_at
    }}
}

func SysUser_DeletedAt(deleted_at time.Time) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.DeletedAt = &deleted_at
    }}
}

func SysUser_SideMode(side_mode string) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.SideMode = side_mode
    }}
}

func SysUser_BaseColor(base_color string) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.BaseColor = base_color
    }}
}

func SysUser_ActiveColor(active_color string) *SysUserModelOptions {
    return &SysUserModelOptions{func(model *SysUserModel) {
        model.ActiveColor = active_color
    }}
}


func(model *SysUserModel) WithID(id int64) *SysUserModel  {
    model.ID = id
    return model
}

func(model *SysUserModel) WithEmail(email string) *SysUserModel  {
    model.Email = email
    return model
}

func(model *SysUserModel) WithPhone(phone string) *SysUserModel  {
    model.Phone = phone
    return model
}

func(model *SysUserModel) WithEmailVerifiedAt(email_verified_at time.Time) *SysUserModel  {
    model.EmailVerifiedAt = &email_verified_at
    return model
}

func(model *SysUserModel) WithAccountName(account_name string) *SysUserModel  {
    model.AccountName = account_name
    return model
}

func(model *SysUserModel) WithNickName(nick_name string) *SysUserModel  {
    model.NickName = nick_name
    return model
}

func(model *SysUserModel) WithFullName(full_name string) *SysUserModel  {
    model.FullName = full_name
    return model
}

func(model *SysUserModel) WithPassword(password string) *SysUserModel  {
    model.Password = password
    return model
}

func(model *SysUserModel) WithAvatar(avatar string) *SysUserModel  {
    model.Avatar = avatar
    return model
}

func(model *SysUserModel) WithLastLoginAt(last_login_at time.Time) *SysUserModel  {
    model.LastLoginAt = &last_login_at
    return model
}

func(model *SysUserModel) WithLastToken(last_token string) *SysUserModel  {
    model.LastToken = last_token
    return model
}

func(model *SysUserModel) WithLastIP(last_ip string) *SysUserModel  {
    model.LastIP = last_ip
    return model
}

func(model *SysUserModel) WithRoleID(role_id int) *SysUserModel  {
    model.RoleID = role_id
    return model
}

func(model *SysUserModel) WithStatus(status int) *SysUserModel  {
    model.Status = status
    return model
}

func(model *SysUserModel) WithDepartmentID(department_id int) *SysUserModel  {
    model.DepartmentID = department_id
    return model
}

func(model *SysUserModel) WithCreatedAt(created_at time.Time) *SysUserModel  {
    model.CreatedAt = &created_at
    return model
}

func(model *SysUserModel) WithUpdatedAt(updated_at time.Time) *SysUserModel  {
    model.UpdatedAt = &updated_at
    return model
}

func(model *SysUserModel) WithDeletedAt(deleted_at time.Time) *SysUserModel  {
    model.DeletedAt = &deleted_at
    return model
}

func(model *SysUserModel) WithSideMode(side_mode string) *SysUserModel  {
    model.SideMode = side_mode
    return model
}

func(model *SysUserModel) WithBaseColor(base_color string) *SysUserModel  {
    model.BaseColor = base_color
    return model
}

func(model *SysUserModel) WithActiveColor(active_color string) *SysUserModel  {
    model.ActiveColor = active_color
    return model
}









