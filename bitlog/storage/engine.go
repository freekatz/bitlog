package storage

import (
	"github.com/1uvu/bitlog/storage/driver"
)

// Engine is a wrapper of Driver
type Engine struct {
	driver driver.Driver
}

func WithDefault() *Engine {
	return WithDriver(driver.Default())
}

func WithDriver(d driver.Driver) *Engine {
	return &Engine{d}
}

//func (e *Engine) StoreBlock(id string, b *types.BlockStatus) bool {
//
//	return true
//}
//
//func (e *Engine) StoreChain(id string, c *types.ChainStatus) bool {
//
//	return true
//}
//
//func (e *Engine) StoreFork(id string, f *types.ForkStatus) bool {
//
//	return true
//}

func (e *Engine) store(id, obj string) bool {
	return e.driver.Store(id, obj)
}

//func (e *Engine) LoadBlock(id string) (b *types.BlockStatus, ok bool) {
//
//	return b, ok
//}
//
//func (e *Engine) LoadChain(id string) (c *types.ChainStatus, ok bool) {
//
//	return c, ok
//}
//
//func (e *Engine) LoadFork(id string) (f *types.ForkStatus, ok bool) {
//
//	return f, ok
//}

func (e *Engine) load(id string) (obj interface{}, ok bool) {
	obj, ok = e.driver.Load(id)
	return obj, ok
}
