//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package aop

import (
	testingx "testing"

	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	"github.com/alibaba/ioc-golang/autowire/singleton"
	"github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &normalApp_{}
		},
	})
	normalAppStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &NormalApp{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	normal.RegisterStructDescriptor(normalAppStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &recursiveApp_{}
		},
	})
	recursiveAppStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &RecursiveApp{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(recursiveAppStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceImpl1_{}
		},
	})
	serviceImpl1StructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ServiceImpl1{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(serviceImpl1StructDescriptor)
	normal.RegisterStructDescriptor(serviceImpl1StructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceImpl2_{}
		},
	})
	serviceImpl2StructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ServiceImpl2{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(serviceImpl2StructDescriptor)
	normal.RegisterStructDescriptor(serviceImpl2StructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceStruct_{}
		},
	})
	serviceStructStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ServiceStruct{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(serviceStructStructDescriptor)
	normal.RegisterStructDescriptor(serviceStructStructDescriptor)
}

type normalApp_ struct {
	RunTest_ func(t *testingx.T)
}

func (n *normalApp_) RunTest(t *testingx.T) {
	n.RunTest_(t)
}

type recursiveApp_ struct {
	Reset_   func()
	RunTest_ func(t *testingx.T)
}

func (r *recursiveApp_) Reset() {
	r.Reset_()
}

func (r *recursiveApp_) RunTest(t *testingx.T) {
	r.RunTest_(t)
}

type serviceImpl1_ struct {
	GetHelloString_ func(name string) string
}

func (s *serviceImpl1_) GetHelloString(name string) string {
	return s.GetHelloString_(name)
}

type serviceImpl2_ struct {
	GetHelloString_ func(name string) string
}

func (s *serviceImpl2_) GetHelloString(name string) string {
	return s.GetHelloString_(name)
}

type serviceStruct_ struct {
	GetString_ func(name string) string
}

func (s *serviceStruct_) GetString(name string) string {
	return s.GetString_(name)
}

type NormalAppIOCInterface interface {
	RunTest(t *testingx.T)
}

type RecursiveAppIOCInterface interface {
	Reset()
	RunTest(t *testingx.T)
}

type ServiceImpl1IOCInterface interface {
	GetHelloString(name string) string
}

type ServiceImpl2IOCInterface interface {
	GetHelloString(name string) string
}

type ServiceStructIOCInterface interface {
	GetString(name string) string
}

var _normalAppSDID string

func GetNormalApp() (*NormalApp, error) {
	if _normalAppSDID == "" {
		_normalAppSDID = util.GetSDIDByStructPtr(new(NormalApp))
	}
	i, err := normal.GetImpl(_normalAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*NormalApp)
	return impl, nil
}

