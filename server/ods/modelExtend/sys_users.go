package modelExtend

import "github.com/andphp/go-gin/goby/utils"

type CurrentUser struct {
	ID              int64          `form:"id" gorm:"column:id" json:"id"`                                            //
	Email           string         `form:"email" gorm:"column:email" json:"email"`                                   // 邮箱
	Phone           string         `form:"phone" gorm:"column:phone" json:"phone"`                                   // 手机号
	EmailVerifiedAt utils.JSONTime `form:"email_verified_at" gorm:"column:email_verified_at" json:"emailVerifiedAt"` // 邮箱验证时间
	AccountName     string         `form:"account_name" gorm:"column:account_name" json:"accountName"`               // 用户登录名
	NickName        string         `form:"nick_name" gorm:"column:nick_name" json:"nickName"`                        // 昵称
	FullName        string         `form:"full_name" gorm:"column:full_name" json:"fullName"`                        // 实名
	Avatar          string         `form:"avatar" gorm:"column:avatar" json:"avatar"`                                // 头像
	LastLoginAt     utils.JSONTime `form:"last_login_at" gorm:"column:last_login_at" json:"lastLoginAt"`             // 最后登录日期
	LastToken       string         `form:"last_token" gorm:"column:last_token" json:"lastToken"`                     // 最新登录token
	LastIP          string         `form:"last_ip" gorm:"column:last_ip" json:"lastIP"`                              // 最后登录IP
	RoleID          int            `form:"role_id" gorm:"column:role_id" json:"roleID"`                              // 角色ID
	DepartmentID    int            `form:"department_id" gorm:"column:department_id" json:"departmentID"`            // 部门ID
	SideMode        string         `form:"side_mode" gorm:"column:side_mode" json:"sideMode"`                        // 用户侧边主题
	BaseColor       string         `form:"base_color" gorm:"column:base_color" json:"baseColor"`                     // 基础颜色
	ActiveColor     string         `form:"active_color" gorm:"column:active_color" json:"activeColor"`               // 活跃颜色
}
