package go_designpattern

import "fmt"

// 策略模式 Strategy Pattern
// 定义一组算法，将每个算法封装起来，使它们之间可以互换
// 通常可以和简单工厂模式组合起来使用，简单工厂模式用来创建具体的策略

// Rules of Thumb 经验法则
// 1. 策略模式和模版模式类似，只是粒度不同
// 2. 策略模式让你改变一个对象的内脏，装饰者模式让你改变一个对象的皮肤

type StrategyType int

const (
	StrategyTypeAdd = iota + 1
	StrategyTypeReduce
)

// Strategy 策略接口
type Strategy interface {
	Apply(int, int) int
}

// Operator 策略的执行者
type Operator struct {
	strategy Strategy
}

// SetStrategy 设置具体的策略
func (o *Operator) SetStrategy(strategy Strategy) {
	o.strategy = strategy
}

// Operate 策略执行者执行策略
func (o *Operator) Operate(a, b int) int {
	return o.strategy.Apply(a, b)
}

// Add 具体的策略实现
type Add struct {
}

func (*Add) Apply(a, b int) int {
	return a + b
}

// Reduce 具体的策略实现
type Reduce struct {
}

func (*Reduce) Apply(a, b int) int {
	return a - b
}

// NewStrategy 根据策略类型创建具体的策略
func NewStrategy(strategyType StrategyType) (Strategy, error) {
	switch strategyType {
	case StrategyTypeAdd:
		return &Add{}, nil
	case StrategyTypeReduce:
		return &Reduce{}, nil
	default:
		return nil, fmt.Errorf("unknown stragetyType: %v", strategyType)
	}
}
