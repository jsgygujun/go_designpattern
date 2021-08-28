package go_designpattern

import "sync"

type singleton struct {
}

var (
	SingletonIns *singleton
	once         = &sync.Once{}
)

func GetSingletonIns() *singleton {
	if SingletonIns == nil {
		once.Do(func() {
			SingletonIns = &singleton{}
		})
	}
	return SingletonIns
}
