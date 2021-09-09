package cmd

import (
	"context"
	"github.com/pkg/errors"
	"github.com/prompt-cash/js-transpiler/app"
	"github.com/prompt-cash/js-transpiler/pkg/transpiler"
	"github.com/spf13/cobra"
)

var transpileParams = &transpiler.TranspileParams{}

func init() {
	transpileCmd.Flags().StringVar(&transpileParams.Input, "in", "", "The input dir of JavaScript files.")
	transpileCmd.Flags().StringVar(&transpileParams.Output, "out", "./pkg/ccxt", "The output dir of GoLang files.")
	rootCmd.AddCommand(transpileCmd)
}

var transpileCmd = &cobra.Command{
	Use:   "transpile",
	Short: "Transpile JavaScript code to Go.",
	Long:  `Parse JavaScript files and decompose them into lists of known expressions to write auto-generated Go code of it.'`,

	RunE: func(cmd *cobra.Command, args []string) error {
		logger, err := getLogger()
		if err != nil {
			return err
		}

		// Create the app context
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go listenExitCommand(logger, cancel)

		app, err := app.New(ctx, logger, &app.ApplicationOptions{})
		if err != nil {
			return err
		}

		//if len(transpileParams.Input) == 0 || len(transpileParams.Output) == 0 {
		//	logger.Infof("transpile arguments missing please provide input and output directories")
		//	return  errors.New("invalid parameters")
		//}

		app.Transpiler, err = transpiler.NewTranspiler(&transpiler.TranspilerOptions{
			CCxtDir:   transpileParams.Input,
			CCxtGoDir: transpileParams.Output,
		}, app.Logger)

		err = app.Transpiler.TranspileFiles()

		if err != nil {
			return errors.Wrap(err, "error writing transpiled file")
		}

		return err
	},
}
