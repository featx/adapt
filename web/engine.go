package web

import (
	"log"
)

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
type Process func(Context) error

type Middleware func(Process) Process

type ErrorHandler func(error, Context)

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

	Start(args string) error

	ListenAndServe() error

	SetLogger(logger *log.Logger)

	Use(processArray ...Process)

	Group(path string, process ...Process) RouterGroup
}

func New(engine string, config Config) Engine {
	switch engine {
	case "atreugo":
		return NewAtreugoEngine(config)
	case "echo":
		return NewEchoEngine(config)
	case "gin":
		return NewGinEngine(config)
	case "gorilla":
		return NewGorillaEngine(config)
	case "goji":
		return NewGojiEngine(config)
	//case "iris": return &IrisEngine{}
	//case "revel": return &RevelEngine{}
	//case "buffalo": return &BuffaloEngine{}
	default:
		return NewAtreugoEngine(config)
	}
}
