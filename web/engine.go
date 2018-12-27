package web

import (
	"context"
	"log"
)

type RouterGroup interface {
	GET(relativePath string, middlewareFunc MiddlewareFunc)

	POST(relativePath string, middlewareFunc MiddlewareFunc)

	PUT(relativePath string, middlewareFunc MiddlewareFunc)

	HEAD(relativePath string, middlewareFunc MiddlewareFunc)

	DELETE(relativePath string, middlewareFunc MiddlewareFunc)

	OPTION(relativePath string, middlewareFunc MiddlewareFunc)
}

type Engine interface {
	RouterGroup
	SetLogger(logger *log.Logger)

	Use(middlewareFunc MiddlewareFunc)

	Group(relativePath string, middlewares ...MiddlewareFunc) RouterGroup
}

type MiddlewareFunc func(c *context.Context) error

type Context interface {
	Request()

	Response()

	HTML()

	JSON()

	JSONP()

	XML()
}