func GetNormalAppIOCInterface() (NormalAppIOCInterface, error) {
	if _normalAppSDID == "" {
		_normalAppSDID = util.GetSDIDByStructPtr(new(NormalApp))
	}
	i, err := normal.GetImplWithProxy(_normalAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(NormalAppIOCInterface)
	return impl, nil
}

var _recursiveAppSDID string

func GetRecursiveAppSingleton() (*RecursiveApp, error) {
	if _recursiveAppSDID == "" {
		_recursiveAppSDID = util.GetSDIDByStructPtr(new(RecursiveApp))
	}
	i, err := singleton.GetImpl(_recursiveAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*RecursiveApp)
	return impl, nil
}

func GetRecursiveAppIOCInterfaceSingleton() (RecursiveAppIOCInterface, error) {
	if _recursiveAppSDID == "" {
		_recursiveAppSDID = util.GetSDIDByStructPtr(new(RecursiveApp))
	}
	i, err := singleton.GetImplWithProxy(_recursiveAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(RecursiveAppIOCInterface)
	return impl, nil
}

type ThisRecursiveApp struct {
}

func (t *ThisRecursiveApp) This() RecursiveAppIOCInterface {
	thisPtr, _ := GetRecursiveAppIOCInterfaceSingleton()
	return thisPtr
}

var _serviceImpl1SDID string

func GetServiceImpl1Singleton() (*ServiceImpl1, error) {
	if _serviceImpl1SDID == "" {
		_serviceImpl1SDID = util.GetSDIDByStructPtr(new(ServiceImpl1))
	}
	i, err := singleton.GetImpl(_serviceImpl1SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ServiceImpl1)
	return impl, nil
}

func GetServiceImpl1IOCInterfaceSingleton() (ServiceImpl1IOCInterface, error) {
	if _serviceImpl1SDID == "" {
		_serviceImpl1SDID = util.GetSDIDByStructPtr(new(ServiceImpl1))
	}
	i, err := singleton.GetImplWithProxy(_serviceImpl1SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ServiceImpl1IOCInterface)
	return impl, nil
}

type ThisServiceImpl1 struct {
}

func (t *ThisServiceImpl1) This() ServiceImpl1IOCInterface {
	thisPtr, _ := GetServiceImpl1IOCInterfaceSingleton()
	return thisPtr
}

func GetServiceImpl1() (*ServiceImpl1, error) {
	if _serviceImpl1SDID == "" {
		_serviceImpl1SDID = util.GetSDIDByStructPtr(new(ServiceImpl1))
	}
	i, err := normal.GetImpl(_serviceImpl1SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ServiceImpl1)
	return impl, nil
}

func GetServiceImpl1IOCInterface() (ServiceImpl1IOCInterface, error) {
	if _serviceImpl1SDID == "" {
		_serviceImpl1SDID = util.GetSDIDByStructPtr(new(ServiceImpl1))
	}
	i, err := normal.GetImplWithProxy(_serviceImpl1SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ServiceImpl1IOCInterface)
	return impl, nil
}

var _serviceImpl2SDID string

func GetServiceImpl2Singleton() (*ServiceImpl2, error) {
	if _serviceImpl2SDID == "" {
		_serviceImpl2SDID = util.GetSDIDByStructPtr(new(ServiceImpl2))
	}
	i, err := singleton.GetImpl(_serviceImpl2SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ServiceImpl2)
	return impl, nil
}

func GetServiceImpl2IOCInterfaceSingleton() (ServiceImpl2IOCInterface, error) {
	if _serviceImpl2SDID == "" {
		_serviceImpl2SDID = util.GetSDIDByStructPtr(new(ServiceImpl2))
	}
	i, err := singleton.GetImplWithProxy(_serviceImpl2SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ServiceImpl2IOCInterface)
	return impl, nil
}

type ThisServiceImpl2 struct {
}

func (t *ThisServiceImpl2) This() ServiceImpl2IOCInterface {
	thisPtr, _ := GetServiceImpl2IOCInterfaceSingleton()
	return thisPtr
}

func GetServiceImpl2() (*ServiceImpl2, error) {
	if _serviceImpl2SDID == "" {
		_serviceImpl2SDID = util.GetSDIDByStructPtr(new(ServiceImpl2))
	}
	i, err := normal.GetImpl(_serviceImpl2SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ServiceImpl2)
	return impl, nil
}

func GetServiceImpl2IOCInterface() (ServiceImpl2IOCInterface, error) {
	if _serviceImpl2SDID == "" {
		_serviceImpl2SDID = util.GetSDIDByStructPtr(new(ServiceImpl2))
	}
	i, err := normal.GetImplWithProxy(_serviceImpl2SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ServiceImpl2IOCInterface)
	return impl, nil
}

var _serviceStructSDID string

func GetServiceStructSingleton() (*ServiceStruct, error) {
	if _serviceStructSDID == "" {
		_serviceStructSDID = util.GetSDIDByStructPtr(new(ServiceStruct))
	}
	i, err := singleton.GetImpl(_serviceStructSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ServiceStruct)
	return impl, nil
}

func GetServiceStructIOCInterfaceSingleton() (ServiceStructIOCInterface, error) {
	if _serviceStructSDID == "" {
		_serviceStructSDID = util.GetSDIDByStructPtr(new(ServiceStruct))
	}
	i, err := singleton.GetImplWithProxy(_serviceStructSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ServiceStructIOCInterface)
	return impl, nil
}

type ThisServiceStruct struct {
}

func (t *ThisServiceStruct) This() ServiceStructIOCInterface {
	thisPtr, _ := GetServiceStructIOCInterfaceSingleton()
	return thisPtr
}

func GetServiceStruct() (*ServiceStruct, error) {
	if _serviceStructSDID == "" {
		_serviceStructSDID = util.GetSDIDByStructPtr(new(ServiceStruct))
	}
	i, err := normal.GetImpl(_serviceStructSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ServiceStruct)
	return impl, nil
}

func GetServiceStructIOCInterface() (ServiceStructIOCInterface, error) {
	if _serviceStructSDID == "" {
		_serviceStructSDID = util.GetSDIDByStructPtr(new(ServiceStruct))
	}
	i, err := normal.GetImplWithProxy(_serviceStructSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ServiceStructIOCInterface)
	return impl, nil
}
