package controller

import (
	"server/controller/login"
	"server/controller/unauthority"
)

type ApiGroup struct {
	UnAuthorityGroup unauthority.WhiteListGroup
	Login            login.Login
}

var ApiGroupApp = new(ApiGroup)
