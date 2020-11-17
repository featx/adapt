package engine

import (
	"github.com/featx/goin/web/types"
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
	config     types.Config
}

//NewAtreugoEngine alloc new mem for atreugo impl engine from map config and give out the pointer
func NewGorillaEngine(config types.Config) types.Engine {
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
func (gorillaEngine *GorillaEngine) Use(processArray ...types.Process) {

}

//Group in GorillaEngine
func (gorillaEngine *GorillaEngine) Group(path string, processArray ...types.Process) types.RouterGroup {
	return nil
}

//GET in GorillaEngine
func (gorillaEngine *GorillaEngine) GET(path string, process types.Process) {

}

//POST in GorillaEngine
func (gorillaEngine *GorillaEngine) POST(path string, process types.Process) {

}

//PUT in GorillaEngine
func (gorillaEngine *GorillaEngine) PUT(path string, process types.Process) {

}

//DELETE in GorillaEngine
func (gorillaEngine *GorillaEngine) DELETE(path string, process types.Process) {

}

//PATCH in GorillaEngine
func (gorillaEngine *GorillaEngine) PATCH(path string, process types.Process) {

}

//HEAD in GorillaEngine
func (gorillaEngine *GorillaEngine) HEAD(path string, process types.Process) {

}

//OPTIONS in GorillaEngine
func (gorillaEngine *GorillaEngine) OPTIONS(path string, process types.Process) {

}

//TRACE in GorillaEngine
func (gorillaEngine *GorillaEngine) TRACE(path string, process types.Process) {

}

//CONNECT in GorillaEngine
func (gorillaEngine *GorillaEngine) CONNECT(path string, process types.Process) {

}
