package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "root.. short",
	Long:  "root.. long",
	Run: func(cmd *cobra.Command, args []string) {
		env := viper.GetString("GURUME_ENV")
		logger.WithField("env", env).Info("gurume start")
	},
}

var logger *logrus.Logger

func init() {
	viper.BindEnv("GURUME_ENV")

	viper.BindEnv("LOG_LEVEL")
	logLevel := viper.GetString("LOG_LEVEL")

	logger = logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(level)

}

// Execute is the root cmd entry point
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
