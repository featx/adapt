package web

import (
	"log"

	"github.com/gin-gonic/gin"
)

type GinEngine struct {
	delegate *gin.Engine
	config   Config
}

func NewGinEngine(config Config) Engine {
	return &GinEngine{gin.New(), config}
}

func (ginEngine *GinEngine) Start(args string) error {
	return ginEngine.delegate.Run(args)
}

func (ginEngine *GinEngine) ListenAndServe() error {
	//return ginEngine.delegate.TLSServer.Serve(ginEngine.delegate.TLSListener)
	return ginEngine.delegate.Run()
}

//SetLogger in GinEngine
func (ginEngine *GinEngine) SetLogger(logger *log.Logger) {

}

//ToContext Convert atreugo RequestCtx to Context
func fromGinContext(ctx *gin.Context) Context {
	return nil
}

//Use in GinEngine
func (ginEngine *GinEngine) Use(processArray ...Process) {
	middlewares := make([]gin.HandlerFunc, len(processArray))
	for i, process := range processArray {
		middlewares[i] = func(ctx *gin.Context) {
			err := process(fromGinContext(ctx))
			if err != nil {
				panic(err)
			}
		}
	}
	ginEngine.delegate.Use(middlewares...)
}

//Group in GinEngine
func (ginEngine *GinEngine) Group(path string, processArray ...Process) RouterGroup {
	middlewares := make([]gin.HandlerFunc, len(processArray))
	for i, process := range processArray {
		middlewares[i] = func(ctx *gin.Context) {
			err := process(fromGinContext(ctx))
			if err != nil {
				panic(err)
			}
		}
	}
	//ginRouterGroup
	_ = ginEngine.delegate.Group(path, middlewares...)
	return nil
}

//GET in GinEngine
func (ginEngine *GinEngine) GET(path string, process Process) {
	ginEngine.delegate.GET(path, func(ctx *gin.Context) {
		err := process(fromGinContext(ctx))
		if err != nil {
			panic(err)
		}
	})
}

//POST in GinEngine
func (ginEngine *GinEngine) POST(path string, process Process) {
	ginEngine.delegate.POST(path, func(ctx *gin.Context) {
		err := process(fromGinContext(ctx))
		if err != nil {
			panic(err)
		}
	})
}

//PUT in GinEngine
func (ginEngine *GinEngine) PUT(path string, process Process) {
	ginEngine.delegate.PUT(path, func(ctx *gin.Context) {
		err := process(fromGinContext(ctx))
		if err != nil {
			panic(err)
		}
	})
}

//DELETE in GinEngine
func (ginEngine *GinEngine) DELETE(path string, process Process) {
	ginEngine.delegate.DELETE(path, func(ctx *gin.Context) {
		err := process(fromGinContext(ctx))
		if err != nil {
			panic(err)
		}
	})
}

//PATCH in GinEngine
func (ginEngine *GinEngine) PATCH(path string, process Process) {
	ginEngine.delegate.PATCH(path, func(ctx *gin.Context) {
		err := process(fromGinContext(ctx))
		if err != nil {
			panic(err)
		}
	})
}

//HEAD in GinEngine
func (ginEngine *GinEngine) HEAD(path string, process Process) {
	ginEngine.delegate.HEAD(path, func(ctx *gin.Context) {
		err := process(fromGinContext(ctx))
		if err != nil {
			panic(err)
		}
	})
}

//OPTIONS in GinEngine
func (ginEngine *GinEngine) OPTIONS(path string, process Process) {
	ginEngine.delegate.OPTIONS(path, func(ctx *gin.Context) {
		err := process(fromGinContext(ctx))
		if err != nil {
			panic(err)
		}
	})
}

//TRACE in GinEngine
func (ginEngine *GinEngine) TRACE(path string, process Process) {
	ginEngine.delegate.Handle("trace", path, func(ctx *gin.Context) {
		err := process(fromGinContext(ctx))
		if err != nil {
			panic(err)
		}
	})
}

//CONNECT in GinEngine
func (ginEngine *GinEngine) CONNECT(path string, process Process) {
	ginEngine.delegate.Handle("connect", path, func(ctx *gin.Context) {
		err := process(fromGinContext(ctx))
		if err != nil {
			panic(err)
		}
	})
}
