package web

import (
	"log"

	"github.com/savsgio/atreugo/v10"
)

// AtreugoEngine The atreugo implementation of engine
type AtreugoEngine struct {
	delegate *atreugo.Atreugo
}

//NewAtreugoEngine alloc new mem for atreugo impl engine from map config and give out the pointer
func NewAtreugoEngine(config map[string]interface{}) *AtreugoEngine {
	return &AtreugoEngine{atreugo.New(atreugo.Config{})}
}

//ToContext Convert atreugo RequestCtx to Context
func ToContext(ctx *atreugo.RequestCtx) Context {
	return nil //TODO
}

//GET in AtreugoEngine
func (atreugoEngine *AtreugoEngine) GET(path string, process Process) {
	atreugoEngine.delegate.GET(path, func(ctx *atreugo.RequestCtx) error {
		return process(ToContext(ctx))
	})
}

//POST in AtreugoEngine
func (atreugoEngine *AtreugoEngine) POST(path string, process Process) {
	atreugoEngine.delegate.POST(path, func(ctx *atreugo.RequestCtx) error {
		return process(ToContext(ctx))
	})
}

//PUT in AtreugoEngine
func (atreugoEngine *AtreugoEngine) PUT(path string, process Process) {
	atreugoEngine.delegate.PUT(path, func(ctx *atreugo.RequestCtx) error {
		return process(ToContext(ctx))
	})
}

//DELETE in AtreugoEngine
func (atreugoEngine *AtreugoEngine) DELETE(path string, process Process) {
	atreugoEngine.delegate.DELETE(path, func(ctx *atreugo.RequestCtx) error {
		return process(ToContext(ctx))
	})
}

//PATCH in AtreugoEngine
func (atreugoEngine *AtreugoEngine) PATCH(path string, process Process) {
	atreugoEngine.delegate.PATCH(path, func(ctx *atreugo.RequestCtx) error {
		return process(ToContext(ctx))
	})
}

//HEAD in AtreugoEngine
func (atreugoEngine *AtreugoEngine) HEAD(path string, process Process) {
	atreugoEngine.delegate.HEAD(path, func(ctx *atreugo.RequestCtx) error {
		return process(ToContext(ctx))
	})
}

//OPTIONS in AtreugoEngine
func (atreugoEngine *AtreugoEngine) OPTIONS(path string, process Process) {
	atreugoEngine.delegate.OPTIONS(path, func(ctx *atreugo.RequestCtx) error {
		return process(ToContext(ctx))
	})
}

//TRACE in AtreugoEngine
func (atreugoEngine *AtreugoEngine) TRACE(path string, process Process) {
	atreugoEngine.delegate.Path("TRACE", path, func(ctx *atreugo.RequestCtx) error {
		return process(ToContext(ctx))
	})
}

//CONNECT in AtreugoEngine
func (atreugoEngine *AtreugoEngine) CONNECT(path string, process Process) {
	atreugoEngine.delegate.Path("CONNECT", path, func(ctx *atreugo.RequestCtx) error {
		return process(ToContext(ctx))
	})
}

//SetLogger in AtreugoEngine
func (atreugoEngine *AtreugoEngine) SetLogger(logger *log.Logger) {

}

//Use in AtreugoEngine
func (atreugoEngine *AtreugoEngine) Use(processArray ...Process) {

}

//Group in AtreugoEngine
func (atreugoEngine *AtreugoEngine) Group(path string, processArray ...Process) RouterGroup {
	return nil
}
