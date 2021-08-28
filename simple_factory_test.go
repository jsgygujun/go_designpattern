package go_designpattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAnimal(t *testing.T) {
	cat, err := NewAnimal(AnimalTypeCat)
	assert.Nil(t, err)
	assert.Equal(t, "hello, I'm a cat", cat.SayHello())
	dog, err := NewAnimal(AnimalTypeDog)
	assert.Nil(t, err)
	assert.Equal(t, "hello, I'm a dog", dog.SayHello())
}
