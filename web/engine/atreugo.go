package engine

import (
	"github.com/featx/goin/web/types"
	"log"
	"net"

	"github.com/featx/goin/web/context"
	"github.com/savsgio/atreugo/v11"
)

// AtreugoEngine The atreugo implementation of engine
type AtreugoEngine struct {
	delegate *atreugo.Atreugo
	config   types.Config
}

//NewAtreugoEngine alloc new mem for atreugo impl engine from map config and give out the pointer
func NewAtreugoEngine(config types.Config) types.Engine {
	return &AtreugoEngine{atreugo.New(toAtreugoConfig(config)), config}
}

func toAtreugoConfig(config types.Config) atreugo.Config {
	return atreugo.Config{}
}

func (atreugoEngine *AtreugoEngine) Start(args string) error {
	l, err := net.Listen("tcp", args)
	if err != nil {
		return err
	}
	return atreugoEngine.delegate.Serve(l)
}

func (atreugoEngine *AtreugoEngine) ListenAndServe() error {
	return atreugoEngine.delegate.ListenAndServe()
}

//SetLogger in AtreugoEngine
func (atreugoEngine *AtreugoEngine) SetLogger(logger *log.Logger) {
	atreugoEngine.delegate.SetLogOutput(logger.Writer())
}

//Use in AtreugoEngine
func (atreugoEngine *AtreugoEngine) Use(processArray ...types.Process) {
	processes := make([]atreugo.Middleware, len(processArray))
	for i, process := range processArray {
		processes[i] = func(ctx *atreugo.RequestCtx) error {
			return process(context.FromAtreugoCtx(ctx))
		}
	}
	atreugoEngine.delegate.UseBefore(processes...)
}

//Group in AtreugoEngine
func (atreugoEngine *AtreugoEngine) Group(path string, processArray ...types.Process) types.RouterGroup {
	atreugoEngine.delegate.NewGroupPath(path)
	return nil
}

//GET in AtreugoEngine
func (atreugoEngine *AtreugoEngine) GET(path string, process types.Process) {
	atreugoEngine.delegate.GET(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//POST in AtreugoEngine
func (atreugoEngine *AtreugoEngine) POST(path string, process types.Process) {
	atreugoEngine.delegate.POST(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//PUT in AtreugoEngine
func (atreugoEngine *AtreugoEngine) PUT(path string, process types.Process) {
	atreugoEngine.delegate.PUT(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//DELETE in AtreugoEngine
func (atreugoEngine *AtreugoEngine) DELETE(path string, process types.Process) {
	atreugoEngine.delegate.DELETE(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//PATCH in AtreugoEngine
func (atreugoEngine *AtreugoEngine) PATCH(path string, process types.Process) {
	atreugoEngine.delegate.PATCH(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//HEAD in AtreugoEngine
func (atreugoEngine *AtreugoEngine) HEAD(path string, process types.Process) {
	atreugoEngine.delegate.HEAD(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//OPTIONS in AtreugoEngine
func (atreugoEngine *AtreugoEngine) OPTIONS(path string, process types.Process) {
	atreugoEngine.delegate.OPTIONS(path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//TRACE in AtreugoEngine
func (atreugoEngine *AtreugoEngine) TRACE(path string, process types.Process) {
	atreugoEngine.delegate.Path("TRACE", path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}

//CONNECT in AtreugoEngine
func (atreugoEngine *AtreugoEngine) CONNECT(path string, process types.Process) {
	atreugoEngine.delegate.Path("CONNECT", path, func(ctx *atreugo.RequestCtx) error {
		return process(context.FromAtreugoCtx(ctx))
	})
}
