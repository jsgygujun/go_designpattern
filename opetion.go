package go_designpattern

import "time"

// 选项模式, 一般用在库设计中, 提供给调用方可选的配置，调用方使用方式如下:
// cli, err := NewXXX(xxx, withXXX(), withXXX())
// 其中 xxx 是必选配置, withXXX 则是可选配置
const (
	_defaultTime    = 100 * time.Millisecond
	_defaultCaching = false
)

type Connection struct {
	addr    string
	caching bool
	timeout time.Duration
}

// options 包含可选属性
type options struct {
	timeout time.Duration
	caching bool
}

// Option 接口, 包含一个 apply 方法, 该方法主要是将属性设置到 options 结构体中
type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

// 设置 timeout 属性
func withTimeout(timeout time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = timeout
	})
}

// 设置 caching 属性
func withCaching(caching bool) Option {
	return optionFunc(func(o *options) {
		o.caching = caching
	})
}

// NewConnection 根据参数创建 Connection, 参数包括必选参数和可选参数
func NewConnection(addr string, opts ...Option) (*Connection, error) {
	// 首先创建带有默认属性的 options 结构体
	options := options{
		timeout: _defaultTime,
		caching: _defaultCaching,
	}
	// 根据提供的 opts 参数修改 options 结构体
	for _, o := range opts {
		o.apply(&options)
	}
	return &Connection{
		addr:    addr,
		caching: options.caching,
		timeout: options.timeout,
	}, nil
}
