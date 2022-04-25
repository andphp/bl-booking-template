package setter

import (
	"sync"
)

var Setter *SetterStruct

func init() {
	Setter = NewSetterImplement()
}

// 设置更新===================================
type SetterStruct struct{}

// 设置更新===================================
var setterImplementPool = sync.Pool{
	New: func() interface{} {
		return new(SetterStruct)
	},
}

func NewSetterImplement() *SetterStruct {
	newSetterImplement := setterImplementPool.Get().(*SetterStruct)
	// 回收对象 以备其他协程使用
	defer setterImplementPool.Put(newSetterImplement)
	return newSetterImplement
}
