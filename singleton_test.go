package go_designpattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSingletonIns(t *testing.T) {
	assert.Equal(t, GetSingletonIns(), GetSingletonIns())
}
