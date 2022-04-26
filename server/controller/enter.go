package controller

import (
	"server/controller/login"
	"server/controller/system"
	"server/controller/unauthority"
)

type ApiGroup struct {
	UnAuthorityGroup unauthority.WhiteListGroup
	Login            login.Login
	System           system.Group
}

var ApiGroupApp = new(ApiGroup)
