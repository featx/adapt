package engine

import (
	"github.com/featx/goin/web"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	//"github.com/gorilla/mux"
	//"github.com/gorilla/context"
)

// GorillaEngine The Gorilla implementation of engine
type GorillaEngine struct {
	httpServer *http.Server
	fastServer *fasthttp.Server
	config     web.Config
}

//NewAtreugoEngine alloc new mem for atreugo impl engine from map config and give out the pointer
func NewGorillaEngine(config web.Config) web.Engine {
	return &GorillaEngine{&http.Server{}, &fasthttp.Server{}, config}
}

func (gorillaEngine *GorillaEngine) Start(args string) error {
	gorillaEngine.httpServer.Addr = args
	return gorillaEngine.httpServer.ListenAndServe()
}

func (gorillaEngine *GorillaEngine) ListenAndServe() error {
	//return gorillaEngine.delegate.TLSServer.Serve(gorillaEngine.delegate.TLSListener)
	return gorillaEngine.httpServer.ListenAndServe()
}

//SetLogger in GorillaEngine
func (gorillaEngine *GorillaEngine) SetLogger(logger *log.Logger) {

}

//Use in GorillaEngine
func (gorillaEngine *GorillaEngine) Use(processArray ...web.Process) {

}

//Group in GorillaEngine
func (gorillaEngine *GorillaEngine) Group(path string, processArray ...web.Process) web.RouterGroup {
	return nil
}

//GET in GorillaEngine
func (gorillaEngine *GorillaEngine) GET(path string, process web.Process) {

}

//POST in GorillaEngine
func (gorillaEngine *GorillaEngine) POST(path string, process web.Process) {

}

//PUT in GorillaEngine
func (gorillaEngine *GorillaEngine) PUT(path string, process web.Process) {

}

//DELETE in GorillaEngine
func (gorillaEngine *GorillaEngine) DELETE(path string, process web.Process) {

}

//PATCH in GorillaEngine
func (gorillaEngine *GorillaEngine) PATCH(path string, process web.Process) {

}

//HEAD in GorillaEngine
func (gorillaEngine *GorillaEngine) HEAD(path string, process web.Process) {

}

//OPTIONS in GorillaEngine
func (gorillaEngine *GorillaEngine) OPTIONS(path string, process web.Process) {

}

//TRACE in GorillaEngine
func (gorillaEngine *GorillaEngine) TRACE(path string, process web.Process) {

}

//CONNECT in GorillaEngine
func (gorillaEngine *GorillaEngine) CONNECT(path string, process web.Process) {

}
