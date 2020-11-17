package types

type Config struct {
}

type Request interface {
}

//Response a collection of Response
type Response interface {
}

//Context type of router
type Context interface {
	Request() Request

	Response() Response

	TEXT(context string)

	JSON(object interface{})
}
