package tools

import (
	"server/ods/modelExtend"

	"github.com/andphp/go-gin/goby"
	"github.com/gin-gonic/gin"
)

func GetSession(c *gin.Context) (*modelExtend.CustomSession, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	session, err := j.ParseToken(token)
	if err != nil {
		goby.GOBY_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return session, err
}

// 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) int64 {
	if session, exists := c.Get("session"); !exists {
		if cl, err := GetSession(c); err != nil {
			return 0
		} else {
			return cl.ID
		}
	} else {
		waitUse := session.(*modelExtend.CustomSession)
		return waitUse.ID
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserRoleId(c *gin.Context) int {
	if session, exists := c.Get("session"); !exists {
		if cl, err := GetSession(c); err != nil {
			return 0
		} else {
			return cl.RoleID
		}
	} else {
		waitUse := session.(*modelExtend.CustomSession)
		return waitUse.RoleID
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *modelExtend.CustomSession {
	if session, exists := c.Get("session"); !exists {
		if cl, err := GetSession(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := session.(*modelExtend.CustomSession)
		return waitUse
	}
}
