package cmd

import (
	"github.com/pureugong/gurume/api"
	"github.com/pureugong/gurume/config"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "api data..",
	Long:  "api data...",
	Run:   apiExecute,
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func apiExecute(cmd *cobra.Command, args []string) {
	esClient := config.NewESClient()
	application := api.NewApp(logger, esClient)
	mc := api.NewMainController(application)
	application.Start(mc.Router())
}
