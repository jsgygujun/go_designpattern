package go_designpattern

import (
	"fmt"
	"math/rand"
)

// Obj 是一个非常消耗资源的对象
type Obj struct {
	Name string
}

func (*Obj) Do() {
	// do something
}

// Pool 用来作为 obj 的对象池
type Pool chan *Obj

func NewObjPool(cap int) *Pool {
	p := make(Pool, cap)
	for i := 0; i < cap; i++ {
		obj := new(Obj)
		obj.Name = fmt.Sprintf("obj-%d", rand.Int())
		p <- obj
	}
	return &p
}

// Acquire 从对象池中获取一个对象
func (p Pool) Acquire() *Obj {
	return <-p
}

// Release 释放一个对象到对象池中
func (p Pool) Release(obj *Obj) {
	p <- obj
}
