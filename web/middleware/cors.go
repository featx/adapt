package middleware

import (
	"github.com/featx/goin/web/types"
)

type CORSConfig struct {
	Skipper      func(types.Context) bool
	AllowOrigins []string
	AllowMethods []string
}

func CORSWithConfig(config CORSConfig) types.Process {
	return func(context types.Context) error {
		return nil
	}
}
