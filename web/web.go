package web

import (
	engines "github.com/featx/goin/web/engine"
	"github.com/featx/goin/web/types"
)

func New(engine types.EngineEnum, config types.Config) types.Engine {
	switch engine {
	case types.ATREUGO:
		return engines.NewAtreugoEngine(config)
	case types.ECHO:
		return engines.NewEchoEngine(config)
	case types.GIN:
		return engines.NewGinEngine(config)
	case types.GORILLA:
		return engines.NewGorillaEngine(config)
	case types.GOJI:
		return engines.NewGojiEngine(config)
	//case "iris": return &IrisEngine{}
	//case "revel": return &RevelEngine{}
	//case "buffalo": return &BuffaloEngine{}
	default:
		return engines.NewGorillaEngine(config)
	}
}
