package types

import (
	"log"
	"strings"
)

//Process type of function , expressing the middleware mainly to do
type Process func(Context) error

//Process chain
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

type EngineEnum int32

const (
	GORILLA EngineEnum = 1
	GOJI    EngineEnum = 2
	GIN     EngineEnum = 3
	ECHO    EngineEnum = 4
	ATREUGO EngineEnum = 5
)

func EngineEnumOf(engine string) EngineEnum {
	switch strings.ToLower(strings.TrimSpace(engine)) {
	case "atreugo":
		return ATREUGO
	case "echo":
		return ECHO
	case "gin":
		return GIN
	case "gorilla":
		return GORILLA
	case "goji":
		return GOJI
	default:
		return GORILLA
	}
}