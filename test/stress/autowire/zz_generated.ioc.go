//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package autowire

import (
	testingx "testing"

	ioc_golangautowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&ioc_golangautowire.StructDescriptor{
		Factory: func() interface{} {
			return &normalApp_{}
		},
	})
	normalAppStructDescriptor := &ioc_golangautowire.StructDescriptor{
		Factory: func() interface{} {
			return &NormalApp{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	normal.RegisterStructDescriptor(normalAppStructDescriptor)
	normal.RegisterStructDescriptor(&ioc_golangautowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceImpl1_{}
		},
	})
	serviceImpl1StructDescriptor := &ioc_golangautowire.StructDescriptor{
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
	normal.RegisterStructDescriptor(&ioc_golangautowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceImpl2_{}
		},
	})
	serviceImpl2StructDescriptor := &ioc_golangautowire.StructDescriptor{
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
	normal.RegisterStructDescriptor(&ioc_golangautowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceStruct_{}
		},
	})
	serviceStructStructDescriptor := &ioc_golangautowire.StructDescriptor{
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
	normal.RegisterStructDescriptor(&ioc_golangautowire.StructDescriptor{
		Factory: func() interface{} {
			return &singletonApp_{}
		},
	})
	singletonAppStructDescriptor := &ioc_golangautowire.StructDescriptor{
		Factory: func() interface{} {
			return &SingletonApp{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(singletonAppStructDescriptor)
}

type normalApp_ struct {
	RunTest_ func(t *testingx.T)
}

func (n *normalApp_) RunTest(t *testingx.T) {
	n.RunTest_(t)
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

type singletonApp_ struct {
	RunTest_ func(t *testingx.T)
}

func (s *singletonApp_) RunTest(t *testingx.T) {
	s.RunTest_(t)
}

type NormalAppIOCInterface interface {
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

type SingletonAppIOCInterface interface {
	RunTest(t *testingx.T)
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

var _singletonAppSDID string

func GetSingletonAppSingleton() (*SingletonApp, error) {
	if _singletonAppSDID == "" {
		_singletonAppSDID = util.GetSDIDByStructPtr(new(SingletonApp))
	}
	i, err := singleton.GetImpl(_singletonAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*SingletonApp)
	return impl, nil
}

func GetSingletonAppIOCInterfaceSingleton() (SingletonAppIOCInterface, error) {
	if _singletonAppSDID == "" {
		_singletonAppSDID = util.GetSDIDByStructPtr(new(SingletonApp))
	}
	i, err := singleton.GetImplWithProxy(_singletonAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(SingletonAppIOCInterface)
	return impl, nil
}

type ThisSingletonApp struct {
}

func (t *ThisSingletonApp) This() SingletonAppIOCInterface {
	thisPtr, _ := GetSingletonAppIOCInterfaceSingleton()
	return thisPtr
}
