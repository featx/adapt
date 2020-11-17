package engine

import (
	"github.com/featx/goin/web"
	"github.com/labstack/echo/v4"
	glog "github.com/labstack/gommon/log"
	"log"
)

//EchoEngine The echo implementation of engine
type EchoEngine struct {
	delegate *echo.Echo
	config   web.Config
}

//NewEchoEngine The Echo implementation of engine
func NewEchoEngine(config web.Config) web.Engine {
	return &EchoEngine{echo.New(), config}
}

func (echoEngine *EchoEngine) Start(args string) error {
	return echoEngine.delegate.Start(args)
}

func (echoEngine *EchoEngine) ListenAndServe() error {
	//return echoEngine.delegate.TLSServer.Serve(echoEngine.delegate.TLSListener)
	return echoEngine.delegate.Server.Serve(echoEngine.delegate.Listener)
}

//SetLogger in EchoEngine
func (echoEngine *EchoEngine) SetLogger(logger *log.Logger) {
	echoEngine.delegate.Logger = glog.New(logger.Prefix())
	echoEngine.delegate.Logger.SetOutput(logger.Writer())
	//echoEngine.delegate.Logger.SetHeader()
	lev := uint8(logger.Flags()) //TODO
	echoEngine.delegate.Logger.SetLevel(glog.Lvl(lev))
}

//ToContext Convert atreugo RequestCtx to Context
func fromEchoContext(ctx echo.Context) web.Context {
	return nil
}

//Use in EchoEngine
func (echoEngine *EchoEngine) Use(processArray ...web.Process) {
	middlewares := make([]echo.MiddlewareFunc, len(processArray))
	for i, process := range processArray {
		middlewares[i] = func(handerFunc echo.HandlerFunc) echo.HandlerFunc {
			return func(ctx echo.Context) error {
				err := process(fromEchoContext(ctx))
				if err == nil {
					return handerFunc(ctx)
				}
				return err
			}
		}
	}
	echoEngine.delegate.Use(middlewares...)
}

//Group in EchoEngine
func (echoEngine *EchoEngine) Group(path string, processArray ...web.Process) web.RouterGroup {
	return nil
}

//GET in EchoEngine
func (echoEngine *EchoEngine) GET(path string, process web.Process) {
	echoEngine.delegate.GET(path, func(ctx echo.Context) error {
		return process(fromEchoContext(ctx))
	})
}

//POST in EchoEngine
func (echoEngine *EchoEngine) POST(path string, process web.Process) {
	echoEngine.delegate.POST(path, func(ctx echo.Context) error {
		return process(fromEchoContext(ctx))
	})
}

//PUT in EchoEngine
func (echoEngine *EchoEngine) PUT(path string, process web.Process) {
	echoEngine.delegate.PUT(path, func(ctx echo.Context) error {
		return process(fromEchoContext(ctx))
	})
}

//DELETE in EchoEngine
func (echoEngine *EchoEngine) DELETE(path string, process web.Process) {
	echoEngine.delegate.DELETE(path, func(ctx echo.Context) error {
		return process(fromEchoContext(ctx))
	})
}

//PATCH in EchoEngine
func (echoEngine *EchoEngine) PATCH(path string, process web.Process) {
	echoEngine.delegate.PATCH(path, func(ctx echo.Context) error {
		return process(fromEchoContext(ctx))
	})
}

//HEAD in EchoEngine
func (echoEngine *EchoEngine) HEAD(path string, process web.Process) {
	echoEngine.delegate.HEAD(path, func(ctx echo.Context) error {
		return process(fromEchoContext(ctx))
	})
}

//OPTIONS in EchoEngine
func (echoEngine *EchoEngine) OPTIONS(path string, process web.Process) {
	echoEngine.delegate.OPTIONS(path, func(ctx echo.Context) error {
		return process(fromEchoContext(ctx))
	})
}

//TRACE in EchoEngine
func (echoEngine *EchoEngine) TRACE(path string, process web.Process) {
	echoEngine.delegate.TRACE(path, func(ctx echo.Context) error {
		return process(fromEchoContext(ctx))
	})
}

//CONNECT in EchoEngine
func (echoEngine *EchoEngine) CONNECT(path string, process web.Process) {
	echoEngine.delegate.CONNECT(path, func(ctx echo.Context) error {
		return process(fromEchoContext(ctx))
	})
}
