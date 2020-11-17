package middleware

import (
	"github.com/featx/goin/web"
)

type CORSConfig struct {
	Skipper      func(web.Context) bool
	AllowOrigins []string
	AllowMethods []string
}

func CORSWithConfig(config CORSConfig) web.Process {
	return func(context web.Context) error {
		return nil
	}
}
