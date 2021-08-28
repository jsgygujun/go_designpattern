package go_designpattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewObjPool(t *testing.T) {
	objPool := NewObjPool(1)
	obj1 := objPool.Acquire()
	objPool.Release(obj1)
	obj2 := objPool.Acquire()
	objPool.Release(obj2)
	assert.Equal(t, obj1.Name, obj2.Name)

	objPool = NewObjPool(2)
	obj1 = objPool.Acquire()
	objPool.Release(obj1)
	obj2 = objPool.Acquire()
	objPool.Release(obj2)
	assert.NotEqual(t, obj1.Name, obj2.Name)
}
