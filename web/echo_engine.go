package web

import "github.com/labstack/echo/v4"

//NewEchoEngine The Echo implementation of engine
func NewEchoEngine() *EchoEngine {
	return &EchoEngine{echo.New()}
}

//EchoEngine The echo implementation of engine
type EchoEngine struct {
	delegate *echo.Echo
}

//Use in EchoEngine
func (echo *EchoEngine) Use(process Process) {
	echo.delegate.Use()
}

//Group in EchoEngine
func (echo *EchoEngine) Group(relativePath string, processArray ...Process) RouterGroup {
	echo.delegate.Group()
}
