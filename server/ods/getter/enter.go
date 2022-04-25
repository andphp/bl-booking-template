package getter

import (
	"server/ods/getter/system"
	"sync"
)

var Getter *GetterStruct

func init() {
	Getter = NewGetterImplement()
}

// 查询获取====================================
type GetterStruct struct {
	SysUser system.SysUserGetter
	SysMenu system.SysMenuGetter
}

// 查询获取====================================
var getterImplementPool = sync.Pool{
	New: func() interface{} {
		return new(GetterStruct)
	},
}

func NewGetterImplement() *GetterStruct {
	newGetterImplement := getterImplementPool.Get().(*GetterStruct)
	// 回收对象 以备其他协程使用
	defer getterImplementPool.Put(newGetterImplement)
	return newGetterImplement
}
