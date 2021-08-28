package go_designpattern

import "fmt"

// 策略模式 Strategy Pattern
// 定义一组算法，将每个算法封装起来，使它们之间可以互换
// 通常可以和简单工厂模式组合起来使用，简单工厂模式用来创建具体的策略

type StrategyType int

const (
	StrategyTypeAdd    = 0
	StrategyTypeReduce = 1
)

// Strategy 策略接口
type Strategy interface {
	Exec(int, int) int
}

// Add 具体的策略实现
type Add struct {
}

func (*Add) Exec(a, b int) int {
	return a + b
}

// Reduce 具体的策略实现
type Reduce struct {
}

func (*Reduce) Exec(a, b int) int {
	return a - b
}

// Operator 策略的执行者
type Operator struct {
	strategy Strategy
}

// SetStrategy 给策略的执行者设置具体的策略
func (o *Operator) SetStrategy(strategy Strategy) {
	o.strategy = strategy
}

// Calculate 策略执行者执行策略
func (o *Operator) Calculate(a, b int) int {
	return o.strategy.Exec(a, b)
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
