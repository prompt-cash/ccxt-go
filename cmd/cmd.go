package cmd

import (
	"context"
	"fmt"
	"github.com/prompt-cash/js-transpiler/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "Start the auth app. You must specify additional parameters.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is config.yaml)")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	}

	viper.SetDefault("Log.Level", "debug")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("unable to read config: %v\n", err)
		os.Exit(1)
	}
}

func listenExitCommand(logger log.Logger, cancel context.CancelFunc) {
	var count int
	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGTERM, os.Interrupt)
	for {
		select {
		case <-c:
			count++
			if count == 2 {
				logger.Infof("Forcefully exiting...")
				os.Exit(1)
			}
			logger.Infof("Signal SIGTERM caught. shutting down...")
			logger.Infof("Catching SIGTERM one more time will forcefully exit")
			cancel()
		}
	}
}

func getLogger() (log.Logger, error) {
	return log.NewLogger(log.NewConfig(viper.GetViper()), log.DefaultLogger)
}
