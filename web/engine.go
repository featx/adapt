package web

import (
	"log"
)

type RouterGroup interface {
	GET(path string, middlewareFunc MiddlewareFunc)

	POST(path string, middlewareFunc MiddlewareFunc)

	PUT(path string, middlewareFunc MiddlewareFunc)

	DELETE(path string, middlewareFunc MiddlewareFunc)

	PATCH(path string, middlewareFunc MiddlewareFunc)

	HEAD(path string, middlewareFunc MiddlewareFunc)

	OPTIONS(path string, middlewareFunc MiddlewareFunc)

	TRACE(path string, middlewareFunc MiddlewareFunc)

	CONNECT(path string, middlewareFunc MiddlewareFunc)
}

type Engine interface {
	RouterGroup

	SetLogger(logger *log.Logger)

	Use(middlewareArray ...MiddlewareFunc)

	Group(path string, middlewareArray ...MiddlewareFunc) RouterGroup
}

type MiddlewareFunc func(c Context) error

type Context interface {
	Request()

	Response()

	HTML()

	JSON()

	JSONP()

	XML()
}
