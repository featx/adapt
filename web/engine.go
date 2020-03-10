package web

import (
	"log"
)

//RouterGroup type of router
type RouterGroup interface {
	GET(path string, process Process)

	POST(path string, process Process)

	PUT(path string, process Process)

	DELETE(path string, process Process)

	PATCH(path string, process Process)

	HEAD(path string, process Process)

	OPTIONS(path string, process Process)

	TRACE(path string, process Process)

	CONNECT(path string, process Process)
}

//Engine type of router
type Engine interface {
	RouterGroup

	SetLogger(logger *log.Logger)

	Use(processArray ...Process)

	Group(path string, process ...Process) RouterGroup
}

//Request a collection of request
type Request interface {
}

//Response a collection of Response
type Response interface {
}

//Context type of router
type Context interface {
	Request() Request

	Response() Response

	TEXT()

	HTML()

	JSON()

	JSONP()

	XML()
}

//Process type of function , expressing the middleware mainly to do
type Process func(c Context) error
