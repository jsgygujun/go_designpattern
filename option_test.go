package go_designpattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewConnection(t *testing.T) {
	conn, err := NewConnection("addr")
	assert.Nil(t, err)
	assert.Equal(t, 100*time.Millisecond, conn.timeout)
	assert.Equal(t, false, conn.caching)
	conn, err = NewConnection("addr", withTimeout(time.Second))
	assert.Nil(t, err)
	assert.Equal(t, time.Second, conn.timeout)
	assert.Equal(t, false, conn.caching)
	conn, err = NewConnection("addr", withCaching(true))
	assert.Nil(t, err)
	assert.Equal(t, 100*time.Millisecond, conn.timeout)
	assert.Equal(t, true, conn.caching)
	conn, err = NewConnection("addr", withTimeout(time.Second), withCaching(true))
	assert.Nil(t, err)
	assert.Equal(t, time.Second, conn.timeout)
	assert.Equal(t, true, conn.caching)
}
