package web

import "github.com/labstack/echo/v4"

func NewEchoEngine() *EchoEngine {
	return &EchoEngine{echo.New()}
}

type EchoEngine struct {
	delegate *echo.Echo
}

func (echo *EchoEngine) Use(middlewareFunc MiddlewareFunc) {
	echo.delegate.Use()
}

func (echo *EchoEngine) Group(relativePath string, middlewares ...MiddlewareFunc) RouterGroup {
	echo.delegate.Group()
}
