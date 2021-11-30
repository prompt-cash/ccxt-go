package app

import (
	"context"
	"github.com/prompt-cash/ccxt-go/log"
	"github.com/prompt-cash/ccxt-go/utils/randstr"
)

type Context struct {
	// App config
	Logger log.Logger
	Config *Config

	RandomGenerator *randstr.RandomStringGenerator

	// User request data
	//RemoteAddress   string    // The remote IP address this request came from
	//ApiFunctionName string    // The API function this request called.
	//BeginTime       time.Time // The time we received this request
	//RequestID       string    // unique ID per request

	ctx context.Context // app main thread context
}

func (app *App) NewContext() *Context {
	return &Context{
		Logger: app.Logger,
		Config: app.Config,

		RandomGenerator: app.RandomGenerator,

		ctx: app.ctx,
	}
}

func (ctx *Context) WithLogger(logger log.Logger) *Context {
	ret := *ctx
	ret.Logger = logger
	return &ret
}

/*func (ctx *Context) WithRemoteAddress(address string) *Context {
	ret := *ctx
	ret.RemoteAddress = address
	return &ret
}*/
