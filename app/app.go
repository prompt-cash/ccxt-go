package app

import (
	"context"
	"github.com/prompt-cash/ccxt-go/log"
	"github.com/prompt-cash/ccxt-go/pkg/transpiler"
	"github.com/prompt-cash/ccxt-go/utils/randstr"
)

type App struct {
	Logger log.Logger
	Config *Config

	Transpiler      *transpiler.Transpiler
	RandomGenerator *randstr.RandomStringGenerator

	ctx context.Context // app main thread context
}

var (
//currencies TokenSet
)

type TokenSet map[string]struct{}

// Exists returns true if the key exists in our pre-defined currency list.
func (c TokenSet) Exists(key string) bool {
	_, ok := c[key]
	return ok
}

// Return all keys in this Set.
func (c TokenSet) Keys() []string {
	keys := make([]string, 0, len(c))
	for key, _ := range c {
		keys = append(keys, key)
	}
	return keys
}

func init() {
	// nothing to do
}

type ApplicationOptions struct {
}

func New(ctx context.Context, logger log.Logger, options *ApplicationOptions) (app *App, err error) {
	if options == nil {
		options = &ApplicationOptions{}
	}
	app = &App{
		Logger: logger,
		ctx:    ctx,
	}
	app.Config, err = initConfig()
	if err != nil {
		return nil, err
	}

	return app, err
}

func (app *App) Close() error {
	// close DB connections here

	return nil
}

// getCcxtDir returns the CCXT_DIR environment variable with some defaults if empty.
/*func getCcxtDir() string {
	ccxtDir := os.Getenv("CCXT_DIR")
	if len(ccxtDir) != 0 {
		return ccxtDir
	}

	switch runtime.GOOS {
	case "windows":
		return viper.GetString("Testing.CCXTDir.Win")
	case "darwin":
		return viper.GetString("Testing.CCXTDir.OSX")
	default:
		panic(fmt.Sprintf("No CCXT default for OS %s. Please set the env variable CCXT_DIR", runtime.GOOS))
	}
}

// getCcxtGoDir returns the CCXT_GO_DIR environment variable with some defaults if empty.
func getCcxtGoDir() string {
	ccxtGoDir := os.Getenv("CCXT_GO_DIR")
	if len(ccxtGoDir) != 0 {
		return ccxtGoDir
	}

	switch runtime.GOOS {
	case "windows":
		return viper.GetString("Testing.CCXTGoDir.Win")
	case "darwin":
		return viper.GetString("Testing.CCXTGoDir.OSX")
	default:
		panic(fmt.Sprintf("No CCXT default for OS %s. Please set the env variable CCXT_GO_DIR", runtime.GOOS))
	}
}*/
