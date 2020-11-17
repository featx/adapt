package web

//Context type of router
type Context interface {
	Request() Request

	Response() Response

	TEXT(context string)

	JSON(object interface{})
}
