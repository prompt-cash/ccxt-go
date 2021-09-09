package app

import (
	"github.com/prompt-cash/js-transpiler/log"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
)

type Config struct {
	LogLevel log.LoggerLevel
}

func initConfig() (*Config, error) {
	//var err error

	config := &Config{
		LogLevel: log.NormalizeLogLevel(viper.GetString("Log.Level")),
	}

	return config, nil
}

// Returns the path folder relative to the parent directory of one directory down (such as testing inside /db as working dir).
func fixRelativeTestingPath(pathStr string) string {
	if filepath.IsAbs(pathStr) == true {
		return pathStr
	}
	workingDir, _ := os.Getwd()
	dirname := path.Base(workingDir)
	if dirname == "app" { // use dirname != project-name (but some people rename the folder)
		return filepath.Join("..", pathStr)
	}
	return pathStr
}
