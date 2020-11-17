package web

import (
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
	config     Config
}

//NewAtreugoEngine alloc new mem for atreugo impl engine from map config and give out the pointer
func NewGorillaEngine(config Config) Engine {
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
func (gorillaEngine *GorillaEngine) Use(processArray ...Process) {

}

//Group in GorillaEngine
func (gorillaEngine *GorillaEngine) Group(path string, processArray ...Process) RouterGroup {
	return nil
}

//GET in GorillaEngine
func (gorillaEngine *GorillaEngine) GET(path string, process Process) {

}

//POST in GorillaEngine
func (gorillaEngine *GorillaEngine) POST(path string, process Process) {

}

//PUT in GorillaEngine
func (gorillaEngine *GorillaEngine) PUT(path string, process Process) {

}

//DELETE in GorillaEngine
func (gorillaEngine *GorillaEngine) DELETE(path string, process Process) {

}

//PATCH in GorillaEngine
func (gorillaEngine *GorillaEngine) PATCH(path string, process Process) {

}

//HEAD in GorillaEngine
func (gorillaEngine *GorillaEngine) HEAD(path string, process Process) {

}

//OPTIONS in GorillaEngine
func (gorillaEngine *GorillaEngine) OPTIONS(path string, process Process) {

}

//TRACE in GorillaEngine
func (gorillaEngine *GorillaEngine) TRACE(path string, process Process) {

}

//CONNECT in GorillaEngine
func (gorillaEngine *GorillaEngine) CONNECT(path string, process Process) {

}
