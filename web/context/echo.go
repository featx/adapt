package context

import (
	"github.com/featx/goin/web/types"
	"github.com/labstack/echo/v4"
)

type EchoContext struct {
	delegate echo.Context
}

func (ctx *EchoContext) Request() types.Request {
	return nil
}

func (ctx *EchoContext) Response() types.Response {
	return nil
}

func (ctx *EchoContext) TEXT(content string) {
	err := ctx.delegate.String(200, content)
	if err != nil {
		panic(err)
	}
}

func (ctx *EchoContext) JSON(object interface{}) {
	err := ctx.delegate.JSON(200, object)
	if err != nil {
		panic(err)
	}
}

//ToContext Convert echo Context to Context
func FromEchoContext(ctx echo.Context) types.Context {
	return &EchoContext{ctx}
}

