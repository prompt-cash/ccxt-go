package app

import (
	"context"
	"github.com/prompt-cash/js-transpiler/log"
	"github.com/spf13/viper"
	"testing"
)

func getContextForTesting(t *testing.T) (*App, *Context) {
	ctx, _ := context.WithCancel(context.Background())

	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	viper.SetDefault("LogLevel", "debug")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	loggerConf := log.NewConfig(viper.GetViper())
	loggerConf.ConsoleLevel = log.NormalizeLogLevel("debug")
	logger, err := log.NewLogger(loggerConf, log.DefaultLogger)
	if err != nil {
		t.Fatalf("Error creating logger for testing %+v", err)
	}

	appMain, err := New(ctx, logger, &ApplicationOptions{})
	if err != nil {
		t.Fatalf("Error creating application for testing %+v", err)
	}
	appCtx := appMain.NewContext()
	// setup logger fields and API function name too?
	return appMain, appCtx
}
