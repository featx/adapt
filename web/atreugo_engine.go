package web

import (
	"log"

	"github.com/savsgio/atreugo/v11"
)

// AtreugoEngine The atreugo implementation of engine
type AtreugoEngine struct {
	delegate *atreugo.Atreugo
	config   Config
}

//NewAtreugoEngine alloc new mem for atreugo impl engine from map config and give out the pointer
func NewAtreugoEngine(config Config) Engine {
	return &AtreugoEngine{atreugo.New(toAtreugoConfig(config)), config}
}

func toAtreugoConfig(config Config) atreugo.Config {
	return atreugo.Config{}
}

func (atreugoEngine *AtreugoEngine) Start(args string) error {
	config := toAtreugoConfig(atreugoEngine.config)
	config.Addr = args
	atreugoEngine.delegate = atreugo.New(config)
	return atreugoEngine.delegate.ListenAndServe()
}

func (atreugoEngine *AtreugoEngine) ListenAndServe() error {
	return atreugoEngine.delegate.ListenAndServe()
}

//SetLogger in AtreugoEngine
func (atreugoEngine *AtreugoEngine) SetLogger(logger *log.Logger) {
	atreugoEngine.delegate.SetLogOutput(logger.Writer())
}

//Use in AtreugoEngine
func (atreugoEngine *AtreugoEngine) Use(processArray ...Process) {
	processes := make([]atreugo.Middleware, len(processArray))
	for i, process := range processArray {
		processes[i] = func(ctx *atreugo.RequestCtx) error {
			return process(fromAtreugoCtx(ctx))
		}
	}
	atreugoEngine.delegate.UseBefore(processes...)
}

//Group in AtreugoEngine
func (atreugoEngine *AtreugoEngine) Group(path string, processArray ...Process) RouterGroup {
	atreugoEngine.delegate.NewGroupPath(path)
	return nil
}

//ToContext Convert atreugo RequestCtx to Context
func fromAtreugoCtx(ctx *atreugo.RequestCtx) Context {
	return nil
}

//GET in AtreugoEngine
func (atreugoEngine *AtreugoEngine) GET(path string, process Process) {
	atreugoEngine.delegate.GET(path, func(ctx *atreugo.RequestCtx) error {
		return process(fromAtreugoCtx(ctx))
	})
}

//POST in AtreugoEngine
func (atreugoEngine *AtreugoEngine) POST(path string, process Process) {
	atreugoEngine.delegate.POST(path, func(ctx *atreugo.RequestCtx) error {
		return process(fromAtreugoCtx(ctx))
	})
}

//PUT in AtreugoEngine
func (atreugoEngine *AtreugoEngine) PUT(path string, process Process) {
	atreugoEngine.delegate.PUT(path, func(ctx *atreugo.RequestCtx) error {
		return process(fromAtreugoCtx(ctx))
	})
}

//DELETE in AtreugoEngine
func (atreugoEngine *AtreugoEngine) DELETE(path string, process Process) {
	atreugoEngine.delegate.DELETE(path, func(ctx *atreugo.RequestCtx) error {
		return process(fromAtreugoCtx(ctx))
	})
}

//PATCH in AtreugoEngine
func (atreugoEngine *AtreugoEngine) PATCH(path string, process Process) {
	atreugoEngine.delegate.PATCH(path, func(ctx *atreugo.RequestCtx) error {
		return process(fromAtreugoCtx(ctx))
	})
}

//HEAD in AtreugoEngine
func (atreugoEngine *AtreugoEngine) HEAD(path string, process Process) {
	atreugoEngine.delegate.HEAD(path, func(ctx *atreugo.RequestCtx) error {
		return process(fromAtreugoCtx(ctx))
	})
}

//OPTIONS in AtreugoEngine
func (atreugoEngine *AtreugoEngine) OPTIONS(path string, process Process) {
	atreugoEngine.delegate.OPTIONS(path, func(ctx *atreugo.RequestCtx) error {
		return process(fromAtreugoCtx(ctx))
	})
}

//TRACE in AtreugoEngine
func (atreugoEngine *AtreugoEngine) TRACE(path string, process Process) {
	atreugoEngine.delegate.Path("TRACE", path, func(ctx *atreugo.RequestCtx) error {
		return process(fromAtreugoCtx(ctx))
	})
}

//CONNECT in AtreugoEngine
func (atreugoEngine *AtreugoEngine) CONNECT(path string, process Process) {
	atreugoEngine.delegate.Path("CONNECT", path, func(ctx *atreugo.RequestCtx) error {
		return process(fromAtreugoCtx(ctx))
	})
}
