//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package config

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	autowireconfig "github.com/alibaba/ioc-golang/extension/autowire/config"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configFloat64_{}
		},
	})
	configFloat64StructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(ConfigFloat64)
		},
		ParamFactory: func() interface{} {
			return new(ConfigFloat64)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configFloat64Interface)
			impl := i.(*ConfigFloat64)
			return param.new(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	autowireconfig.RegisterStructDescriptor(configFloat64StructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configInt64_{}
		},
	})
	configInt64StructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(ConfigInt64)
		},
		ParamFactory: func() interface{} {
			return new(ConfigInt64)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configInt64Interface)
			impl := i.(*ConfigInt64)
			return param.new(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	autowireconfig.RegisterStructDescriptor(configInt64StructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configInt_{}
		},
	})
	configIntStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(ConfigInt)
		},
		ParamFactory: func() interface{} {
			return new(ConfigInt)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configIntInterface)
			impl := i.(*ConfigInt)
			return param.new(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	autowireconfig.RegisterStructDescriptor(configIntStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configMap_{}
		},
	})
	configMapStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(ConfigMap)
		},
		ParamFactory: func() interface{} {
			return new(ConfigMap)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configMapInterface)
			impl := i.(*ConfigMap)
			return param.new(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	autowireconfig.RegisterStructDescriptor(configMapStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configSlice_{}
		},
	})
	configSliceStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(ConfigSlice)
		},
		ParamFactory: func() interface{} {
			return new(ConfigSlice)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configSliceInterface)
			impl := i.(*ConfigSlice)
			return param.new(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	autowireconfig.RegisterStructDescriptor(configSliceStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configString_{}
		},
	})
	configStringStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(ConfigString)
		},
		ParamFactory: func() interface{} {
			return new(ConfigString)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configStringInterface)
			impl := i.(*ConfigString)
			return param.new(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	autowireconfig.RegisterStructDescriptor(configStringStructDescriptor)
}

type configStringInterface interface {
	new(impl *ConfigString) (*ConfigString, error)
}
type configFloat64Interface interface {
	new(impl *ConfigFloat64) (*ConfigFloat64, error)
}
type configInt64Interface interface {
	new(impl *ConfigInt64) (*ConfigInt64, error)
}
type configIntInterface interface {
	new(impl *ConfigInt) (*ConfigInt, error)
}
type configMapInterface interface {
	new(impl *ConfigMap) (*ConfigMap, error)
}
type configSliceInterface interface {
	new(impl *ConfigSlice) (*ConfigSlice, error)
}
type configFloat64_ struct {
	Value_ func() float64
}

func (c *configFloat64_) Value() float64 {
	return c.Value_()
}

type configInt64_ struct {
	Value_ func() int64
}

func (c *configInt64_) Value() int64 {
	return c.Value_()
}

type configInt_ struct {
	Value_ func() int
}

func (c *configInt_) Value() int {
	return c.Value_()
}

type configMap_ struct {
	Value_ func() map[string]interface{}
}

func (c *configMap_) Value() map[string]interface{} {
	return c.Value_()
}

type configSlice_ struct {
	Value_ func() []interface{}
}

func (c *configSlice_) Value() []interface{} {
	return c.Value_()
}

type configString_ struct {
	Value_ func() string
}

func (c *configString_) Value() string {
	return c.Value_()
}

type ConfigFloat64IOCInterface interface {
	Value() float64
}

type ConfigInt64IOCInterface interface {
	Value() int64
}

type ConfigIntIOCInterface interface {
	Value() int
}

type ConfigMapIOCInterface interface {
	Value() map[string]interface{}
}

type ConfigSliceIOCInterface interface {
	Value() []interface{}
}

type ConfigStringIOCInterface interface {
	Value() string
}

var _configFloat64SDID string
var _configInt64SDID string
var _configIntSDID string
var _configMapSDID string
var _configSliceSDID string
var _configStringSDID string
