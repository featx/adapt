package web

import (
	"github.com/savsgio/atreugo/v10"
	"log"
)

type AtreugoEngine struct {
	delegate *atreugo.Atreugo
}

func NewAtreugoEngine(config map[string]interface{}) *AtreugoEngine {
	return &AtreugoEngine{atreugo.New(atreugo.Config{})}
}

func ToContext(ctx *atreugo.RequestCtx) Context {
	return nil //TODO
}

func (atreugoEngine *AtreugoEngine) GET(path string, middlewareFunc MiddlewareFunc) {
	atreugoEngine.delegate.GET(path, func(ctx *atreugo.RequestCtx) error {
		return middlewareFunc(ToContext(ctx))
	})
}

func (atreugoEngine *AtreugoEngine) POST(path string, middlewareFunc MiddlewareFunc) {
	atreugoEngine.delegate.POST(path, func(ctx *atreugo.RequestCtx) error {
		return middlewareFunc(ToContext(ctx))
	})
}

func (atreugoEngine *AtreugoEngine) PUT(path string, middlewareFunc MiddlewareFunc) {
	atreugoEngine.delegate.PUT(path, func(ctx *atreugo.RequestCtx) error {
		return middlewareFunc(ToContext(ctx))
	})
}

func (atreugoEngine *AtreugoEngine) DELETE(path string, middlewareFunc MiddlewareFunc) {
	atreugoEngine.delegate.DELETE(path, func(ctx *atreugo.RequestCtx) error {
		return middlewareFunc(ToContext(ctx))
	})
}

func (atreugoEngine *AtreugoEngine) PATCH(path string, middlewareFunc MiddlewareFunc) {
	atreugoEngine.delegate.PATCH(path, func(ctx *atreugo.RequestCtx) error {
		return middlewareFunc(ToContext(ctx))
	})
}

func (atreugoEngine *AtreugoEngine) HEAD(path string, middlewareFunc MiddlewareFunc) {
	atreugoEngine.delegate.HEAD(path, func(ctx *atreugo.RequestCtx) error {
		return middlewareFunc(ToContext(ctx))
	})
}

func (atreugoEngine *AtreugoEngine) OPTIONS(path string, middlewareFunc MiddlewareFunc) {
	atreugoEngine.delegate.OPTIONS(path, func(ctx *atreugo.RequestCtx) error {
		return middlewareFunc(ToContext(ctx))
	})
}

func (atreugoEngine *AtreugoEngine) TRACE(path string, middlewareFunc MiddlewareFunc) {
	atreugoEngine.delegate.Path("TRACE", path, func(ctx *atreugo.RequestCtx) error {
		return middlewareFunc(ToContext(ctx))
	})
}

func (atreugoEngine *AtreugoEngine) CONNECT(path string, middlewareFunc MiddlewareFunc) {
	atreugoEngine.delegate.Path("CONNECT", path, func(ctx *atreugo.RequestCtx) error {
		return middlewareFunc(ToContext(ctx))
	})
}

func (atreugoEngine *AtreugoEngine) SetLogger(logger *log.Logger) {

}

func (atreugoEngine *AtreugoEngine) Use(middlewareArray ...MiddlewareFunc) {

}

func (atreugo *AtreugoEngine) Group(path string, middlewareArray ...MiddlewareFunc) RouterGroup {
	return nil
}
