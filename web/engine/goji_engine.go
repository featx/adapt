package engine

import (
	"github.com/featx/goin/web/types"
	"log"
	"net/http"

	"github.com/zenazn/goji/web"
)

// GojiEngine The goji implementation of engine
type GojiEngine struct {
	delegate *web.Mux
	config   types.Config
}

//NewGojiEngine alloc new mem for atreugo impl engine from map config and give out the pointer
func NewGojiEngine(config types.Config) types.Engine {
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
func (gojiEngine *GojiEngine) Use(processArray ...types.Process) {
	for _, process := range processArray {
		gojiEngine.delegate.Use(process)
	}
}

//Group in GojiEngine
func (gojiEngine *GojiEngine) Group(path string, processArray ...types.Process) types.RouterGroup {
	return nil
}

//GET in GojiEngine
func (gojiEngine *GojiEngine) GET(path string, process types.Process) {
	gojiEngine.delegate.Get(path, process)
}

//POST in GojiEngine
func (gojiEngine *GojiEngine) POST(path string, process types.Process) {
	gojiEngine.delegate.Post(path, process)
}

//PUT in GojiEngine
func (gojiEngine *GojiEngine) PUT(path string, process types.Process) {
	gojiEngine.delegate.Put(path, process)
}

//DELETE in GojiEngine
func (gojiEngine *GojiEngine) DELETE(path string, process types.Process) {
	gojiEngine.delegate.Delete(path, process)
}

//PATCH in GojiEngine
func (gojiEngine *GojiEngine) PATCH(path string, process types.Process) {
	gojiEngine.delegate.Patch(path, process)
}

//HEAD in GojiEngine
func (gojiEngine *GojiEngine) HEAD(path string, process types.Process) {
	gojiEngine.delegate.Head(path, process)
}

//OPTIONS in GojiEngine
func (gojiEngine *GojiEngine) OPTIONS(path string, process types.Process) {
	gojiEngine.delegate.Options(path, process)
}

//TRACE in GojiEngine
func (gojiEngine *GojiEngine) TRACE(path string, process types.Process) {
	gojiEngine.delegate.Trace(path, process)
}

//CONNECT in GojiEngine
func (gojiEngine *GojiEngine) CONNECT(path string, process types.Process) {
	gojiEngine.delegate.Connect(path, process)
}
