package go_designpattern

import "sync"

// SingletonObject 单例对象
type SingletonObject struct {
}

var (
	SingletonIns  *SingletonObject
	singletonOnce = &sync.Once{}
)

// GetSingletonIns 获取单例对象实例
func GetSingletonIns() *SingletonObject {
	if SingletonIns == nil {
		singletonOnce.Do(func() {
			SingletonIns = &SingletonObject{}
		})
	}
	return SingletonIns
}
