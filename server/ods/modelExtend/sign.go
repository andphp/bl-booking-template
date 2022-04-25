package modelExtend

import "github.com/andphp/go-gin/goby/utils"

type Login struct {
	AccountName string         `form:"account_name" binding:"required"` // 用户名
	Password    string         `form:"password" binding:"required"`     // 密码
	NickName    string         `form:"nick_name"`
	ID          int64          `form:"id"`
	RoleID      int            `form:"role_id"`
	UpdatedAt   utils.JSONTime `form:"updated_at"`
}
