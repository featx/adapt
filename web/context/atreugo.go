package context

import (
	"github.com/featx/goin/web"
	"github.com/savsgio/atreugo/v11"
)

type AtreugoContext struct {
	delegate *atreugo.RequestCtx
}

func (ctx *AtreugoContext) Request() web.Request {
	return nil
}

func (ctx *AtreugoContext) Response() web.Response {
	return nil
}

func (ctx *AtreugoContext) TEXT(content string) {
	err := ctx.delegate.TextResponse(content, 200)
	if err != nil {
		panic(err)
	}
}

func (ctx *AtreugoContext) JSON(object interface{}) {
	err := ctx.delegate.JSONResponse(object, 200)
	if err != nil {
		panic(err)
	}
}

//ToContext Convert atreugo RequestCtx to Context
func FromAtreugoCtx(ctx *atreugo.RequestCtx) web.Context {
	return &AtreugoContext{ctx}
}
