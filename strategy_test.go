package go_designpattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperator_Calculate(t *testing.T) {
	strategy, err := NewStrategy(StrategyTypeAdd)
	assert.Nil(t, err)
	operator := new(Operator)
	operator.SetStrategy(strategy)
	assert.Equal(t, 10, operator.Operate(8, 2))
	strategy, err = NewStrategy(StrategyTypeReduce)
	assert.Nil(t, err)
	operator.SetStrategy(strategy)
	assert.Equal(t, 6, operator.Operate(8, 2))
}
