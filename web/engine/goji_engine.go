package engine

import (
	web2 "github.com/featx/goin/web"
	"log"
	"net/http"

	"github.com/zenazn/goji/web"
)

// GojiEngine The goji implementation of engine
type GojiEngine struct {
	delegate *web.Mux
	config   web2.Config
}

//NewGojiEngine alloc new mem for atreugo impl engine from map config and give out the pointer
func NewGojiEngine(config web2.Config) web2.Engine {
	return &GojiEngine{web.New(), config}
}

func (gojiEngine *GojiEngine) Start(args string) error {
	return http.ListenAndServe(args, gojiEngine.delegate)
}

func (gojiEngine *GojiEngine) ListenAndServe() error {
	return http.ListenAndServe(":8080", gojiEngine.delegate)
}

//SetLogger in GojiEngine
func (gojiEngine *GojiEngine) SetLogger(logger *log.Logger) {

}

//Use in GojiEngine
func (gojiEngine *GojiEngine) Use(processArray ...web2.Process) {
	for _, process := range processArray {
		gojiEngine.delegate.Use(process)
	}
}

//Group in GojiEngine
func (gojiEngine *GojiEngine) Group(path string, processArray ...web2.Process) web2.RouterGroup {
	return nil
}

//GET in GojiEngine
func (gojiEngine *GojiEngine) GET(path string, process web2.Process) {
	gojiEngine.delegate.Get(path, process)
}

//POST in GojiEngine
func (gojiEngine *GojiEngine) POST(path string, process web2.Process) {
	gojiEngine.delegate.Post(path, process)
}

//PUT in GojiEngine
func (gojiEngine *GojiEngine) PUT(path string, process web2.Process) {
	gojiEngine.delegate.Put(path, process)
}

//DELETE in GojiEngine
func (gojiEngine *GojiEngine) DELETE(path string, process web2.Process) {
	gojiEngine.delegate.Delete(path, process)
}

//PATCH in GojiEngine
func (gojiEngine *GojiEngine) PATCH(path string, process web2.Process) {
	gojiEngine.delegate.Patch(path, process)
}

//HEAD in GojiEngine
func (gojiEngine *GojiEngine) HEAD(path string, process web2.Process) {
	gojiEngine.delegate.Head(path, process)
}

//OPTIONS in GojiEngine
func (gojiEngine *GojiEngine) OPTIONS(path string, process web2.Process) {
	gojiEngine.delegate.Options(path, process)
}

//TRACE in GojiEngine
func (gojiEngine *GojiEngine) TRACE(path string, process web2.Process) {
	gojiEngine.delegate.Trace(path, process)
}

//CONNECT in GojiEngine
func (gojiEngine *GojiEngine) CONNECT(path string, process web2.Process) {
	gojiEngine.delegate.Connect(path, process)
}
