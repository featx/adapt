package engine

import (
	"log"

	"github.com/savsgio/atreugo/v11"

	"github.com/featx/goin/web"
	"github.com/featx/goin/web/context"
)

// AtreugoEngine The atreugo implementation of engine
type AtreugoEngine struct {
	delegate *atreugo.Atreugo
	config   web.Config
}

//NewAtreugoEngine alloc new mem for atreugo impl engine from map config and give out the pointer
func NewAtreugoEngine(config web.Config) web.Engine {
	return &AtreugoEngine{atreugo.New(toAtreugoConfig(config)), config}
}

func toAtreugoConfig(config web.Config) atreugo.Config {
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
func (atreugoEngine *AtreugoEngine) Use(processArray ...web.Process) {
	processes := make([]atreugo.Middleware, len(processArray))
	for i, process := range processArray {
		processes[i] = func(ctx *atreugo.RequestCtx) error {
			return process(context.FromAtreugoCtx(ctx))
		}
	}
	atreugoEngine.delegate.UseBefore(processes...)
}

//Group in AtreugoEngine
func (atreugoEngine *AtreugoEngine) Group(path string, processArray ...web.Process) web.RouterGroup {
	atreugoEngine.delegate.NewGroupPath(path)
	return nil
}

//GET in AtreugoEngine
func (atreugoEngine *AtreugoEngine) GET(path string, process web.Process) {
	atreugoEngine.delegate.GET(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//POST in AtreugoEngine
func (atreugoEngine *AtreugoEngine) POST(path string, process web.Process) {
	atreugoEngine.delegate.POST(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//PUT in AtreugoEngine
func (atreugoEngine *AtreugoEngine) PUT(path string, process web.Process) {
	atreugoEngine.delegate.PUT(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//DELETE in AtreugoEngine
func (atreugoEngine *AtreugoEngine) DELETE(path string, process web.Process) {
	atreugoEngine.delegate.DELETE(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//PATCH in AtreugoEngine
func (atreugoEngine *AtreugoEngine) PATCH(path string, process web.Process) {
	atreugoEngine.delegate.PATCH(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//HEAD in AtreugoEngine
func (atreugoEngine *AtreugoEngine) HEAD(path string, process web.Process) {
	atreugoEngine.delegate.HEAD(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//OPTIONS in AtreugoEngine
func (atreugoEngine *AtreugoEngine) OPTIONS(path string, process web.Process) {
	atreugoEngine.delegate.OPTIONS(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//TRACE in AtreugoEngine
func (atreugoEngine *AtreugoEngine) TRACE(path string, process web.Process) {
	atreugoEngine.delegate.Path("TRACE", path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//CONNECT in AtreugoEngine
func (atreugoEngine *AtreugoEngine) CONNECT(path string, process web.Process) {
	atreugoEngine.delegate.Path("CONNECT", path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}
